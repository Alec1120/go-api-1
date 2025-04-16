# go-api-1
# 📝 Go Todo API (In-Memory Version) - 使用 Gin + Docker 构建的简易任务系统

A simple Todo REST API built with Go and Gin, storing tasks in memory and designed for containerized development via Docker.

---

## 📦 技术栈 | Tech Stack

- Go 1.22
- Gin - A high-performance HTTP web framework
- Docker - Development container environment

---

## 📁 项目结构 | Project Structure

go-api/ 
├── build
│ └── build.sh 
├── handlers/ # 路由处理模块 | Route handlers 
│ └── todos.go 
└── docker/
| └── Dockerfile # Dockerfile
├── main.go # 主程序入口 | Main entry point
├── run.sh
├── curls_cmd # test cases
└── README.md

### 🐳 构建镜像 | Build Docker Image

```bash
docker build -f docker/Dockerfile.dev -t go-api-dev .
```

## 🚀 运行 | Run
```bash
docker run -it --rm -v "$(pwd)/src:/app/go-api/src" -p 8080:8080 go-api-dev
```
