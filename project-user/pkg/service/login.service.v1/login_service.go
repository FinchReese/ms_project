package login_service_v1

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	common "test.com/project-common"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-common/jwts"
	"test.com/project-common/time_format"
	"test.com/project-grpc/user/login"
	"test.com/project-user/config"
	"test.com/project-user/pkg/data/member"
	"test.com/project-user/pkg/data/organization"
	"test.com/project-user/pkg/database/gorm"
	tran "test.com/project-user/pkg/database/trans"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	Cache            repo.Cache
	MemberRepo       repo.Member
	OrganizationRepo repo.OrganizationRepo
	Tran             *tran.TransactionImpl
}

const (
	mobileCaptchaRedisKeyPrefix = "REGISTER_"
	redisPutTimeout             = 2  // redis Put操作超时时间设置为2秒
	redisGetTimeout             = 2  // redis Get操作超时时间设置为2秒
	captchaRedisCacheMin        = 15 // 验证码在Redis的缓存时间设置为15分钟
	aesKey                      = "sdfgyrhgbxcdgryfhgywertd"
	tokenType                   = "bearer"
)

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *login.CaptchaMessage) (*login.CaptchaResponse, error) {
	//1. 获取参数
	mobile := msg.GetMobile()
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("调用短信平台发送短信")
		zap.L().Debug("调用短信平台发送短信debug")
		zap.L().Warn("调用短信平台发送短信warn")
		zap.L().Error("调用短信平台发送短信error")
		//发送成功 存入redis
		ctx, cancel := context.WithTimeout(context.Background(), redisPutTimeout*time.Second)
		defer cancel()
		err := ls.Cache.Put(ctx, mobileCaptchaRedisKeyPrefix+mobile, code, captchaRedisCacheMin*time.Minute)
		if err != nil {
			log.Println("验证码存入redis发生错误，cause by :", err)
		}
		log.Println("发送短信成功")
	}()
	return &login.CaptchaResponse{Code: code}, nil
}

