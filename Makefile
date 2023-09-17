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

.PHONY: server sqlc