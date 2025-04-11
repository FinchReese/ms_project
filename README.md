# docker运行项目
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
在连接mysql数据库，补充表结构定义

**员工表**

```sql
CREATE TABLE `ms_member`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '系统前台用户表',
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户登陆账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登陆密码',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态',
  `last_login_time` bigint(0) NULL DEFAULT NULL COMMENT '上次登录时间',
  `sex` tinyint(0) NULL DEFAULT 0 COMMENT '性别',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像',
  `idcard` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证',
  `province` int(0) NULL DEFAULT 0 COMMENT '省',
  `city` int(0) NULL DEFAULT 0 COMMENT '市',
  `area` int(0) NULL DEFAULT 0 COMMENT '区',
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在地址',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `dingtalk_openid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉openid',
  `dingtalk_unionid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉unionid',
  `dingtalk_userid` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '钉钉用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1000 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;
```
**组织表**

```sql
CREATE TABLE `ms_organization`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `avatar` varchar(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `member_id` bigint(0) NULL DEFAULT NULL COMMENT '拥有者',
  `create_time` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
  `personal` tinyint(1) NULL DEFAULT 0 COMMENT '是否个人项目',
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `province` int(0) NULL DEFAULT 0 COMMENT '省',
  `city` int(0) NULL DEFAULT 0 COMMENT '市',
  `area` int(0) NULL DEFAULT 0 COMMENT '区',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织表' ROW_FORMAT = COMPACT;
```

# 单独调试镜像
调试project-api镜像的命令如下：
```shell
docker run -d -p 80:80 project-api:latest
```

调试project-user镜像的命令如下：
```shell
docker run -d -p 8080:8080 -p 8881:8881 project-user:latest
```

