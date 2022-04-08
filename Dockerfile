FROM golang:1.18.0-alpine3.15

WORKDIR /go/src/github.com/ac2dia/iam
COPY . .
RUN go build -o iam main.go

