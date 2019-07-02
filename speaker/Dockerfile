FROM alpine:latest

ADD speaker /speaker
add etcdv3.conf /etcdv3.conf
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ["/speaker"]
