package login

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"test.com/project-api/config"
	"test.com/project-common/service_discover"
	login_service_v1 "test.com/project-user/pkg/service/login.service.v1"
)

var LoginServiceClient login_service_v1.LoginServiceClient

func InitUserRpc() {
	resolver.Register(service_discover.NewEtcdBuilder(config.AppConf.EtcdConf.Addrs))
	conn, err := grpc.Dial("etcd:///login", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login_service_v1.NewLoginServiceClient(conn)
}
