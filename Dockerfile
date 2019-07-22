FROM quay.io/prometheus/busybox:latest
LABEL maintainer="gavin.zhou@gmail.com"

COPY janus /bin/janus

ENTRYPOINT [ "/bin/janus" ]
