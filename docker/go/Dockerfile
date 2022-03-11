FROM golang:1.16.7-alpine3.14

COPY src /go/src

WORKDIR /go/src

RUN apk update
RUN apk add --no-cache git
RUN go get -u github.com/gin-gonic/gin
RUN go mod init mainModules
RUN go get github.com/jinzhu/gorm