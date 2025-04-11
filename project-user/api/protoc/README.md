# 新增GRPC服务流程
1、在微服务目录的api/protoc目录下，创建服务的proto文件，然后执行以下命令在gen目录下生成对应的go文件（将命令中的login_service.proto替换为新服务的proto文件名）：
```shell
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative  login_service.proto
```
2、把gen目录生成的go文件拷贝到project-grpc\user的对应目录下，修改两个go文件的package（不修改package为gen），后续GRPC服务器和客户端都使用这个目录下的go文件。<br/>
**补充**：为什么不直接把生成的go文件放在project-grpc\user目录下，而是先放在project-user\api\protoc\gen目录下，再拷贝？<br/>
避免生成的go文件直接覆盖旧版的go文件，如果新生成的go文件有问题，还能先用旧版的go文件