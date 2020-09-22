FROM golang:1.15.2-alpine3.12 
RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
    libressl \
    tar \
    git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

WORKDIR /root
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /root/librdkafka
RUN /root/librdkafka/configure
RUN make
RUN make install
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /go/src/kafka_producer
COPY . ./
COPY ca-certificates.crt /etc/ssl/certs/
RUN go build -tags dynamic -o ./server .
ENTRYPOINT ["./server"]