postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root bank_db

dropdb:
	docker exec -it postgres15 dropdb bank_db

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_db?sslmode=disable" -verbose down

.PHONY:postgres createdb dropdb migrateup migratedown


