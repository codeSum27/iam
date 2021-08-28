server-run:
	docker-compose up -d
	go build -o iam main.go
	sleep 5
	docker exec -it -uroot mariadb-server mysql -uroot -pqwe1212 -e  "create database IF NOT EXISTS iam"
	./iam
server-stop:
	docker-compose down

docker-build:
	docker build . -t orgensthe/codesum-iam:latest

build:
	go generate -v ./...
	go mod tidy
	go test ./...
	go build -o iam main.go
