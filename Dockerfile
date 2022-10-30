#FROM golang:alpine as builder
#MAINTAINER Li Kai Feng<lkfeng2016@163.com>
#ENV VERSION 1.0.0
#
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOPROXY=https://goproxy.cn,direct \
#    GOOS=linux \
#    GOARCH=amd64
#
##RUN go env -w GO111MODULE=on \
##    && go env -w GOPROXY=https://goproxy.cn,direct \
##    && go env -w CGO_ENABLED=0 \
##    && go env \
##    && go mod tidy \
##    && go build -o server . \
#
## 移动到工作目录：/build
#WORKDIR ./build
#
## 将代码复制到容器中
#COPY . .
#
## 将我们的代码编译成二进制可执行文件app
#RUN go build -o app .
#
## 移动到用于存放生成的二进制文件的 /dist 目录
#WORKDIR ./dist
#
## 将二进制文件从 ./build 目录复制到这里
#RUN cp ./build/app .
#
## 声明服务端口
#EXPOSE 8888
#ENTRYPOINT ./server -c config.docker.yaml
#
## 启动容器时运行的命令
#CMD ["./dist/app"]

# 声明镜像来源为golang:alpine
FROM golang:alpine

# 声明工作目录
WORKDIR /go/src/yu-miaomiao-service

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
WORKDIR /go/src/yu-miaomiao-service

# 把/go/src/gin-vue-admin整个文件夹的文件到当前工作目录
COPY --from=0 /go/src/yu-miaomiao-service ./

EXPOSE 8888

# 运行打包好的二进制 并用-c 指定config.docker.yaml配置文件
ENTRYPOINT ./server -c config.docker.yaml




