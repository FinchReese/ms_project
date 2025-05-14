package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-grpc/menu"
)

func getMenuTree(ctx *gin.Context) {
	result := &common.Result{}
	// 调用grpc服务获取菜单树
	menuTree, err := MenuServiceClient.GetMenuTree(ctx, &menu.GetMenuTreeReq{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// 组织回复消息
	var menuList []*model_project.Menu
	copier.Copy(&menuList, menuTree.MenuTree)
	ctx.JSON(http.StatusOK, result.Success(menuList))
}
