server-run:
	docker-compose up -d
	go build -o iam main.go
	sleep 5
	docker exec -it -uroot codesum-mariadb mysql -uroot -p"qwe1212!Q" -e  "create database IF NOT EXISTS iam"
	./iam
server-stop:
	docker-compose down

docker-build:
	docker build . -t codeSum27/iam:2.0.0

build:
	go generate -v ./...
	go mod tidy
	go test ./...
	go build -o iam main.go
