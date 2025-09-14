# 第一阶段：编译阶段
FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 禁用 cgo，去掉调试符号
ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# 第二阶段：精简运行镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main ./main

EXPOSE 8080  # 根据你程序监听端口改

CMD ["./main"]