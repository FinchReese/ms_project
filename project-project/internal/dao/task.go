package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	custom_gorm "test.com/project-project/internal/database/gorm"
)

type TaskDAO struct {
	conn *custom_gorm.MysqlConn
}

func NewTaskDAO() *TaskDAO {
	return &TaskDAO{
		conn: custom_gorm.NewMysqlConn(),
	}
}

func (td *TaskDAO) FindTasksByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Model(&data.Task{}).
		Where("stage_code = ? and deleted = 0", stageCode).
		Order("sort asc").
		Find(&list).Error
	return
}

// GetMaxIdNumByProjectID 根据项目id获取任务表中id_num字段的最大值
func (td *TaskDAO) GetMaxIdNumByProjectID(ctx context.Context, projectID int64) (int, error) {
	var maxIdNum int

	err := td.conn.Db.Session(&gorm.Session{Context: ctx}).
		Model(&data.Task{}).
		Where("project_code = ? and deleted = 0", projectID).
		Select("IFNULL(MAX(id_num), 0) as max_id_num").
		Pluck("max_id_num", &maxIdNum).Error

	if err != nil {
		return 0, err
	}

	return maxIdNum, nil
}

// 根据项目id和阶段编码获取任务表中sort字段的最大值
func (td *TaskDAO) GetMaxSortByProjectIDAndStageCode(ctx context.Context, projectID int64, stageCode int) (int, error) {
	var maxSort int

	err := td.conn.Db.Session(&gorm.Session{Context: ctx}).
		Model(&data.Task{}).
		Where("project_code = ? AND stage_code = ? AND deleted = 0", projectID, stageCode).
		Select("IFNULL(MAX(sort), 0) as max_sort").
		Pluck("max_sort", &maxSort).Error

	if err != nil {
		return 0, err
	}

	return maxSort, nil
}

// SaveTask 保存任务
func (td *TaskDAO) SaveTask(ctx context.Context, task *data.Task, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Save(task).Error
}

// 修改任务所属步骤
func (td *TaskDAO) ModifyStageCode(ctx context.Context, taskId int64, stageCode int, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Model(&data.Task{}).Where("id = ?", taskId).Update("stage_code", stageCode).Error
}

// 将指定步骤的大于等于sort阈值的任务的sort加1
func (td *TaskDAO) IncreaseSort(ctx context.Context, projectID int64, stageCode int, sort int, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Model(&data.Task{}).Where("project_code = ? AND stage_code = ? AND sort >= ?", projectID, stageCode, sort).Update("sort", gorm.Expr("sort + 1")).Error
}

func (td *TaskDAO) GetTaskById(ctx context.Context, taskId int64) (*data.Task, error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	var task data.Task
	err := session.Model(&data.Task{}).Where("id = ?", taskId).First(&task).Error
	return &task, err
}

func (td *TaskDAO) ModifyTaskSort(ctx context.Context, taskId int64, sort int32, db *gorm.DB) error {
	if db == nil {
		return errors.New("db is nil")
	}
	return db.Model(&data.Task{}).Where("id = ?", taskId).Update("sort", sort).Error
}
