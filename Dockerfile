FROM golang:alpine as builder
MAINTAINER Li Kai Feng<lkfeng2016@163.com>
ENV VERSION 1.0.0

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server . \

# 移动到工作目录：/build
WORKDIR ./build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR ./dist

# 将二进制文件从 ./build 目录复制到这里
RUN cp ./build/app .

# 声明服务端口
EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml

# 启动容器时运行的命令
CMD ["./dist/app"]

#FROM alpine:latest
#
#LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"
#
#WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server
#
##COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/server ./
##COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/resource ./resource/
#COPY --from=0 /go/src/github.com/flipped-aurora/gin-vue-admin/server/config.docker.yaml ./
#
#EXPOSE 8888
#ENTRYPOINT ./server -c config.docker.yaml


