package router

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/user/login"
	"test.com/project-user/config"
	"test.com/project-user/pkg/dao"
	tran "test.com/project-user/pkg/database/trans"
	login_service_v1 "test.com/project-user/pkg/service/login.service.v1"
)

const (
	grpcServiceTTL int64 = 10 // grpc服务租约TTL为10秒
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
	s := grpc.NewServer()
	login.RegisterLoginServiceServer(s, &login_service_v1.LoginService{
		Cache:            dao.Rc,
		MemberRepo:       dao.MDao,
		OrganizationRepo: dao.ODao,
		Tran:             tran.NewTransaction(),
	})
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
