FROM scratch
LABEL MAINTAINER="Xijun Dai <daixijun1990@gmail.com>"

COPY nacos-initial /
ENTRYPOINT [ "/nacos-initial" ]
CMD [ "-h" ]
