# 新增GRPC服务流程
1、在微服务目录的api/protoc目录下，创建服务的proto文件，然后执行以下命令在gen目录下生成对应的go文件：
project服务命令：
```shell
protoc --go_out=../../../project-grpc/project --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/project --go-grpc_opt=paths=source_relative  project_service.proto
```
task服务命令：
```shell
protoc --go_out=../../../project-grpc/task --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/task --go-grpc_opt=paths=source_relative  task_service.proto
```
account服务命令：
```
protoc --go_out=../../../project-grpc/account --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/account --go-grpc_opt=paths=source_relative  account_service.proto
```

department服务命令：
```
protoc --go_out=../../../project-grpc/department --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/department --go-grpc_opt=paths=source_relative  department.proto
```

projectAuth服务命令
```
protoc --go_out=../../../project-grpc/project_auth --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/project_auth --go-grpc_opt=paths=source_relative  project_auth.proto
```

menu服务命令
```
protoc --go_out=../../../project-grpc/menu --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/menu --go-grpc_opt=paths=source_relative  menu.proto
```

project node服务
```
protoc --go_out=../../../project-grpc/project_node --go_opt=paths=source_relative --go-grpc_out=../../../project-grpc/project_node --go-grpc_opt=paths=source_relative  project_node.proto
```


2、把gen目录生成的go文件拷贝到project-grpc\project的对应目录下，修改两个go文件的package（不修改package为gen），后续GRPC服务器和客户端都使用这个目录下的go文件。<br/>
**补充**：为什么不直接把生成的go文件放在project-grpc\user目录下，而是先放在project-user\api\protoc\gen目录下，再拷贝？<br/>
避免生成的go文件直接覆盖旧版的go文件，如果新生成的go文件有问题，还能先用旧版的go文件