package task_service_v1

import (
	"context"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-common/time_format"
	"test.com/project-grpc/task"
	"test.com/project-grpc/user/login"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/database/gorm"
	"test.com/project-project/internal/database/trans"
	"test.com/project-project/internal/repo"
	"test.com/project-project/internal/rpc"
	"test.com/project-project/pkg/model"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
	taskStage  repo.TaskStageRepo
	task       repo.TaskRepo
	taskMember repo.TaskMemberRepo
	project    repo.ProjectRepo
	tran       *trans.TransactionImpl
}

func NewTaskService(ts repo.TaskStageRepo, t repo.TaskRepo, tm repo.TaskMemberRepo, p repo.ProjectRepo, tran *trans.TransactionImpl) *TaskService {
	return &TaskService{
		taskStage:  ts,
		task:       t,
		taskMember: tm,
		project:    p,
		tran:       tran,
	}
}

func (ts *TaskService) GetTaskStages(ctx context.Context, req *task.GetTaskStagesReq) (*task.GetTaskStagesResp, error) {
	// 解析请求消息获得参数
	projectCodeStr, _ := encrypt.Decrypt(req.ProjectCode, model.AESKey)
	projectId, _ := strconv.ParseInt(projectCodeStr, 10, 64)
	taskStageList, total, err := ts.taskStage.FindStagesByProjectId(ctx, projectId, int(req.Page), int(req.PageSize))
	if err != nil {
		zap.L().Error("encrypt project id err", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskStagesError)
	}

	var taskStages []*task.TaskStage
	err = copier.Copy(&taskStages, taskStageList)
	if err != nil {
		zap.L().Error("copy task stage list err", zap.Error(err))
		return nil, errs.GrpcError(model.CopyTaskStageListError)
	}
	if taskStages == nil {
		return &task.GetTaskStagesResp{List: taskStages, Total: 0}, nil
	}
	for _, taskStage := range taskStages {
		taskStage.Code, err = encrypt.EncryptInt64(int64(taskStage.Id), model.AESKey)
		if err != nil {
			zap.L().Error("encrypt task stage id err", zap.Error(err))
			return nil, errs.GrpcError(model.EncryptTaskStageIdError)
		}
		taskStage.ProjectCode = req.ProjectCode
	}

	return &task.GetTaskStagesResp{List: taskStages, Total: total}, nil
}

func (ts *TaskService) GetTasksByStageCode(ctx context.Context, in *task.GetTasksByStageCodeReq) (*task.GetTasksByStageCodeResp, error) {
	// 解析请求消息获得参数
	stageCodeStr, _ := encrypt.Decrypt(in.StageCode, model.AESKey)
	stageId, _ := strconv.ParseInt(stageCodeStr, 10, 64)
	taskList, err := ts.task.FindTasksByStageCode(ctx, int(stageId))
	if err != nil {
		zap.L().Error("find tasks by stage code err", zap.Error(err))
		return nil, errs.GrpcError(model.GetTasksByStageCodeError)
	}

	if taskList == nil || len(taskList) == 0 {
		return &task.GetTasksByStageCodeResp{List: []*task.Task{}}, nil
	}

	var dispTaskList []*task.Task
	memIdList := []int64{}
	for _, task := range taskList {
		dispTask := task.ToDisplayTask()
		// 私人的任务需要判断是否是当前用户所有
		if dispTask.Private == 1 {
			taskMemberList, err := ts.taskMember.FindTaskMembers(ctx, dispTask.Id, in.MemberId)
			if err != nil {
				zap.L().Error("find task members err", zap.Error(err))
				return nil, errs.GrpcError(model.GetTaskMembersError)
			}
			if taskMemberList != nil && len(taskMemberList) > 0 {
				dispTask.CanRead = model.CanRead
			} else {
				dispTask.CanRead = model.CanNotRead
			}
		}
		assignToStr, _ := encrypt.Decrypt(dispTask.AssignTo, model.AESKey)
		memId, _ := strconv.ParseInt(assignToStr, 10, 64)
		memIdList = append(memIdList, memId)
		dispTaskList = append(dispTaskList, dispTask)
	}

	// 向login服务一次性查询所有成员信息，提升性能
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memIdToInfo[member.Id] = member
	}
	for _, dispTask := range dispTaskList {
		assignToStr, _ := encrypt.Decrypt(dispTask.AssignTo, model.AESKey)
		memId, _ := strconv.ParseInt(assignToStr, 10, 64)
		member := memIdToInfo[memId]
		e := &task.ExecutorMessage{
			Name:   member.Name,
			Avatar: member.Avatar,
		}
		dispTask.Executor = e
	}
	return &task.GetTasksByStageCodeResp{List: dispTaskList}, nil
}

