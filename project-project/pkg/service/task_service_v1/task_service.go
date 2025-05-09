package task_service_v1

import (
	"context"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
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
	taskStage    repo.TaskStageRepo
	task         repo.TaskRepo
	taskMember   repo.TaskMemberRepo
	project      repo.ProjectRepo
	tran         *trans.TransactionImpl
	projectLog   repo.ProjectLogRepo
	taskWorkTime repo.TaskWorkTimeRepo
	file         repo.FileRepo
	sourceLink   repo.SourceLinkRepo
}

func NewTaskService(ts repo.TaskStageRepo, t repo.TaskRepo, tm repo.TaskMemberRepo, p repo.ProjectRepo, pl repo.ProjectLogRepo, twt repo.TaskWorkTimeRepo, f repo.FileRepo, sl repo.SourceLinkRepo, tran *trans.TransactionImpl) *TaskService {
	return &TaskService{
		taskStage:    ts,
		task:         t,
		taskMember:   tm,
		project:      p,
		tran:         tran,
		projectLog:   pl,
		taskWorkTime: twt,
		file:         f,
		sourceLink:   sl,
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
	// 记录日志
	projectLog := &data.ProjectLog{
		MemberCode:   req.MemberId,
		ToMemberCode: newTask.AssignTo,
		SourceCode:   newTask.Id,
		Content:      newTask.Name,
		Remark:       "创建任务",
		CreateTime:   time.Now().UnixMilli(),
		Type:         "create",
		ActionType:   "task",
		ProjectCode:  projectID,
		Icon:         "plus",
		IsComment:    model.NotCommentLog,
		IsRobot:      0,
	}
	err = ts.projectLog.CreateProjectLog(ctx, projectLog)
	if err != nil {
		zap.L().Error("create project log error", zap.Error(err))
		return nil, errs.GrpcError(model.CreateProjectLogError)
	}

	dispTask := newTask.ToDisplayTask()
	// 9. 组织返回消息
	resp := &task.SaveTaskResp{}
	copier.Copy(&resp, dispTask)

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

func (ts *TaskService) GetTaskList(ctx context.Context, req *task.GetTaskListReq) (*task.GetTaskListResp, error) {
	var err error
	var list []*data.Task
	var total int64
	switch req.TaskType {
	case model.TaskTypeAssignedTo:
		list, total, err = ts.task.GetTasksByAssignToAndDone(ctx, req.MemberId, int(req.Done), int(req.Page), int(req.PageSize))
	case model.TaskTypeCreatedBy:
		list, total, err = ts.task.GetTasksByCreateByAndDone(ctx, req.MemberId, int(req.Done), int(req.Page), int(req.PageSize))
	case model.TaskTypeInvolved:
		list, total, err = ts.task.GetTasksByMemberIdAndDone(ctx, req.MemberId, int(req.Done), int(req.Page), int(req.PageSize))
	default:
		zap.L().Error("invalid task type", zap.Int32("task_type", req.TaskType))
		return nil, errs.GrpcError(model.InvalidTaskType)
	}
	if err != nil {
		zap.L().Error("get task list error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskListError)
	}
	// 收集任务的项目id和成员id，统一查询项目和成员信息，避免重复查询，提升性能
	projectIdList := []int64{}
	memberIdList := []int64{}
	for _, task := range list {
		projectIdList = append(projectIdList, task.ProjectCode)
		memberIdList = append(memberIdList, task.CreateBy)
	}
	// 向login服务一次性查询所有成员信息，提升性能
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memberIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memIdToInfo[member.Id] = member
	}
	// 向一次性查询所有项目信息，提升性能
	projects, err := ts.project.GetProjectsByIds(ctx, projectIdList)
	if err != nil {
		zap.L().Error("get projects info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetProjectsInfoError)
	}
	projectIdToInfo := make(map[int64]*data.Project)
	for _, project := range projects {
		projectIdToInfo[project.Id] = project
	}
	var mtdList []*data.MyTaskDisplay
	for _, v := range list {
		memberMessage := memIdToInfo[v.AssignTo]
		name := memberMessage.Name
		avatar := memberMessage.Avatar
		mtd := v.ToMyTaskDisplay(projectIdToInfo[v.ProjectCode], name, avatar)
		mtdList = append(mtdList, mtd)
	}
	var myMsgs []*task.TaskMessage
	copier.Copy(&myMsgs, mtdList)
	return &task.GetTaskListResp{List: myMsgs, Total: total}, nil
}

func (ts *TaskService) GetTaskDetail(ctx context.Context, req *task.GetTaskDetailReq) (*task.GetTaskDetailResp, error) {
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	taskInfo, err := ts.task.GetTaskById(ctx, taskCode)
	if err != nil {
		zap.L().Error("get task by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskError)
	}
	dispTask := taskInfo.ToTaskDisplay()
	// 允许访问的条件是Private字段为1且在ms_task_member表中存在记录
	if taskInfo.Private == 1 {
		taskMemberList, err := ts.taskMember.FindTaskMembers(ctx, taskInfo.Id, req.MemberId)
		if err != nil {
			zap.L().Error("find task members error", zap.Error(err))
			return nil, errs.GrpcError(model.GetTaskMembersError)
		}
		if taskMemberList == nil || len(taskMemberList) == 0 {
			dispTask.CanRead = model.CanNotRead
		} else {
			dispTask.CanRead = model.CanRead
		}
	}
	// 在ms_project表查询项目名字
	project, err := ts.project.GetProjectByID(ctx, taskInfo.ProjectCode)
	if err != nil {
		zap.L().Error("get project by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetProjectError)
	}
	dispTask.ProjectName = project.Name
	// 在ms_task_stage表查询任务步骤名字
	taskStage, err := ts.taskStage.GetTaskStageByID(ctx, taskInfo.StageCode)
	if err != nil {
		zap.L().Error("get task stage by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskStageError)
	}
	dispTask.StageName = taskStage.Name
	// 借助login服务查找执行者信息
	member, err := rpc.LoginServiceClient.GetMemberById(ctx, &login.GetMemberByIdReq{MemberId: taskInfo.AssignTo})
	if err != nil {
		zap.L().Error("get member by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMemberByIdError)
	}
	dispTask.Executor.Name = member.Name
	dispTask.Executor.Avatar = member.Avatar
	// 组织回复消息
	resp := &task.GetTaskDetailResp{}
	copier.Copy(&resp, dispTask)
	return resp, nil
}

func (ts *TaskService) GetTaskMemberList(ctx context.Context, req *task.GetTaskMemberListReq) (*task.GetTaskMemberListResp, error) {
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	taskMemberList, total, err := ts.taskMember.GetTaskMemberList(ctx, taskCode, int(req.Page), int(req.PageSize))
	if err != nil {
		zap.L().Error("find task members error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskMembersError)
	}
	// 收集成员任务id，统一查找成员信息
	memberIdList := []int64{}
	for _, taskMember := range taskMemberList {
		memberIdList = append(memberIdList, taskMember.MemberCode)
	}
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memberIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memberIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memberIdToInfo[member.Id] = member
	}

	taskMemberMessageList := []*task.TaskMember{}

	for _, taskMember := range taskMemberList {
		code, _ := encrypt.EncryptInt64(taskMember.MemberCode, model.AESKey)
		taskMemberMessage := &task.TaskMember{
			Id:                taskMember.Id,
			Name:              memberIdToInfo[taskMember.MemberCode].Name,
			Avatar:            memberIdToInfo[taskMember.MemberCode].Avatar,
			Code:              code,
			MembarAccountCode: memberIdToInfo[taskMember.MemberCode].Account,
			IsExecutor:        int32(taskMember.IsExecutor),
			IsOwner:           int32(taskMember.IsOwner),
		}
		taskMemberMessageList = append(taskMemberMessageList, taskMemberMessage)
	}
	return &task.GetTaskMemberListResp{List: taskMemberMessageList, Total: total}, nil
}

func (ts *TaskService) GetTaskLogList(ctx context.Context, req *task.GetTaskLogListReq) (*task.GetTaskLogListResp, error) {
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	var logList []*data.ProjectLog
	var total int64
	var err error
	if req.All == 1 {
		logList, total, err = ts.projectLog.GetAllProjectLogListByTaskId(ctx, taskCode, int(req.Comment))
		if err != nil {
			zap.L().Error("get project log list by task id error", zap.Error(err))
			return nil, errs.GrpcError(model.GetProjectLogListError)
		}
	} else {
		logList, total, err = ts.projectLog.GetProjectLogListByTaskIdAndPage(ctx, taskCode, int(req.Comment), int64(req.Page), int64(req.PageSize))
		if err != nil {
			zap.L().Error("get project log list by task id and page error", zap.Error(err))
			return nil, errs.GrpcError(model.GetProjectLogListError)
		}
	}
	// 收集成员id,统一查找成员信息
	memberIdList := []int64{}
	for _, log := range logList {
		memberIdList = append(memberIdList, log.MemberCode)
	}
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memberIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memberIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memberIdToInfo[member.Id] = member
	}

	dispLogList := []*data.ProjectLogDisplay{}
	for _, log := range logList {
		dispLog := log.ToDisplay()
		memberInfo := memberIdToInfo[log.MemberCode]
		dispLog.Member.Name = memberInfo.Name
		dispLog.Member.Avatar = memberInfo.Avatar
		dispLog.Member.Id = log.MemberCode
		dispLog.Member.Code, _ = encrypt.EncryptInt64(log.MemberCode, model.AESKey)
		dispLogList = append(dispLogList, dispLog)
	}
	var taskLogList []*task.TaskLog
	copier.Copy(&taskLogList, dispLogList)
	return &task.GetTaskLogListResp{List: taskLogList, Total: total, Page: req.Page}, nil
}

func (ts *TaskService) SaveTaskWorkTime(ctx context.Context, req *task.SaveTaskWorkTimeReq) (*task.SaveTaskWorkTimeResp, error) {
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	// 组织taskWorkTime
	taskWorkTime := &data.TaskWorkTime{
		TaskCode:   taskCode,
		MemberCode: req.MemberId,
		Content:    req.Content,
		Num:        int(req.Num),
		BeginTime:  req.BeginTime,
	}
	err := ts.taskWorkTime.SaveTaskWorkTime(ctx, taskWorkTime)
	if err != nil {
		zap.L().Error("save task work time error", zap.Error(err))
		return nil, errs.GrpcError(model.SaveTaskWorkTimeError)
	}
	return &task.SaveTaskWorkTimeResp{}, nil
}

func (ts *TaskService) GetTaskWorkTimeList(ctx context.Context, req *task.GetTaskWorkTimeListReq) (*task.GetTaskWorkTimeListResp, error) {
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	taskWorkTimeList, err := ts.taskWorkTime.GetTaskWorkTimeList(ctx, taskCode)
	if err != nil {
		zap.L().Error("get task work time list error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskWorkTimeListError)
	}
	if len(taskWorkTimeList) == 0 {
		return &task.GetTaskWorkTimeListResp{List: []*task.TaskWorkTime{}}, nil
	}
	// 收集任务工时的成员id，统一查找成员信息
	memberIdList := []int64{}
	for _, taskWorkTime := range taskWorkTimeList {
		memberIdList = append(memberIdList, taskWorkTime.MemberCode)
	}
	members, err := rpc.LoginServiceClient.GetMembersByIds(ctx, &login.GetMembersByIdsReq{
		MemberIds: memberIdList,
	})
	if err != nil {
		zap.L().Error("get members info error", zap.Error(err))
		return nil, errs.GrpcError(model.GetMembersInfoError)
	}
	memberIdToInfo := make(map[int64]*login.MemberMessage)
	for _, member := range members.List {
		memberIdToInfo[member.Id] = member
	}

	taskWorkTimeDisplayList := []*data.TaskWorkTimeDisplay{}
	for _, taskWorkTime := range taskWorkTimeList {
		taskWorkTimeDisplay := taskWorkTime.ToDisplay()
		memberInfo := memberIdToInfo[taskWorkTime.MemberCode]
		taskWorkTimeDisplay.Member.Name = memberInfo.Name
		taskWorkTimeDisplay.Member.Avatar = memberInfo.Avatar
		taskWorkTimeDisplay.Member.Id = memberInfo.Id
		taskWorkTimeDisplay.Member.Code, _ = encrypt.EncryptInt64(memberInfo.Id, model.AESKey)
		taskWorkTimeDisplayList = append(taskWorkTimeDisplayList, taskWorkTimeDisplay)
	}
	// 组织回复消息
	var taskWorkTimeListResp []*task.TaskWorkTime
	copier.Copy(&taskWorkTimeListResp, taskWorkTimeDisplayList)
	return &task.GetTaskWorkTimeListResp{List: taskWorkTimeListResp}, nil
}

// SaveUploadFileInfo 保存上传文件信息
func (ts *TaskService) SaveUploadFileInfo(ctx context.Context, req *task.SaveUploadFileInfoReq) (*task.SaveUploadFileInfoResp, error) {
	// 1. 解析请求参数
	var projectID int64 = 0
	if req.ProjectCode != "" {
		projectCodeStr, _ := encrypt.Decrypt(req.ProjectCode, model.AESKey)
		projectID, _ = strconv.ParseInt(projectCodeStr, 10, 64)
	}

	var taskID int64 = 0
	if req.TaskCode != "" {
		taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
		taskID, _ = strconv.ParseInt(taskCodeStr, 10, 64)
	}
	// 根据memberId查询organization code，把第一个作为成员的organization code
	var organizationCode int64 = 0
	grpcResp, err := rpc.LoginServiceClient.GetOrganizationList(ctx, &login.GetOrganizationListReq{MemberId: req.MemberId})
	if err != nil {
		zap.L().Error("get organization err", zap.Error(err))
		return nil, errs.GrpcError(model.GetOrganizationError)
	}
	orgs := grpcResp.OrgList
	if len(orgs) > 0 { // 教程把成员的第一个组织作为当前组织，实际应该是在前端选择当前组织，后端记录
		organizationCode = orgs[0].Id
	}

	// 2. 创建文件记录
	now := time.Now().UnixMilli()
	fileRecord := &data.File{
		PathName:         req.PathName,
		Title:            req.FileName,
		Extension:        req.Extension,
		Size:             int(req.Size),
		ObjectType:       "", // 默认为任务类型
		OrganizationCode: organizationCode,
		TaskCode:         taskID,
		ProjectCode:      projectID,
		CreateBy:         req.MemberId,
		CreateTime:       now,
		Downloads:        0,                // 默认下载次数为0
		Extra:            "",               // 默认无额外信息
		Deleted:          model.NotDeleted, // 默认未删除
		FileUrl:          req.FileUrl,
		FileType:         req.FileType,
	}

	// 3. 开启事务，保存文件信息和资源关联关系
	err = ts.tran.ExecTran(func(db trans.DbConn) error {
		conn := db.(*gorm.MysqlConn)
		// 保存文件信息
		if err := ts.file.SaveFile(ctx, fileRecord, conn.TranDb); err != nil {
			zap.L().Error("save file error", zap.Error(err))
			return errs.GrpcError(model.SaveFileError)
		}
		// 保存源文件链接信息
		sourceLink := &data.SourceLink{
			SourceType:       "file",
			SourceCode:       int64(fileRecord.Id),
			LinkType:         model.LinkTypeTask,
			LinkCode:         taskID,
			OrganizationCode: organizationCode,
			CreateBy:         req.MemberId,
			CreateTime:       strconv.FormatInt(now, 10),
			Sort:             0, // 默认排序
		}

		if err := ts.sourceLink.SaveSourceLink(ctx, sourceLink, conn.TranDb); err != nil {
			zap.L().Error("save source link error", zap.Error(err))
			return errs.GrpcError(model.SaveSourceLinkError)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &task.SaveUploadFileInfoResp{}, nil
}

func (ts *TaskService) GetTaskLinkFiles(ctx context.Context, req *task.GetTaskLinkFilesReq) (*task.GetTaskLinkFilesResp, error) {
	// 解析task code
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskCode, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	// 根据task code获取资源关联列表
	sourceLinkList, err := ts.sourceLink.GetSourceLinkList(ctx, model.LinkTypeTask, taskCode)
	if err != nil {
		zap.L().Error("get source link list error", zap.Error(err))
		return nil, errs.GrpcError(model.GetSourceLinkListError)
	}
	// 先收集文件id，再统一查询文件信息，提升性能
	fileIdList := []int64{}
	for _, sourceLink := range sourceLinkList {
		fileIdList = append(fileIdList, sourceLink.SourceCode)
	}
	files, err := ts.file.GetFileListByIds(ctx, fileIdList)
	if err != nil {
		zap.L().Error("get file list by ids error", zap.Error(err))
		return nil, errs.GrpcError(model.GetFileListByIdsError)
	}
	// 创建文件id到文件的映射
	fileIdToFile := make(map[int64]*data.File)
	for _, file := range files {
		fileIdToFile[int64(file.Id)] = file
	}
	// 整理展示的文件信息
	var sourceLinkDisplayList []*data.SourceLinkDisplay
	for _, sourceLink := range sourceLinkList {
		sourceLinkDisplay := sourceLink.ToDisplay(fileIdToFile[sourceLink.SourceCode])
		sourceLinkDisplayList = append(sourceLinkDisplayList, sourceLinkDisplay)
	}
	// 组织回复消息
	var taskLinkFiles []*task.TaskLinkFile
	copier.Copy(&taskLinkFiles, sourceLinkDisplayList)
	return &task.GetTaskLinkFilesResp{List: taskLinkFiles}, nil
}

func (ts *TaskService) CreateTaskComment(ctx context.Context, req *task.CreateTaskCommentReq) (*task.CreateTaskCommentResp, error) {
	// 解析task code
	taskCodeStr, _ := encrypt.Decrypt(req.TaskCode, model.AESKey)
	taskId, _ := strconv.ParseInt(taskCodeStr, 10, 64)
	// 根据Id获取任务
	taskInfo, err := ts.task.GetTaskById(ctx, taskId)
	if err != nil {
		zap.L().Error("get task by id error", zap.Error(err))
		return nil, errs.GrpcError(model.GetTaskError)
	}
	pl := &data.ProjectLog{
		MemberCode:   req.MemberId,
		Content:      req.CommentContent,
		Remark:       req.CommentContent,
		Type:         "createComment",
		CreateTime:   time.Now().UnixMilli(),
		SourceCode:   taskId,
		ActionType:   "task",
		ToMemberCode: 0,
		IsComment:    model.CommentLog,
		ProjectCode:  taskInfo.ProjectCode,
		Icon:         "plus",
		IsRobot:      0,
	}
	err = ts.projectLog.CreateProjectLog(ctx, pl)
	if err != nil {
		zap.L().Error("create project log error", zap.Error(err))
		return nil, errs.GrpcError(model.CreateProjectLogError)
	}
	return &task.CreateTaskCommentResp{}, nil
}
