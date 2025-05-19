# 使用官方的 Go 基础镜像
FROM golang:1.23
WORKDIR /app
COPY ./target/project-project .
COPY ./config/app-docker.yaml .
RUN  mkdir config && mv app-docker.yaml config/app.yaml
EXPOSE 8082 8882
ENTRYPOINT ["./project-project"]