# 使用官方的 Go 基础镜像
FROM golang:1.23
WORKDIR /app
COPY ./target/project-api .
COPY ./config/app-docker.yaml .
RUN mkdir -p /logs/debug
RUN mkdir -p /logs/info
RUN mkdir -p /logs/error
RUN  mkdir config
RUN  mv app-docker.yaml ./config/app.yaml
RUN chmod +x project-api
EXPOSE 80
ENTRYPOINT ["./project-api"]