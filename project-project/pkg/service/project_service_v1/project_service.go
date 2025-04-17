package project_service_v1

import (
	"context"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"test.com/project-common/encrypt"
	"test.com/project-common/errs"
	"test.com/project-common/time_format"
	"test.com/project-grpc/project"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/database/gorm"
	"test.com/project-project/internal/database/trans"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
	user_repo "test.com/project-user/pkg/repo"
)

const (
	getAllProjectTemplates    = -1
	getCustomProjectTemplates = 0
	getSystemProjectTemplates = 1
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	menuRepo              repo.MenuRepo
	projectMemberRepo     repo.ProjectMemberRepo
	projectTemplateRepo   repo.ProjectTemplateRepo
	templateTaskStageRepo repo.TemplateTaskStageRepo
	projectRepo           repo.ProjectRepo
	organizationRepo      user_repo.OrganizationRepo
	tran                  *trans.TransactionImpl
}

func NewProjectService(mr repo.MenuRepo, pmr repo.ProjectMemberRepo, ptr repo.ProjectTemplateRepo, ttsr repo.TemplateTaskStageRepo, pr repo.ProjectRepo,
	or user_repo.OrganizationRepo, t *trans.TransactionImpl) *ProjectService {
	return &ProjectService{
		menuRepo:              mr,
		projectMemberRepo:     pmr,
		projectTemplateRepo:   ptr,
		templateTaskStageRepo: ttsr,
		projectRepo:           pr,
		organizationRepo:      or,
		tran:                  t,
	}
}

func (p *ProjectService) Index(ctx context.Context, req *project.IndexMessage) (resp *project.IndexResponse, err error) {
	// 数据库查询所有菜单信息
	menuList, err := p.menuRepo.GetAllMenus(context.TODO())
	if err != nil {
		return nil, errs.GrpcError(model.GetAllMenusError)
	}
	menuTree := data.ConvertMenuListToTreeList(menuList)
	var menus []*project.MenuMessage
	copier.Copy(&menus, menuTree)
	return &project.IndexResponse{Menus: menus}, nil
}

func (p *ProjectService) GetProjectList(ctx context.Context, req *project.GetProjectListReq) (*project.GetProjectListResp, error) {
	projectList, total, err := p.projectMemberRepo.GetProjectList(ctx, req.GetMemberId(), req.GetSelectBy(), req.GetPage(), req.GetSize())
	if err != nil {
		return nil, errs.GrpcError(model.GetProjectListError)
	}
	resp := &project.GetProjectListResp{}
	copier.Copy(&resp.ProjectList, projectList)
	for _, project := range resp.ProjectList {
		project.OwnerName = req.MemberName
	}
	resp.Total = total
	return resp, nil
}

func (p *ProjectService) GetProjectTemplates(ctx context.Context, req *project.GetProjectTemplatesReq) (*project.GetProjectTemplatesResp, error) {
	var templates []data.ProjectTemplate
	var total int64
	var err error
	switch req.ViewType {
	case getAllProjectTemplates:
		templates, total, err = p.projectTemplateRepo.GetAllProjectTemplates(ctx, req.Page, req.PageSize)
	case getCustomProjectTemplates:
		templates, total, err = p.projectTemplateRepo.GetCustomProjectTemplates(ctx, req.MemberId, req.Page, req.PageSize)
	case getSystemProjectTemplates:
		templates, total, err = p.projectTemplateRepo.GetSystemProjectTemplates(ctx, req.Page, req.PageSize)
	default:
		return nil, errs.GrpcError(model.InvalidViewType)
	}
	if err != nil {
		return nil, errs.GrpcError(model.QueryProjectTemplateError)
	}

	taskStages, err := p.templateTaskStageRepo.GetTaskStagesByTemplateIds(ctx, data.ToProjectTemplateIds(templates))
	if err != nil {
		return nil, errs.GrpcError(model.QueryTemplateTaskStagesError)
	}
	// 获得模板id到该模板与任务步骤列表的映射
	templateId2TaskStages := data.CovertProjectMap(taskStages)
	completeTemplates := []*data.CompleteProjectTemplate{}

	for _, template := range templates {
		// 根据模板id获取任务步骤列表
		templateTaskStages := templateId2TaskStages[template.Id]
		// 转换为grpc回复消息的模板格式
		completeTemplates = append(completeTemplates, template.Convert(templateTaskStages))
	}

	projectTemplateMessages := []*project.ProjectTemplateMessage{}
	copier.Copy(&projectTemplateMessages, &completeTemplates)
	return &project.GetProjectTemplatesResp{Ptm: projectTemplateMessages, Total: total}, nil
}

func (p *ProjectService) SaveProject(ctx context.Context, req *project.SaveProjectReq) (*project.SaveProjectResp, error) {
	var organizationCode int64 = 0
	orgs, err := p.organizationRepo.GetOrganizationByMemberId(ctx, req.MemberId)
	if err != nil {
		zap.L().Error("get organization err", zap.Error(err))
		return nil, errs.GrpcError(model.GetOrganizationError)
	}
	if len(orgs) > 0 {
		// 教程默认把成员的第一个组织当做当前组织
		organizationCode = orgs[0].Id
	}
	templateCodeStr, _ := encrypt.Decrypt(req.TemplateCode, model.AESKey)
	templateCode, _ := strconv.ParseInt(templateCodeStr, 10, 64)
	pro := &data.Project{
		Name:              req.Name,
		Description:       req.Description,
		TemplateCode:      int(templateCode),
		CreateTime:        time.Now().UnixMilli(),
		Cover:             "https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500",
		Deleted:           model.NoDeleted,
		Archive:           model.NoArchive,
		OrganizationCode:  organizationCode,
		AccessControlType: model.Open,
		TaskBoardTheme:    model.Simple,
	}

	err = p.tran.ExecTran(func(dbConn trans.DbConn) error {
		// 存入项目表
		conn := dbConn.(*gorm.MysqlConn)
		err := p.projectRepo.SaveProject(ctx, pro, conn.TranDb)
		if err != nil {
			zap.L().Error("save project err", zap.Error(err))
			return errs.GrpcError(model.SaveProjectError)
		}
		pm := &data.ProjectMember{
			ProjectCode: pro.Id,
			MemberCode:  req.MemberId,
			JoinTime:    time.Now().UnixMilli(),
			IsOwner:     req.MemberId,
			Authorize:   "",
		}
		// 存入项目组织表
		err = p.projectMemberRepo.SaveProjectMember(ctx, pm, conn.TranDb)
		if err != nil {
			zap.L().Error("save project member err", zap.Error(err))
			return errs.GrpcError(model.SaveProjectMembertError)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	proCode, err := encrypt.EncryptInt64(pro.Id, model.AESKey)
	if err != nil {
		zap.L().Error("encrypt project id err", zap.Error(err))
		return nil, errs.GrpcError(model.EncryptProjectIdError)
	}
	organizationCodeStr, err := encrypt.EncryptInt64(organizationCode, model.AESKey)
	if err != nil {
		zap.L().Error("encrypt organization id err", zap.Error(err))
		return nil, errs.GrpcError(model.EncryptOrganizationIdError)
	}
	resp := &project.SaveProjectResp{
		Id:               pro.Id,
		Code:             proCode,
		OrganizationCode: organizationCodeStr,
		Name:             pro.Name,
		Cover:            pro.Cover,
		CreateTime:       time_format.ConvertMsecToString(pro.CreateTime),
		TaskBoardTheme:   pro.TaskBoardTheme,
	}
	return resp, nil
}
