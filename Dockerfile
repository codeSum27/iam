FROM golang:1.16

WORKDIR /go/src/github.com/codeSum27/iam
COPY . .
RUN go build -o iam main.go

