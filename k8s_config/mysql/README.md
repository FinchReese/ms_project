# mysql导入表结构定义流程
1. 在本地mysql环境创建所有的表，根据需要插入初始数据，然后导出sql文件，这里我们命名为：msproject.sql
2. 在k8s集群的某个工作节点创建一个configmap存储表结构定义，这里命名为：mysql-structure
```shell
kubectl create configmap mysql-structure --from-file=msproject.sql -n ms-project
```

3. 在创建mysql pod的时候，将mysql-structure挂载到/docker-entrypoint-initdb.d目录下

补充：/docker-entrypoint-initdb.d 是 Docker 官方 MySQL 镜像中一个特殊目录，用于在容器首次启动时自动执行初始化脚本（如 SQL 文件、Shell 脚本等）。它的设计目的是为了方便用户在容器启动时自动完成数据库的初始化配置（如表结构创建、初始数据插入等）。

# 进入pod执行命令
```shell
kubectl exec -it mysql-master-0 -n ms-project -- bash
```