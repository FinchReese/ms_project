package project

import (
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"test.com/project-api/config"
	"test.com/project-common/service_discover"
	"test.com/project-grpc/account"
	"test.com/project-grpc/department"
	"test.com/project-grpc/menu"
	"test.com/project-grpc/project"
	"test.com/project-grpc/project_auth"
	"test.com/project-grpc/project_node"
	"test.com/project-grpc/task"
)

var projectServiceClient project.ProjectServiceClient
var TaskServiceClient task.TaskServiceClient
var AccountServiceClient account.AccountServiceClient
var DepartmentServiceClient department.DepartmentServiceClient
var ProjectAuthServiceClient project_auth.ProjectAuthServiceClient
var MenuServiceClient menu.MenuServiceClient
var ProjectNodeServiceClient project_node.ProjectNodeServiceClient

func InitProjectRpc() {
	resolver.Register(service_discover.NewEtcdBuilder(config.AppConf.EtcdConf.Addrs))
	conn, err := grpc.Dial(
		"etcd:///project",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	projectServiceClient = project.NewProjectServiceClient(conn)
	TaskServiceClient = task.NewTaskServiceClient(conn)
	AccountServiceClient = account.NewAccountServiceClient(conn)
	DepartmentServiceClient = department.NewDepartmentServiceClient(conn)
	ProjectAuthServiceClient = project_auth.NewProjectAuthServiceClient(conn)
	MenuServiceClient = menu.NewMenuServiceClient(conn)
	ProjectNodeServiceClient = project_node.NewProjectNodeServiceClient(conn)
}
