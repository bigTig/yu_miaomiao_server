# 声明镜像来源为golang:alpine
FROM golang:alpine

# 声明工作目录
WORKDIR /go/src/gin-web-server

# 拷贝整个server项目到工作目录
COPY . .

# go generate 编译前自动执行代码
# go env 查看go的环境变量
# go build -o server . 打包项目生成文件名为server的二进制文件
RUN go generate && go env && go build -o server .

# ==================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为alpine:latest
FROM alpine:latest

# 镜像编写者及邮箱
LABEL MAINTAINER="lkfeng2016@163.com"

# 声明工作目录
WORKDIR /go/src/gin-web-server

# 把/go/src/gin-vue-admin整个文件夹的文件到当前工作目录
COPY --from=0 /go/src/gin-web-server ./

EXPOSE 8888

# 运行打包好的二进制 并用-c 指定config.docker.yaml配置文件
ENTRYPOINT ./server -c config.docker.yaml
