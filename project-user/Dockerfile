# 使用官方的 Go 基础镜像
FROM golang:1.23
WORKDIR /app
COPY ./target/project-user .
COPY ./config/app-docker.yaml .
RUN  mkdir config && mv app-docker.yaml config/app.yaml
RUN mkdir -p /logs/debug
RUN mkdir -p /logs/info
RUN mkdir -p /logs/error
EXPOSE 8080 8881
ENTRYPOINT ["./project-user"]