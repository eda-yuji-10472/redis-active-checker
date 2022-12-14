FROM golang:1.17-alpine

RUN apk update && apk add git

RUN go get github.com/gomodule/redigo/redis

ADD . /go/src/redis-active-checker
RUN go install github.com/eda-yuji-10472/redis-active-checker@main

ENV REDISHOST redis
ENV REDISPORT 6379

ENTRYPOINT /go/bin/redis-active-checker
