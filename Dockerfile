FROM golang:1.21.4-alpine as builder
WORKDIR /data/gin-otel-demo-code
ENV GOPROXY=https://goproxy.cn
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache upx ca-certificates tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o gin-otel-demo

FROM alpine:3.12 as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /data/gin-otel-demo-code/gin-otel-demo /gin-otel-demo
EXPOSE 9090
CMD ["/gin-otel-demo"]
