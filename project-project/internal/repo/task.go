package repo

import (
	"context"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
)

type TaskStageRepo interface {
	SaveTaskStage(ctx context.Context, ts *data.TaskStage, db *gorm.DB) error
	FindStagesByProjectId(ctx context.Context, projectId int64, page int, pageSize int) (list []*data.TaskStage, total int64, err error)
	GetTaskStageByID(ctx context.Context, id int) (*data.TaskStage, error)
}

type TaskRepo interface {
	FindTasksByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error)
	GetMaxIdNumByProjectID(ctx context.Context, projectID int64) (int, error)
	// 根据项目id和阶段编码获取任务表中sort字段的最大值
	GetMaxSortByProjectIDAndStageCode(ctx context.Context, projectID int64, stageCode int) (int, error)
	// 保存任务
	SaveTask(ctx context.Context, task *data.Task, db *gorm.DB) error
	// 修改任务所属步骤
	ModifyStageCode(ctx context.Context, taskId int64, stageCode int, db *gorm.DB) error
	// 将指定步骤的大于等于sort阈值的任务的sort加1
	IncreaseSort(ctx context.Context, projectID int64, stageCode int, sort int, db *gorm.DB) error
	// 根据id获取任务
	GetTaskById(ctx context.Context, taskId int64) (*data.Task, error)
	// 修改指定任务的sort
	ModifyTaskSort(ctx context.Context, taskId int64, sort int32, db *gorm.DB) error
	// 指定assign_to、done字段筛选任务，再根据指定的页号和页大小返回任务列表
	GetTasksByAssignToAndDone(ctx context.Context, assignTo int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error)
	// 指定成员id、done字段筛选任务，再根据指定的页号和页大小返回任务列表
	GetTasksByMemberIdAndDone(ctx context.Context, memberId int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error)
	// 指定create_by、done字段筛选任务，再根据指定的页号和页大小返回任务列表
	GetTasksByCreateByAndDone(ctx context.Context, createBy int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error)
}

type TaskMemberRepo interface {
	FindTaskMembers(ctx context.Context, taskCode int64, memberCode int64) (list []*data.TaskMember, err error)
	// 保存任务成员关系
	SaveTaskMember(ctx context.Context, taskMember *data.TaskMember, db *gorm.DB) error
}
