FROM golang:latest AS builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=auto \
    && go env -w CGO_ENABLED=0 \
    && set -ex \
    && go build -ldflags "-s -w -extldflags '-static'" -o main

FROM alpine:latest

RUN apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone && \
    rm -rf /var/cache/apk/*

COPY --from=builder  /build/main /usr/bin/main

RUN chmod +x /usr/bin/main

WORKDIR /data

ENTRYPOINT [ "/usr/bin/main" ]