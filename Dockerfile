FROM golang:1.18

WORKDIR /go/src/github.com/ac2dia/iam
COPY . .
RUN go build -o iam main.go

