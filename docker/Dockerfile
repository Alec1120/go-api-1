# docker/Dockerfile.dev
#### 使用 golang:1.22 作为基础镜像
FROM golang:1.22

#### 构建初始化环境
# 设置代理（可选）
# ENV GOPROXY=https://goproxy.cn,direct

# 设置 Go module mode
ENV GO111MODULE=on

# 构建时拷贝源码或者docker run时挂载源码
# COPY . .

# ✅ 进入工作目录
WORKDIR /app/go-api

# ✅ 自动初始化 Go module
RUN go mod init go-api || true

# ✅ 拉取依赖（如 gin）
RUN go get github.com/gin-gonic/gin

EXPOSE 8080

#### 运行阶段（调试时可以注释掉以下代码进入docker）
# ✅ 运行后端代码
CMD ["go", "run", "src/main.go"]