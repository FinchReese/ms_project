#!/bin/bash

# Nacos清理脚本 - 删除所有Nacos相关资源

echo "开始清理Nacos资源..."

# 1. 删除Nacos StatefulSet
echo "1. 删除Nacos StatefulSet..."
kubectl delete statefulset nacos -n ms-project

# 2. 删除Nacos Service
echo "2. 删除Nacos Service..."
kubectl delete -f nacos-service.yaml

# 3. 删除Nacos MySQL
echo "3. 删除Nacos MySQL..."
kubectl delete -f nacos-mysql.yaml

# 4. 删除PVC（可选）
echo "4. 清理PVC..."
read -p "是否删除数据存储卷（PVC）？这将永久删除数据 [y/N]: " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    kubectl delete pvc -l app=nacos -n ms-project
    kubectl delete pvc -l app=nacos-mysql -n ms-project
    echo "PVC已删除"
else
    echo "保留PVC数据"
fi

echo "Nacos清理完成！"

# 查看清理结果
echo "查看剩余资源..."
kubectl get pods -n ms-project | grep nacos || echo "无Nacos相关Pod"
kubectl get svc -n ms-project | grep nacos || echo "无Nacos相关Service"
kubectl get pvc -n ms-project | grep nacos || echo "无Nacos相关PVC" 