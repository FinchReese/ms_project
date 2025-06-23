# docker运行项目
## 1、在根目录下打开git bash，执行以下命令编译得到可执行文件
```shell
GOOS=linux GOARCH=amd64 go build -o project-api/target/project-api project-api/api_main.go
GOOS=linux GOARCH=amd64 go build -o project-project/target/project-project project-project/project_main.go
GOOS=linux GOARCH=amd64 go build -o project-user/target/project-user project-user/user_main.go
```
## 2、切换到project-api目录下，执行以下命令创建project-api镜像
```shell
docker build -t project-api:v1.0.3 .
```

## 3、切换到project-user目录下，执行以下命令创建project-user镜像
```shell
docker build -t project-user:v1.0.3 .
```

## 4、切换到project-project目录下，执行以下命令创建project-project镜像
```shell
docker build -t project-project:v1.0.3 .
```

## 5、回到根目录，执行以下命令启动容器
```shell
docker-compose up -d
```

## 6、单独调试镜像
调试project-api镜像的命令如下：
```shell
docker run -d -p 80:80 project-api:latest
```

调试project-user镜像的命令如下：
```shell
docker run -d -p 8080:8080 -p 8881:8881 project-user:latest
```

## 7、上传镜像到腾讯云私有仓库
### 登录
```shell
docker login hkccr.ccs.tencentyun.com --username=100013395802
```

关联远程仓库
```shell
docker tag project-user:v1.0.3 hkccr.ccs.tencentyun.com/ms-project/project-user:v1.0.3
docker tag project-api:v1.0.3 hkccr.ccs.tencentyun.com/ms-project/project-api:v1.0.3
docker tag project-project:v1.0.3 hkccr.ccs.tencentyun.com/ms-project/project-project:v1.0.3
```

推送远程仓库
```shell
docker push hkccr.ccs.tencentyun.com/ms-project/project-user:v1.0.3
docker push hkccr.ccs.tencentyun.com/ms-project/project-api:v1.0.3
docker push hkccr.ccs.tencentyun.com/ms-project/project-project:v1.0.3
```

# nacos上传配置文件
http://localhost:8848/nacos/index.html
账号和密码默认都是nacos

# 查看链路追踪访问
http://localhost:16686/search

# ES安装ik分词器命令：
```
docker-compose exec es elasticsearch-plugin install https://get.infini.cloud/elasticsearch/analysis-ik/8.6.0
```
