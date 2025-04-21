添加一个新路由的流程（参考/project/login）：
# 一、注册路由
## 1、注册路由
在project-api\api目录的子目录的route.go注册URL和处理函数的映射关系

## 2、定义请求消息格式和回复消息格式
在project-api\pkg\model目录的子目录下定义请求消息格式和回复消息格式, 如果消息格式不是结构体类型则不定义

## 3、实现处理函数
在同目录下的handler.go实现对应的处理函数，处理函数借助对应的GRPC接口来实现功能。
借助commom.Result结构体组织回复消息。

# 二、GRPC接口
## 1、在proto文件新增rpc接口
proto文件路径在对应目录的\api\protoc

## 2、根据protoc文件生成go文件
在proto文件同级目录下，执行以下命令：
```
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative  project_service.proto
```
就会在gen目录下生成两个go文件，把2个go文件剪切到project-grpc目录的对应子目录下，然后修改两个go文件的package为所在目录名。

## 3、实现GRPC接口
在对应目录的service目录的子目录下，借助DAO层封装的数据库操作接口实现处理函数使用的grpc接口。
用到的枚举值和常量需要定义在对应目录的pkg\model\biz.go


# 三、DAO层
在对应目录的repo目录下先定义数据库操作接口，然后在dao目录下实现接口。
