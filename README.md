# 运行项目
在根目录下打开git bash，执行以下命令编译得到可执行文件
```shell
GOOS=linux GOARCH=amd64 go build -o project-api/target/project-api project-api/main.go
GOOS=linux GOARCH=amd64 go build -o project-user/target/project-user project-user/main.go
```
切换到project-api目录下，执行以下命令创建project-api镜像
```shell
docker build -t project-api:latest .
```

切换到project-user目录下，执行以下命令创建project-user镜像
```shell
docker build -t project-user:latest .
```

回到根目录，执行以下命令启动容器
```shell
docker-compose up -d
```
# 调试镜像
调试project-api镜像的命令如下：
```shell
docker run -d -p 80:80 project-api:latest
```

调试project-user镜像的命令如下：
```shell
docker run -d -p 8080:8080 -p 8881:8881 project-user:latest
```

# 新增GRPC服务流程
1、切换到project-user\api\protoc目录下，创建服务的proto文件，然后执行以下命令在gen目录下生成对应的go文件（将命令中的login_service.proto替换为新服务的proto文件名）：
```shell
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative  login_service.proto
```
2、把gen目录生成的go文件拷贝到project-grpc\user的对应目录下，后续GRPC服务器和客户端都使用这个目录下的go文件。<br/>
**补充**：为什么不直接把生成的go文件放在project-grpc\user目录下，而是先放在project-user\api\protoc\gen目录下，再拷贝？<br/>
避免生成的go文件直接覆盖旧版的go文件，如果新生成的go文件有问题，还能先用旧版的go文件
