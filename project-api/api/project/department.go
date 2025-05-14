package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	model_project "test.com/project-api/pkg/model/project"
	common "test.com/project-common"
	"test.com/project-common/errs"
	"test.com/project-grpc/department"
)

func getDepartmentList(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	var req model_project.GetDepartmentListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}
	// 调用rpc接口获取部门列表
	departmentList, err := DepartmentServiceClient.GetDepartmentList(ctx, &department.GetDepartmentListReq{
		MemberId: ctx.GetInt64("memberId"),
		Pcode:    req.Pcode,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 组织回复消息
	resp := &model_project.GetDepartmentListResp{
		Page:  req.Page,
		Total: departmentList.Total,
	}
	copier.Copy(&resp.Departments, departmentList.Departments)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func addDepartment(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	var req model_project.AddDepartmentReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
	}
	// 调用rpc接口添加部门
	department, err := DepartmentServiceClient.AddDepartment(ctx, &department.AddDepartmentReq{
		MemberId:       ctx.GetInt64("memberId"),
		DepartmentCode: req.DepartmentCode,
		Pcode:          req.ParentDepartmentCode,
		Name:           req.Name,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
		return
	}
	// 组织回复消息
	resp := &model_project.Department{}
	copier.Copy(resp, department)
	ctx.JSON(http.StatusOK, result.Success(resp))
}

func getDepartmentById(ctx *gin.Context) {
	// 解析请求消息
	result := &common.Result{}
	var req model_project.GetDepartmentById
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
	}
	// 调用rpc接口获取部门信息
	department, err := DepartmentServiceClient.GetDepartmentById(ctx, &department.GetDepartmentByIdReq{
		DepartmentId: req.DepartmentCode,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusInternalServerError, result.Fail(code, msg))
	}
	// 组织回复消息
	resp := &model_project.Department{}
	copier.Copy(resp, department)
	ctx.JSON(http.StatusOK, result.Success(resp))
}
