package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/account"
)

func GetAccountList(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	//1. 获取参数
	memberId := ctx.GetInt64("memberId")
	var req model_project.GetAccountListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	if req.SearchType == 0 {
		req.SearchType = 1
	}
	// 调用grpc服务获取数据
	grpcResp, err := AccountServiceClient.GetAccountList(ctx, &account.GetAccountListReq{
		MemberId:       memberId,
		Page:           int64(req.Page),
		PageSize:       int64(req.PageSize),
		SearchType:     int32(req.SearchType),
		DepartmentCode: req.DepartmentCode,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 组织回复消息
	resp := &model_project.GetAccountListResp{}
	copier.Copy(&resp.ProjectAuthList, grpcResp.AuthList)
	copier.Copy(&resp.MemberAccountList, grpcResp.AccountList)
	resp.Total = grpcResp.Total
	resp.Page = req.Page
	ctx.JSON(http.StatusOK, result.Success(resp))
}
