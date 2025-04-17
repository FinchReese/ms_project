package project

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"test.com/project-api/config"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/project"
)

var projectServiceClient project.ProjectServiceClient

func InitProjectRpc() {
	resolver.Register(service_discover.NewEtcdBuilder(config.AppConf.EtcdConf.Addrs))
	conn, err := grpc.Dial("etcd:///project", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	projectServiceClient = project.NewProjectServiceClient(conn)
}
