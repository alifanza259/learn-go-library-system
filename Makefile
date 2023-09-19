postgres-create: |
	docker run --name learn-go-library-system-postgres -d -p 5432\:5432 -e POSTGRES_PASSWORD=postgres postgres
	sleep 5; # Use timeout 5 if os=windows
	docker exec -it learn-go-library-system-postgres psql -U postgres -w -c "create database library"

postgres-destroy: |
	docker rm learn-go-library-system-postgres

migrateup:
	migrate -database ${DB_URL} -path db/migration up

migrateup1:
	migrate -database ${DB_URL} -path db/migration up 1

migratedown:
	migrate -database ${DB_URL} -path db/migration down

migratedown1:
	migrate -database ${DB_URL} -path db/migration down 1

sqlc:
	sqlc generate

server:
	go run main.go;

docker-build: |
	docker build -t my-golang-app .
	docker run --name my-golang-process -p 3000:3000 --network library-network my-golang-app

gomock:
	mockgen -destination db/mock/library.go github.com/alifanza259/learn-go-library-system/db/sqlc Library

.PHONY: server sqlc