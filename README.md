# go-gin-project

> 一个最简可运行、具备完整目录的 Gin REST API 示例。

## 启动
```bash
# 安装依赖
$ go mod tidy

# 本地运行
$ go run cmd/main.go

# 访问
curl http://localhost:8080/api/v1/ping # => {"message":"pong"}
```

## Docker
```bash
$ docker build -t go-gin-project .
$ docker run -p 8080:8080 go-gin-project
```