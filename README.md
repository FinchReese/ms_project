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
# 调试
调试project-api镜像的命令如下：
```shell
docker run -d -p 80:80 project-api:latest
```

调试project-user镜像的命令如下：
```shell
docker run -d -p 8080:8080 -p 8881:8881 project-user:latest
```