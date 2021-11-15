FROM alpine:3.14
LABEL MAINTAINER="Xijun Dai <daixijun1990@gmail.com>"

COPY nacos-initial /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/nacos-initial" ]
CMD [ "--help" ]
