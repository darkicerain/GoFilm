FROM golang:1.20

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

MAINTAINER "bai"
 
WORKDIR /opt/server
 
ADD ../../film/server /opt/server
 
RUN go mod tidy && go build main.go
 
EXPOSE 3061