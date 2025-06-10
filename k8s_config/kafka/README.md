# 测试kafka
## 1、创建主题test-topic
进入kafka脚本目录后，敲下面的命令：
```shell
./kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

## 2、查看主题列表
```shell
./kafka-topics.sh --list --bootstrap-server localhost:9092
```