func (ls *LoginService) Register(ctx context.Context, msg *login.RegisterMessage) (*login.RegisterResponse, error) {
	// 校验验证码
	mobile := msg.GetMobile()
	inputCaptcha := msg.GetCaptcha()
	getCtx, cancel := context.WithTimeout(context.Background(), redisGetTimeout*time.Second)
	defer cancel()
	expectCaptcha, err := ls.Cache.Get(getCtx, mobileCaptchaRedisKeyPrefix+mobile)
	if err != nil {
		zap.L().Error("获取手机验证码缓存失败", zap.String("手机号", mobile))
		return nil, errs.GrpcError(model.CaptchaNotExist)
	}
	if inputCaptcha != expectCaptcha {
		zap.L().Error("验证码校验失败", zap.String("手机号", mobile), zap.String("输入验证码", inputCaptcha))
		return nil, errs.GrpcError(model.CaptchaError)
	}
	// 校验账号是否已经被注册
	account := msg.GetName()
	ret, err := ls.MemberRepo.IsAccountRegistered(ctx, account)
	if err != nil {
		zap.L().Error("数据库查询账号异常", zap.String("账号", account))
		return nil, errs.GrpcError(model.QueryDbError)
	}
	if ret {
		zap.L().Error("账号已存在", zap.String("账号", account))
		return nil, errs.GrpcError(model.AccountExist)
	}
	// 校验邮箱是否已经被注册
	email := msg.GetEmail()
	ret, err = ls.MemberRepo.IsEmailRegistered(ctx, email)
	if err != nil {
		zap.L().Error("数据库查询邮箱异常", zap.String("邮箱", email))
		return nil, errs.GrpcError(model.QueryDbError)
	}
	if ret {
		zap.L().Error("邮箱已存在", zap.String("邮箱", email))
		return nil, errs.GrpcError(model.EmailExist)
	}
	// 校验手机号是否已经被注册
	ret, err = ls.MemberRepo.IsMobileRegistered(ctx, mobile)
	if err != nil {
		zap.L().Error("数据库查询手机号异常", zap.String("手机号", mobile))
		return nil, errs.GrpcError(model.QueryDbError)
	}
	if ret {
		zap.L().Error("手机号已存在", zap.String("手机号", mobile))
		return nil, errs.GrpcError(model.MobileExist)
	}
	err = ls.Tran.ExecTran(func(dbConn tran.DbConn) error {

		// 如果前面校验通过，则把数据存入成员表
		pwd := encrypt.Md5(msg.GetPassword())
		member := &member.Member{
			Account:       msg.Name,
			Password:      pwd,
			Name:          msg.Name,
			Mobile:        msg.Mobile,
			Email:         msg.Email,
			CreateTime:    time.Now().UnixMilli(),
			LastLoginTime: time.Now().UnixMilli(),
			Status:        member.MemberStatusNormal,
		}
		conn := dbConn.(*gorm.MysqlConn)
		err = ls.MemberRepo.RegisterMember(ctx, member, conn.TranDb)
		if err != nil {
			zap.L().Error("register member err", zap.Error(err))
			return errs.GrpcError(model.RegisterMemberError)
		}
		// 存入组织表
		//存入组织
		org := &organization.Organization{
			Name:       member.Name + "个人组织",
			MemberId:   member.Id,
			CreateTime: time.Now().UnixMilli(),
			Personal:   organization.OrganizationPersion,
			Avatar:     "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		}
		err = ls.OrganizationRepo.RegisterOrganization(ctx, org, conn.TranDb)
		if err != nil {
			zap.L().Error("register SaveOrganization err", zap.Error(err))
			return errs.GrpcError(model.RegisterOrganizationError)
		}
		return nil
	})

	return &login.RegisterResponse{}, err
}

func (ls *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	account := msg.GetAccount()
	pwd := encrypt.Md5(msg.GetPassword())
	member, err := ls.MemberRepo.LoginVerify(ctx, account, pwd)
	if err != nil {
		zap.L().Error("login vefiry err", zap.Error(err))
		return nil, errs.GrpcError(model.AccountOrPwdError)
	}

	memMsg := &login.MemberMessage{}
	err = copier.Copy(&memMsg, &member)
	if err != nil {
		zap.L().Error("copy member msg err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyMemberMsgError)
	}
	memMsg.Code, err = encrypt.EncryptInt64(memMsg.Id, aesKey)
	if err != nil {
		zap.L().Error("encrypt member id error", zap.Error(err))
		return nil, errs.GrpcError(model.EncryptMemberIdError)
	}
	memMsg.LastLoginTime = time_format.ConvertMsecToString(member.LastLoginTime)
	memMsg.CreateTime = time_format.ConvertMsecToString(member.CreateTime)

	orgs, err := ls.OrganizationRepo.GetOrganizationByMemberId(ctx, member.Id)
	if err != nil {
		zap.L().Error("get organization msg err", zap.Error(err))
		return nil, errs.GrpcError(model.GetOrganizationMsgError)
	}
	orgMsgs := []*login.OrganizationMessage{}
	err = copier.Copy(&orgMsgs, &orgs)
	if err != nil {
		zap.L().Error("copy organization msg err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyOrganizationMsgError)
	}
	for _, orgMsg := range orgMsgs {
		orgMsg.Code, err = encrypt.EncryptInt64(orgMsg.Id, aesKey)
		if err != nil {
			zap.L().Error("encrypt member id error", zap.Error(err))
			return nil, errs.GrpcError(model.EncryptOrganizationIdError)
		}
		orgMsg.OwnerCode = memMsg.Code
	}
	if len(orgMsgs) > 0 {
		memMsg.OrganizationCode = orgMsgs[0].Code
	}

	exp := time.Duration(config.AppConf.JwtConf.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.AppConf.JwtConf.RefreshExp*3600*24) * time.Second
	token := jwts.CreateToken(member.Id, exp, config.AppConf.JwtConf.AccessSecret, rExp, config.AppConf.JwtConf.RefreshSecret, msg.Ip)
	tokenMsg := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessTokenExp: token.AccessExp,
		TokenType:      tokenType,
	}
	// 缓存成员信息到redis
	marshal, _ := json.Marshal(member)
	ls.Cache.Put(ctx, model.MemberKeyPrefix+"::"+strconv.FormatInt(member.Id, 10), string(marshal), exp)

	resp := &login.LoginResponse{
		Member:           memMsg,
		OrganizationList: orgMsgs,
		TokenList:        tokenMsg,
	}

	return resp, nil
}

func (ls *LoginService) VerifyToken(ctx context.Context, req *login.VerifyTokenReq) (*login.VerifyTokenResp, error) {
	// 解析请求消息获得token
	token := req.GetToken()
	if strings.Contains(token, tokenType) {
		token = strings.ReplaceAll(token, tokenType+" ", "")
	}
	memberId, err := jwts.ParseToken(token, config.AppConf.JwtConf.AccessSecret, req.Ip)
	if err != nil {
		zap.L().Error("verify token error", zap.Error(err))
		return nil, errs.GrpcError(model.VerifyTokenError)
	}
	// 从redis中获取member
	memberStr, err := ls.Cache.Get(ctx, model.MemberKeyPrefix+"::"+strconv.FormatInt(memberId, 10))
	if err != nil {
		zap.L().Error("get member from redis err", zap.Error(err))
		return nil, errs.GrpcError(model.GetMemberFromRedisError)
	}
	if memberStr == "" {
		zap.L().Error("get member expire", zap.Int64("memberId", memberId))
		return nil, errs.GrpcError(model.LoginTimeoutError)
	}
	var member member.Member
	err = json.Unmarshal([]byte(memberStr), &member)
	if err != nil {
		zap.L().Error("unmarshal member from redis err", zap.Error(err))
		return nil, errs.GrpcError(model.UnmarshalMemberFromRedisError)
	}

	memMsg := login.MemberMessage{}
	copier.Copy(&memMsg, member)
	resp := &login.VerifyTokenResp{Member: &memMsg}
	return resp, nil
}

func (ls *LoginService) GetOrganizationList(ctx context.Context, req *login.GetOrganizationListReq) (*login.GetOrganizationListResp, error) {
	memberId := req.GetMemberId()
	orgs, err := ls.OrganizationRepo.GetOrganizationByMemberId(ctx, memberId)
	if err != nil {
		zap.L().Error("get organization msg err", zap.Error(err))
		return nil, errs.GrpcError(model.GetOrganizationMsgError)
	}
	orgMsgs := []*login.OrganizationMessage{}
	err = copier.Copy(&orgMsgs, &orgs)
	if err != nil {
		zap.L().Error("copy organization msg err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyOrganizationMsgError)
	}
	for _, orgMsg := range orgMsgs {
		orgMsg.Code, err = encrypt.EncryptInt64(orgMsg.Id, aesKey)
		if err != nil {
			zap.L().Error("encrypt member id error", zap.Error(err))
			return nil, errs.GrpcError(model.EncryptOrganizationIdError)
		}
	}
	resp := &login.GetOrganizationListResp{
		OrgList: orgMsgs,
	}
	return resp, nil
}

func (ls *LoginService) GetMemberById(ctx context.Context, req *login.GetMemberByIdReq) (*login.MemberMessage, error) {
	// 解析参数
	memberId := req.GetMemberId()
	// 数据库查询
	member, err := ls.MemberRepo.FindMemberById(ctx, memberId)
	if err != nil {
		zap.L().Error("find member by id err", zap.Error(err))
		return nil, errs.GrpcError(model.FindMemberByIdError)
	}
	resp := &login.MemberMessage{}
	copier.Copy(&resp, member)
	return resp, nil
}

func (ls *LoginService) GetMembersByIds(ctx context.Context, req *login.GetMembersByIdsReq) (*login.GetMembersByIdsResp, error) {
	// 调用之前实现的 DAO 层方法获取成员列表
	members, err := ls.MemberRepo.FindMembersByIds(ctx, req.MemberIds)
	if err != nil {
		zap.L().Error("find members by ids error", zap.Error(err))
		return nil, errs.GrpcError(model.FindMembersByIdsError)
	}

	// 转换为响应消息
	memberMessages := make([]*login.MemberMessage, 0)
	for _, m := range members {
		memberMsg := &login.MemberMessage{}
		copier.Copy(memberMsg, m)
		memberMessages = append(memberMessages, memberMsg)
	}

	return &login.GetMembersByIdsResp{
		List: memberMessages,
	}, nil
}
