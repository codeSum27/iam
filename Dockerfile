FROM golang:1.18.0-alpine3.15

WORKDIR /go/src/github.com/codeSum27/iam
COPY . .
RUN go build -o iam main.go

