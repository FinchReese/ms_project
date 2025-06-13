# 上传本地镜像到腾讯云私人仓库（需要关闭clash等代理工具）
1. 私人仓库需要先登录腾讯云容器镜像服务 Docker Registry
```shell
docker login hkccr.ccs.tencentyun.com --username=100013395802
```
2. 本地镜像关联远程仓库
```shell
docker tag project-user:v1.0.2 hkccr.ccs.tencentyun.com/ms-project/project-user:v1.0.2
```
```shell
docker tag project-project:v1.0.2 hkccr.ccs.tencentyun.com/ms-project/project-project:v1.0.2
```

3. 推送本地镜像到远程仓库
```shell
docker push hkccr.ccs.tencentyun.com/ms-project/project-user:v1.0.2
``` 
```shell
docker push hkccr.ccs.tencentyun.com/ms-project/project-project:v1.0.2
``` 