package login

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"test.com/project-api/api/rpc"
	"test.com/project-api/pkg/model/user"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/user/login"
)

const (
	serviceTimeOut = 2 // 服务超时时间设置为2秒
)

func getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	rsp, err := rpc.LoginServiceClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}

func registerUser(c *gin.Context) {
	// 解析参数
	result := &common.Result{}
	var req user.RegisterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	// 校验参数
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//处理业务
	msg := &login.RegisterMessage{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
		Captcha:  req.Captcha,
	}
	_, err = rpc.LoginServiceClient.Register(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(nil))
}

func projectLogin(c *gin.Context) {
	result := &common.Result{}
	// 解析参数
	var req user.LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	// 调用grpc接口完成登录验证
	ctx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	msg := &login.LoginMessage{
		Account:  req.Account,
		Password: req.Password,
	}
	loginRsp, err := rpc.LoginServiceClient.Login(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	rsp := &user.LoginRsp{}
	err = copier.Copy(rsp, loginRsp)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "copy有误"))
		return
	}
	c.JSON(http.StatusOK, result.Success(rsp))
}

func getOrgList(ctx *gin.Context) {
	result := &common.Result{}
	val, _ := ctx.Get("memberId")
	memberId := val.(int64)
	fmt.Println("call getOrgList, memberId:", memberId)
	grpcCtx, cancelFunc := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancelFunc()
	resp, err := rpc.LoginServiceClient.GetOrganizationList(grpcCtx, &login.GetOrganizationListReq{MemberId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	if resp.OrgList == nil {
		ctx.JSON(http.StatusOK, result.Success([]*user.OrganizationList{}))
		return
	}
	var orgList []*user.OrganizationList
	err = copier.Copy(&orgList, resp.OrgList)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "copy有误"))
		return
	}
	fmt.Println("call getOrgList, orgList:", orgList)
	ctx.JSON(http.StatusOK, result.Success(orgList))
}
