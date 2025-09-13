# 使用官方的golang镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go模块文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用alpine作为运行环境
FROM alpine:latest

# 安装ca证书
RUN apk --no-cache add ca-certificates

# 创建工作目录
WORKDIR /root/

# 从构建环境复制可执行文件
COPY --from=builder /app/main .

# 复制配置文件
COPY --from=builder /app/.env .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]