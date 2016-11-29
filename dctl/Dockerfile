FROM alpine:latest
MAINTAINER zyfdegh <zyfdegg@gmail.com>

WORKDIR /root
# fix library dependencies
# otherwise golang binary may encounter 'not found' error
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY bin/dctl /root/dctl

VOLUME /var/run

CMD ["./dctl"]
