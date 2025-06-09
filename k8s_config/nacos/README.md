# mysql导入Nacos 表结构定义流程
1. 通过执行以下命令下载Nacos 表结构定义的sql文件（nacos-mysql.sql）
```shell
curl -o nacos-mysql.sql https://raw.githubusercontent.com/alibaba/nacos/2.2.0/distribution/conf/mysql-schema.sql
```
2. 在k8s集群的某个工作节点创建一个configmap存储表结构定义，这里命名为：mysql-structure
```shell
kubectl create configmap nacos-mysql-structure --from-file=nacos-mysql.sql -n ms-project
```

# 测试nacos
```shell
kubectl port-forward --address 0.0.0.0 svc/nacos-service -n ms-project 8848:8848
```