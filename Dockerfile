FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /storyKitchen

COPY . .

RUN go build ./main.go

FROM alpine

ENV TZ=Asia/Shanghai\
    LANG=zh_CN.utf8

WORKDIR /storyKitchen

COPY --from=builder /storyKitchen .

EXPOSE 8001

CMD ["./main"]