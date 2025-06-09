#!/bin/bash

# Nacos完整部署脚本 - 使用独立MySQL

echo "开始部署Nacos - 使用独立MySQL..."

# 1. 部署Nacos专用MySQL
echo "1. 部署Nacos专用MySQL..."
kubectl apply -f nacos-mysql.yaml

# 等待MySQL Pod准备就绪
echo "等待MySQL Pod启动..."
kubectl wait --for=condition=ready pod -l app=nacos-mysql -n ms-project --timeout=300s

# 2. 部署Nacos服务
echo "2. 部署Nacos Service..."
kubectl apply -f nacos-service.yaml

# 3. 部署Nacos StatefulSet
echo "3. 部署Nacos StatefulSet..."
kubectl apply -f nacos-statefulset.yaml

# 等待Nacos Pod准备就绪
echo "等待Nacos Pod启动..."
kubectl wait --for=condition=ready pod -l app=nacos -n ms-project --timeout=300s

echo "Nacos部署完成！"

# 查看部署状态
echo "查看部署状态..."
kubectl get pods -n ms-project -l app=nacos
kubectl get pods -n ms-project -l app=nacos-mysql
kubectl get svc -n ms-project | grep -E "(nacos|nacos-mysql)"

# 获取Nacos访问信息
NACOS_CLUSTER_IP=$(kubectl get svc nacos -n ms-project -o jsonpath='{.spec.clusterIP}')
NACOS_PORT=$(kubectl get svc nacos -n ms-project -o jsonpath='{.spec.ports[0].port}')

echo "Nacos集群内访问地址: http://${NACOS_CLUSTER_IP}:${NACOS_PORT}/nacos"
echo "Nacos服务名访问: http://nacos.ms-project.svc.cluster.local:8848/nacos"
echo "默认用户名/密码: nacos/nacos"
echo ""
echo "外部访问方法:"
echo "1. 端口转发: kubectl port-forward --address 0.0.0.0 svc/nacos 8848:8848 -n ms-project"
echo "2. 然后访问: http://执行端口转发命令的节点ip:8848/nacos"
echo ""
echo "验证服务状态:"
echo "kubectl run test-pod --image=curlimages/curl -it --rm --restart=Never -n ms-project -- curl http://nacos.ms-project.svc.cluster.local:8848/nacos/actuator/health" 