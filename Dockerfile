# 使用一个精简的基础镜像
FROM alpine:latest

# 安装运行二进制文件所需的任何依赖
RUN apk add --no-cache ca-certificates

# 设置工作目录
WORKDIR /app

# 将二进制文件从构建上下文复制到容器内
COPY /opt/uisApp /app/

# 使你的应用成为容器启动的默认程序
ENTRYPOINT ["./uisApp"]

# 暴露应用监听的端口
EXPOSE 8000