// SaveTask 创建任务
func (ts *TaskService) SaveTask(ctx context.Context, req *task.SaveTaskReq) (*task.SaveTaskResp, error) {
	// 1. 解析请求参数
	stageCodeStr, _ := encrypt.Decrypt(req.StageCode, model.AESKey)
	stageID, _ := strconv.ParseInt(stageCodeStr, 10, 64)

	projectCodeStr, _ := encrypt.Decrypt(req.ProjectCode, model.AESKey)
	projectID, _ := strconv.ParseInt(projectCodeStr, 10, 64)

	// 2. 获取项目步骤信息
	stage, err := ts.taskStage.GetTaskStageByID(ctx, int(stageID))
	if err != nil || stage == nil {
		zap.L().Error("get task stage by id error", zap.Error(err))
		return nil, errs.GrpcError(model.TaskStageNotExistError)
	}

	// 3. 获取项目信息
	project, err := ts.project.GetProjectByID(ctx, projectID)
	if err != nil || project == nil {
		zap.L().Error("get project by id error", zap.Error(err))
		return nil, errs.GrpcError(model.ProjectNotExistError)
	}

	// 检查项目是否已删除
	if project.Deleted == model.Deleted {
		return nil, errs.GrpcError(model.ProjectDeletedError)
	}

	// 4. 获取任务表中id_num字段最大值
	maxIdNum, err := ts.task.GetMaxIdNumByProjectID(ctx, projectID)
	if err != nil {
		zap.L().Error("get max id_num error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMaxIdNumError)
	}

	// 5. 获取任务表中sort字段最大值
	maxSort, err := ts.task.GetMaxSortByProjectIDAndStageCode(ctx, projectID, int(stageID))
	if err != nil {
		zap.L().Error("get max sort error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMaxSortError)
	}

	// 6. 解析任务执行者
	var assignTo int64 = 0
	if req.AssignTo != "" {
		assignToStr, _ := encrypt.Decrypt(req.AssignTo, model.AESKey)
		assignTo, _ = strconv.ParseInt(assignToStr, 10, 64)
	}

	// 7. 组织Task信息
	now := time.Now().UnixMilli()
	newTask := &data.Task{
		ProjectCode:   projectID,
		Name:          req.Name,
		Pri:           0,                // 默认优先级
		ExecuteStatus: data.Wait,        // 默认等待状态
		Description:   "",               // 默认无描述
		CreateBy:      req.MemberId,     // 创建者为当前成员
		CreateTime:    now,              // 当前时间
		AssignTo:      assignTo,         // 指派给谁
		Deleted:       model.NotDeleted, // 默认未删除
		StageCode:     int(stageID),     // 所属阶段
		Sort:          maxSort + 1,      // 排序值为当前最大值+1
		IdNum:         maxIdNum + 1,     // 编号为当前最大值+1
		Private:       project.OpenTaskPrivate,
		BeginTime:     now,
		EndTime:       time.Now().Add(2 * 24 * time.Hour).UnixMilli(),
	}

	// 8. 开启事务，保存任务和任务成员关系
	err = ts.tran.ExecTran(func(db trans.DbConn) error {
		conn := db.(*gorm.MysqlConn)

		// 保存任务信息
		if err := ts.task.SaveTask(ctx, newTask, conn.TranDb); err != nil {
			zap.L().Error("save task error", zap.Error(err))
			return errs.GrpcError(model.SaveTaskError)
		}

		// 保存任务成员关系 - 创建者
		taskMember := &data.TaskMember{
			TaskCode:   newTask.Id,
			MemberCode: assignTo, // 当前成员
			JoinTime:   now,      // 当前时间
			IsOwner:    1,        // 是创建人
		}
		if assignTo == req.MemberId {
			taskMember.IsExecutor = model.Executor
		}

		if err := ts.taskMember.SaveTaskMember(ctx, taskMember, conn.TranDb); err != nil {
			zap.L().Error("save task member error", zap.Error(err))
			return errs.GrpcError(model.SaveTaskMemberError)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 9. 组织返回消息
	resp := &task.SaveTaskResp{}
	// 根据需要填充返回值
	resp.Id = newTask.Id
	resp.ProjectCode, _ = encrypt.EncryptInt64(newTask.ProjectCode, model.AESKey)
	resp.Name = newTask.Name
	resp.Pri = int32(newTask.Pri)
	resp.ExecuteStatus = newTask.GetExecuteStatusStr()
	resp.Description = newTask.Description
	resp.CreateBy, _ = encrypt.EncryptInt64(newTask.CreateBy, model.AESKey)
	resp.CreateTime = time_format.ConvertMsecToString(newTask.CreateTime)
	resp.AssignTo, _ = encrypt.EncryptInt64(newTask.AssignTo, model.AESKey)
	resp.Deleted = int32(newTask.Deleted)
	resp.StageCode, _ = encrypt.EncryptInt64(int64(newTask.StageCode), model.AESKey)
	resp.Done = int32(newTask.Done)
	resp.Sort = int32(newTask.Sort)
	resp.Private = int32(newTask.Private)
	resp.IdNum = int32(newTask.IdNum)
	resp.Code, _ = encrypt.EncryptInt64(newTask.Id, model.AESKey)

	// 如果有任务执行者，获取执行者信息
	if assignTo > 0 {
		memberResp, err := rpc.LoginServiceClient.GetMemberById(ctx, &login.GetMemberByIdReq{MemberId: assignTo})
		if err == nil && memberResp != nil {
			resp.Executor = &task.ExecutorMessage{
				Name:   memberResp.Name,
				Avatar: memberResp.Avatar,
				Code:   memberResp.Code,
			}
		}
	}

	return resp, nil
}

func (ts *TaskService) MoveTask(ctx context.Context, req *task.MoveTaskReq) (*task.MoveTaskResp, error) {
	// 解析请求消息
	orginTaskCodeStr, _ := encrypt.Decrypt(req.OriginTaskCode, model.AESKey)
	orginTaskId, _ := strconv.ParseInt(orginTaskCodeStr, 10, 64)

	targetStageCodeStr, _ := encrypt.Decrypt(req.TargetStageCode, model.AESKey)
	targetStageCode, _ := strconv.ParseInt(targetStageCodeStr, 10, 64)
	// 获取目标任务步骤信息
	targetStage, err := ts.taskStage.GetTaskStageByID(ctx, int(targetStageCode))
	if err != nil {
		zap.L().Error("get task stage by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskStageError)
	}
	// 获取原始任务信息
	originTask, err := ts.task.GetTaskById(ctx, orginTaskId)
	if err != nil {
		zap.L().Error("get task stage code error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskError)
	}

	// 在一个事务中完成移动任务的相关sql操作
	err = ts.tran.ExecTran(func(db trans.DbConn) error {
		conn := db.(*gorm.MysqlConn)
		var err error
		modifyStageCode := false
		// 如果目标stageCode与原始stageCode不相同，则需要更新任务的stageCode
		if int64(originTask.StageCode) != targetStageCode {
			// 更新原始任务的步骤id
			err = ts.task.ModifyStageCode(ctx, orginTaskId, int(targetStageCode), conn.TranDb)
			if err != nil {
				zap.L().Error("update task stage code error", zap.Error(err))
				return errs.GrpcError(model.ModifyStageCodeError)
			}
			modifyStageCode = true
		}
		if req.TargetTaskCode == "" {
			if !modifyStageCode {
				return nil
			}
			// 如果修改了任务所属步骤且没有目标任务，意味着移到了新步骤的末尾，需要修改原任务sort为新步骤的最大值+1
			maxSort, err := ts.task.GetMaxSortByProjectIDAndStageCode(ctx, targetStage.ProjectCode, int(targetStageCode))
			if err != nil {
				zap.L().Error("get max sort error", zap.Error(err))
				return errs.GrpcError(model.GetMaxSortError)
			}
			err = ts.task.ModifyTaskSort(ctx, orginTaskId, int32(maxSort+1), conn.TranDb)
			if err != nil {
				zap.L().Error("modify task sort error", zap.Error(err))
				return errs.GrpcError(model.ModifyTaskSortError)
			}
		} else {
			targetTaskCodeStr, _ := encrypt.Decrypt(req.TargetTaskCode, model.AESKey)
			targetTaskCode, _ := strconv.ParseInt(targetTaskCodeStr, 10, 64)
			// 获取目标任务信息
			targetTask, err := ts.task.GetTaskById(ctx, targetTaskCode)
			if err != nil {
				zap.L().Error("get task by id error", zap.Error(err))
				return errs.GrpcError(model.GetTaskError)
			}
			// 获取目标任务的sort
			targetTaskSort := targetTask.Sort
			// 获取目标任务所属项目id
			targetProjectId := targetTask.ProjectCode
			// 如果目标TaskCode不为空，则需要将目标任务及其之后的任务的sort加1
			err = ts.task.IncreaseSort(ctx, targetProjectId, int(targetStageCode), targetTaskSort, conn.TranDb)
			if err != nil {
				zap.L().Error("increase sort error", zap.Error(err))
				return errs.GrpcError(model.IncreaseSortError)
			}
			// 原任务的sort改为目标任务的sort
			err = ts.task.ModifyTaskSort(ctx, orginTaskId, int32(targetTaskSort), conn.TranDb)
			if err != nil {
				zap.L().Error("modify task sort error", zap.Error(err))
				return errs.GrpcError(model.ModifyTaskSortError)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &task.MoveTaskResp{}, nil
}
