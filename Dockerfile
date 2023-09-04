FROM golang:1.18.10 as builder
WORKDIR /go/src/github.com/stonebirdjx/topx
COPY . .

RUN bash build.sh

FROM debian:10.13
LABEL maintainer="stonebirdjx <1245863260@qq.com>"
RUN groupadd -g 1005 stonebirdjx \
    && useradd -u 1005 -g 1005 stonebirdjx \
    && sed -i 's#http://deb.debian.org#http://mirrors.aliyun.com#g' /etc/apt/sources.list \
    && sed -i 's#http://security.debian.org#http://mirrors.aliyun.com#g' /etc/apt/sources.list \
    && apt-get update \
    && apt install -y curl \
    && apt-get install -y mongo-tools \
    && apt install -y ffmpeg

USER stonebirdjx

WORKDIR /opt/tiger/stonebirdjx/app/topx/
COPY --from=builder /go/src/github.com/stonebirdjx/topx/output . 

