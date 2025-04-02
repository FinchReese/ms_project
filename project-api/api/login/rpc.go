package login

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"test.com/project-api/config"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/user/login"
)

var LoginServiceClient login.LoginServiceClient

func InitUserRpc() {
	resolver.Register(service_discover.NewEtcdBuilder(config.AppConf.EtcdConf.Addrs))
	conn, err := grpc.Dial("etcd:///login", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
