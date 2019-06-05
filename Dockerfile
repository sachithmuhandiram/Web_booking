FROM golang:alpine as goalpine 

RUN apk update

WORKDIR /go_web