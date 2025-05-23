package project

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/project_auth"
)

func getProjectAuthList(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	var req model_project.GetProjectAuthListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数传递有误"))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}

	// 调用RPC服务获取项目权限列表
	projectAuthList, err := ProjectAuthServiceClient.GetProjectAuthList(ctx, &project_auth.GetProjectAuthListReq{
		MemberId: ctx.GetInt64("memberId"),
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 组织回复消息
	resp := &model_project.GetProjectAuthListResp{
		Total: projectAuthList.Total,
		Page:  req.Page,
	}
	copier.Copy(&resp.List, projectAuthList.List)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getProjectNodeApply(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	var req model_project.ProjectNodeApplyReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	var nodes []string
	if req.Nodes != "" {
		err = json.Unmarshal([]byte(req.Nodes), &nodes)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, result.Fail(http.StatusBadRequest, "参数传递有误"))
			return
		}
	}

	// 调用RPC服务获取项目节点列表
	projectNodeApply, err := ProjectAuthServiceClient.ProjectAuthNodeApply(ctx, &project_auth.ProjectAuthNodeApplyReq{
		AuthId:   req.Id,
		Action:   req.Action,
		NodeList: nodes,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}

	// 组织回复消息
	if req.Action == "getnode" {
		resp := &model_project.ProjectNodeApplyResp{
			CheckedList: projectNodeApply.CheckedList,
		}
		copier.Copy(&resp.List, projectNodeApply.List)
		ctx.JSON(http.StatusOK, result.Success(resp))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(""))
}

func GetAuthNodeUrlList(ctx *gin.Context) ([]string, error) {
	// 调用grpc服务获取member id有权限的节点URL列表
	memberId := ctx.GetInt64("memberId")
	authNodeUrlList, err := ProjectAuthServiceClient.GetAuthNodeUrlList(ctx, &project_auth.GetAuthNodeUrlListReq{
		MemberCode: memberId,
	})
	if err != nil {
		return nil, err
	}
	return authNodeUrlList.List, nil
}
