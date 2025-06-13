创建一个包含腾讯云 CCR 认证信息的 Secret
```shell
kubectl create secret docker-registry tencent-ccr-secret --docker-server=hkccr.ccs.tencentyun.com --docker-username=100013395802 --docker-password=Test_1234! --docker-email=nlp_irefe@163.com -n ms-project --dry-run=client -o yaml
```