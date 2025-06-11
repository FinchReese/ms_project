# 使用官方的 Go 基础镜像
FROM golang:1.23
WORKDIR /app
COPY ./target/project-project .
COPY ./config/app-docker.yaml .
COPY ./config/nacos_config_docker.yaml .
RUN  mkdir config && mv app-docker.yaml config/app.yaml && mv nacos_config_docker.yaml config/nacos_config.yaml
RUN mkdir -p /logs/debug
RUN mkdir -p /logs/info
RUN mkdir -p /logs/error
EXPOSE 8082 8882
ENTRYPOINT ["./project-project"]