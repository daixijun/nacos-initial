FROM golang:1.13 as builder

ENV GOARCH=amd64
ENV GOOS=linux
ENV GO11MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

WORKDIR /workspace
COPY . .
RUN go build -ldflags "-s -w" -o app main.go

FROM alpine:3.10
LABEL MAINTAINER="Xijun Dai <daixijun1990@gmail.com>"

COPY --from=builder /workspace/app /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/app" ]
CMD [ "-h" ]
