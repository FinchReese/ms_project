package project

import (
	"context"
	"net/http"
	"time"

	"test.com/project-grpc/project_node"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
)

func getProjectNodeList(c *gin.Context) {
	// 调用grpc服务获取项目节点列表
	result := &common.Result{}
	grpcCtx, cancel := context.WithTimeout(context.Background(), serviceTimeOut*time.Second)
	defer cancel()
	grpcResp, err := ProjectNodeServiceClient.GetProjectNodeList(grpcCtx, &project_node.GetProjectNodeListReq{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 组织回复消息
	var resp []*model_project.ProjectNodeTree
	copier.Copy(&resp, grpcResp.Nodes)
	c.JSON(http.StatusOK, result.Success(resp))
}
