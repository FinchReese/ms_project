创建 MinIO Secret
```shell
kubectl create secret generic minio-secret -n ms-project \
  --from-literal=root-user=admin \
  --from-literal=root-password=password123
```