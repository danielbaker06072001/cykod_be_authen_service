# Makefile

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=Cel-365.
DB_NAME=postgres
DB_PORT=5432
MIGRATION_PATH=./sql/
CONTAINER_NAME=cykod_postgres

postgres:
	docker run --name $(CONTAINER_NAME) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:12-alpine

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb $(DB_NAME)

migrateup:
	migrate -path ./sql/ -database "postgresql://postgres:Cel-365.@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path ./sql/ -database "postgresql://postgres:Cel-365.@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
