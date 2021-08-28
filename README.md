# IAM
인증&amp;권한 관리 서버


## Prerequisite

1. golang installation (https://golang.org/doc/install)
2. set go path 
```
export GOPATH = ~/go
```

3. install Docker and Docker-compose

https://docs.docker.com/engine/install/ (docker)

https://docs.docker.com/compose/install/ (docker-compose)

## Local server run

1. clone into GOPATH
```
mkdir -p $GOPATH/src/github.com/codeSum27/iam
git clone https://github.com/codeSum27/iam  $GOPATH/src/github.com/codeSum27/iam
```

2. Run server
```
cd $GOPATH/src/github.com/codeSum27/iam
make server-run
```

## Local server stop
```
make server-stop
```
