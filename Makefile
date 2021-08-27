run:
	docker-compose up -d
	go build -o iam main.go
	sleep 5
	docker exec -it -uroot mariadb-server mysql -uroot -pqwe1212 -e  "create database IF NOT EXISTS iam"
	./iam
stop:
	docker-compose down
