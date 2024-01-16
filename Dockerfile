FROM debian:12
LABEL MAINTAINER="Xijun Dai <daixijun1990@gmail.com>"
ENV TZ=Asia/Shanghai

RUN apt update -y && apt upgrade -y && \
    apt install tzdata ca-certificates -y && \
    apt clean && \
    apt autoremove -y && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone
COPY nacos-initial /usr/local/bin/
ENTRYPOINT [ "/usr/local/bin/nacos-initial" ]
CMD [ "--help" ]
