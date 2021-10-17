FROM golang:alpine

# 启用 go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 切换工作目录
WORKDIR /app

# 将本地的文件都拷贝到 /app 目录下
COPY . .

# 指定OS等，并go build, ps: 由于使用的是 mac m1, 所以 GOARCH=arm64
RUN GOOS=linux GOARCH=arm64 go build -o http_server .

# 暴露容器端口 80
EXPOSE 80

CMD ["/app/http_server"]
