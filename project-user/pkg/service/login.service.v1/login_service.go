package login_service_v1

import (
	"context"
	"log"
	"time"

	"go.uber.org/zap"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	Cache repo.Cache
}

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
		err := ls.Cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis发生错误，cause by :", err)
		}
		log.Println("发送短信成功")
	}()
	return &login.CaptchaResponse{Code: code}, nil
}
