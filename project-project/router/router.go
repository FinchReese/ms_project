package router

import (
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	interceptor "test.com/project-common/Interceptor"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/account"
	"test.com/project-grpc/department"
	"test.com/project-grpc/menu"
	"test.com/project-grpc/project"
	"test.com/project-grpc/project_auth"
	"test.com/project-grpc/project_node"
	"test.com/project-grpc/task"
	"test.com/project-project/config"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/database/trans"
	"test.com/project-project/internal/domain"
	"test.com/project-project/pkg/service/account_service_v1"
	"test.com/project-project/pkg/service/department_service_v1"
	"test.com/project-project/pkg/service/menu_service_v1"
	"test.com/project-project/pkg/service/project_auth_service_v1"
	"test.com/project-project/pkg/service/project_service_v1"
	"test.com/project-project/pkg/service/task_service_v1"
)

const (
	grpcServiceTTL int64 = 10 // grpc服务租约TTL为10秒
	cacheExpire    int64 = 5  // 缓存过期时间5分钟
)

type Router interface {
	Register(r *gin.Engine)
}

var routers = []Router{}

func RegisterRouter(ro ...Router) {
	routers = append(routers, ro...)
}

func InitRouter(r *gin.Engine) {
	for _, router := range routers {
		router.Register(r)
	}
}

func RegisterGrpc() *grpc.Server {
	// 创建拦截器
	methodToConfigMap := interceptor.MethodToConfigMap{
		"/project.service.v1.ProjectService/GetProjectList": {
			Resp:   &project.GetProjectListResp{},
			Expire: time.Duration(cacheExpire) * time.Minute,
		},
		"/task.service.v1.TaskService/GetTasksByStageCode": {
			Resp:   &task.GetTasksByStageCodeResp{},
			Expire: time.Duration(cacheExpire) * time.Minute,
		},
	}
	interceptor := interceptor.NewServiceInterceptor(methodToConfigMap, dao.Rc)
	// 创建GRPC服务器时注册拦截器
	s := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			interceptor.Intercept,
		),
	))
	// 注册project服务
	projectService := project_service_v1.NewProjectService(
		dao.NewMenuDAO(),
		dao.NewProjectMemberDAO(),
		dao.NewProjectTemplateDAO(),
		dao.NewTemplateTaskStageDAO(),
		dao.NewProjectDAO(),
		dao.NewProjectCollectDao(),
		dao.NewTaskStageDAO(),
		domain.NewMenuDomain(dao.NewMenuDAO()),
		trans.NewTransaction())
	project.RegisterProjectServiceServer(s, projectService)
	// 注册task服务
	taskService := task_service_v1.NewTaskService(
		dao.NewTaskStageDAO(),
		dao.NewTaskDAO(),
		dao.NewTaskMemberDAO(),
		dao.NewProjectDAO(),
		dao.NewProjectLogDAO(),
		dao.NewTaskWorkTimeDAO(),
		dao.NewFileDao(),
		dao.NewSourceLinkDao(),
		trans.NewTransaction(),
		domain.NewUserDomain(),
		domain.NewTaskWorkTimeDomain(dao.NewTaskWorkTimeDAO(), domain.NewUserDomain()),
	)
	task.RegisterTaskServiceServer(s, taskService)
	// 注册account服务
	accountService := account_service_v1.NewAccountService(
		domain.NewMemberAccountDomain(dao.NewMemberAccountDAO(), domain.NewUserDomain(), domain.NewDepartmentDomain(dao.NewDepartmentDAO())),
		domain.NewProjectAuthDomain(dao.NewProjectAuthDAO(), dao.NewMemberAccountDAO(), domain.NewProjectAuthNodeDomain(dao.NewProjectAuthNodeDAO()), domain.NewTaskDomain(dao.NewTaskDAO()), domain.NewProjectMemberDomain(dao.NewProjectMemberDAO())),
		domain.NewUserDomain(),
	)
	account.RegisterAccountServiceServer(s, accountService)
	// 注册department服务
	departmentService := department_service_v1.NewDepartmentService(
		domain.NewDepartmentDomain(dao.NewDepartmentDAO()),
		domain.NewUserDomain(),
	)
	department.RegisterDepartmentServiceServer(s, departmentService)
	lis, err := net.Listen("tcp", config.AppConf.GrpcConf.ListenAddr)
	if err != nil {
		log.Println("cannot listen")
	}
	// 注册project_auth服务
	projectAuthService := project_auth_service_v1.NewProjectAuthService(
		domain.NewProjectAuthDomain(dao.NewProjectAuthDAO(), dao.NewMemberAccountDAO(), domain.NewProjectAuthNodeDomain(dao.NewProjectAuthNodeDAO()), domain.NewTaskDomain(dao.NewTaskDAO()), domain.NewProjectMemberDomain(dao.NewProjectMemberDAO())),
		domain.NewUserDomain(),
		domain.NewProjectNodeDomain(dao.NewProjectNodeDAO(), domain.NewProjectAuthNodeDomain(dao.NewProjectAuthNodeDAO())),
		domain.NewProjectAuthNodeDomain(dao.NewProjectAuthNodeDAO()),
		trans.NewTransaction(),
	)
	project_auth.RegisterProjectAuthServiceServer(s, projectAuthService)
	// 注册menu服务
	menuService := menu_service_v1.NewMenuService(
		domain.NewMenuDomain(dao.NewMenuDAO()),
	)
	menu.RegisterMenuServiceServer(s, menuService)
	// 注册project_node服务
	projectNodeService := project_service_v1.NewProjectNodeService(
		domain.NewProjectNodeDomain(dao.NewProjectNodeDAO(), domain.NewProjectAuthNodeDomain(dao.NewProjectAuthNodeDAO())),
	)
	project_node.RegisterProjectNodeServiceServer(s, projectNodeService)
	// 创建协程，让启动grpc服务器不会阻塞到其他模块工作
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}

func RegisterGrpcAddrConf() {
	service_discover.RegisterService(config.AppConf.EtcdConf.Addrs, config.AppConf.GrpcConf.Name, config.AppConf.GrpcConf.ConnectAddr, grpcServiceTTL)
}
