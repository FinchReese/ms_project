package project

import (
	"github.com/gin-gonic/gin"
	"test.com/project-api/api/midd"
	"test.com/project-api/router"
)

type ProjectRouter struct {
}

func (*ProjectRouter) Register(r *gin.Engine) {
	projectGroup := r.Group("/project")
	projectGroup.Use(midd.VerifyToken())
	projectGroup.POST("/index", index)
	projectGroup.POST("/project", selfList)
	projectGroup.POST("/project/selfList", selfList)
	projectGroup.POST("/project_template", projectTemplate)
	projectGroup.POST("/project/save", saveProject)
	projectGroup.POST("/project/read", getProjectInfo)
	projectGroup.POST("/project_collect/collect", collectProject)
	projectGroup.POST("/project/recycle", recycleProject)
	projectGroup.POST("/project/recovery", recoveryProject)
	projectGroup.POST("/project/edit", updateProject)
	projectGroup.POST("/task_stages", getTaskStage)
	projectGroup.POST("/project_member/index", getProjectMemberList)
	projectGroup.POST("/task_stages/tasks", getTaskList)
	projectGroup.POST("/task/save", saveTask)
	projectGroup.POST("/task/sort", moveTask)
	projectGroup.POST("/task/selfList", getTaskListByType)
	projectGroup.POST("/task/read", getTaskDetail)
	projectGroup.POST("/task_member", getTaskMemberList)
	projectGroup.POST("/task/taskLog", getTaskLogList)
	projectGroup.POST("/task/_taskWorkTimeList", getTaskWorkTimeList)
	projectGroup.POST("/task/saveTaskWorkTime", saveTaskWorkTime)
	projectGroup.POST("/file/uploadFiles", uploadFile)
	projectGroup.POST("/task/taskSources", getTaskLinkFiles)
	projectGroup.POST("/task/createComment", createComment)
	projectGroup.POST("/project/getLogBySelfProject", getUserProjectLogList)
	projectGroup.POST("/account", getAccountList)
	projectGroup.POST("/department", getDepartmentList)
	projectGroup.POST("/department/save", addDepartment)
	projectGroup.POST("/department/read", getDepartmentById)
}

func init() {
	InitProjectRpc()
	router.RegisterRouter(&ProjectRouter{})
}
