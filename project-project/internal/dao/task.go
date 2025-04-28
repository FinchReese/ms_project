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

// 指定assign_to、done字段筛选任务，再根据指定的页号和页大小返回任务列表
func (td *TaskDAO) GetTasksByAssignToAndDone(ctx context.Context, assignTo int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Model(&data.Task{}).
		Where("assign_to = ? AND done = ?", assignTo, done).
		Order("sort asc").Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	err = session.Model(&data.Task{}).
		Where("assign_to = ? AND done = ?", assignTo, done).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// 指定成员id、done字段筛选任务，再根据指定的页号和页大小返回任务列表
func (td *TaskDAO) GetTasksByMemberIdAndDone(ctx context.Context, memberId int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	getTaskListSql := "select t.* from ms_task AS t, ms_task_member AS tm where t.id = tm.task_code AND tm.member_code = ? AND t.done = ? order by t.sort asc limit ?, ?"
	err = session.Raw(getTaskListSql, memberId, done, (page-1)*pageSize, pageSize).Scan(&list).Error
	if err != nil {
		return nil, 0, err
	}
	getTaskCountSql := "select count(*) from ms_task AS t, ms_task_member AS tm where t.id = tm.task_code AND tm.member_code = ? AND t.done = ?"
	err = session.Raw(getTaskCountSql, memberId, done).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// 指定create_by、done字段筛选任务，再根据指定的页号和页大小返回任务列表
func (td *TaskDAO) GetTasksByCreateByAndDone(ctx context.Context, createBy int64, done int, page int, pageSize int) (list []*data.Task, total int64, err error) {
	session := td.conn.Db.Session(&gorm.Session{Context: ctx})
	err = session.Model(&data.Task{}).
		Where("create_by = ? AND done = ?", createBy, done).
		Order("sort asc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	err = session.Model(&data.Task{}).
		Where("create_by = ? AND done = ?", createBy, done).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
