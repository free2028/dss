# 使用多阶段构建，减少最终镜像大小
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dss ./cmd/dirstats

# 使用轻量级的 alpine 镜像作为最终镜像
FROM alpine:latest

# 安装 ca-certificates（如果需要 HTTPS 请求）
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/dss .

# 设置可执行权限
RUN chmod +x ./dss

# 设置入口点
ENTRYPOINT ["./dss"]
