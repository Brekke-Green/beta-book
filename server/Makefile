postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root beta-book-db

dropdb:
	docker exec -it postgres dropdb beta-book-db

migrateup:
	 migrate -path db/migrations -database "postgresql://root:password@localhost:5433/beta-book-db?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migrations -database "postgresql://root:password@localhost:5433/beta-book-db?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown

