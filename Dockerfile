FROM golang:1.17-buster as builder
ARG GOPROXY
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=${GOPROXY:-'https://proxy.golang.org,direct'} go build -ldflags "-s -w" -o nacos-initial main.go

FROM alpine:3.14
LABEL MAINTAINER="Xijun Dai <daixijun1990@gmail.com>"

COPY --from=builder /workspace/nacos-initial /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/nacos-initial" ]
CMD [ "-h" ]
