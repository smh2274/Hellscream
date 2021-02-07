# --- 构建编译环境 --
FROM golang:1.15 AS builder

# 设置环境变量
ENV GOPROXY https://goproxy.cn,direct
ENV GOPRIVATE github.com/smh2274/

WORKDIR /go/src/github.com/smh2274/Hellscream

# 拷贝需要编译的文件
COPY . /go/src/github.com/smh2274/Hellscream

# 编译
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOROOT_FINAL="$(pwd)" \
  go build -a -ldflags '-w -extldflags "-static"'  \
  -gcflags=-trimpath="$(pwd)" -asmflags=-trimpath="$(pwd)" cmd/hellscream.go

# --- 构建运行环境 ---
FROM envoyproxy/envoy-alpine:v1.17.0 AS prod

RUN mkdir -p /Azeroth/Hellscream/config \
    && mkdir -p /Azeroth/Hellscream/ssl \
    && mkdir -p /Azeroth/Hellscream/file \
    && mkdir -p /Azeroth/Hellscream/log \
    && touch /Azeroth/Hellscream/log/envoy_access.log \
    && chmod 777 /Azeroth/Hellscream/log/envoy_access.log

# 拷贝编译环境的二进制文件
COPY --from=builder /go/src/github.com/smh2274/Hellscream/hellscream /Azeroth/Hellscream/hellscream
COPY docker_prepare/* /Azeroth/Hellscream/
COPY internal/ssl/* /Azeroth/Hellscream/

RUN mv /Azeroth/Hellscream/envoy.yaml /Azeroth/Hellscream/config/ \
   && mv /Azeroth/Hellscream/hellscream_conf.yaml /Azeroth/Hellscream/config/ \
   && mv /Azeroth/Hellscream/domain.* /Azeroth/Hellscream/ssl/

# 设置时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
 && apk add --no-cache tzdata \
 && echo "Asia/Shanghai" > /etc/timezone \
 && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8808

CMD /Azeroth/Hellscream/run.sh