package router

import (
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	interceptor "test.com/project-common/Interceptor"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/project"
	"test.com/project-grpc/task"
	"test.com/project-project/config"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/database/trans"
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
	}
	interceptor := interceptor.NewServiceInterceptor(methodToConfigMap, dao.Rc)
	// 创建GRPC服务器时注册拦截器
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Intercept))
	// 注册project服务
	projectService := project_service_v1.NewProjectService(
		dao.NewMenuDAO(),
		dao.NewProjectMemberDAO(),
		dao.NewProjectTemplateDAO(),
		dao.NewTemplateTaskStageDAO(),
		dao.NewProjectDAO(),
		dao.NewProjectCollectDao(),
		dao.NewTaskStageDAO(),
		trans.NewTransaction())
	project.RegisterProjectServiceServer(s, projectService)
	// 注册task服务
	taskService := task_service_v1.NewTaskService(
		dao.NewTaskStageDAO(),
		dao.NewTaskDAO(),
		dao.NewTaskMemberDAO(),
		dao.NewProjectDAO(),
		trans.NewTransaction(),
	)
	task.RegisterTaskServiceServer(s, taskService)
	lis, err := net.Listen("tcp", config.AppConf.GrpcConf.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
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
	service_discover.RegisterService(config.AppConf.EtcdConf.Addrs, config.AppConf.GrpcConf.Name, config.AppConf.GrpcConf.Addr, grpcServiceTTL)
}
