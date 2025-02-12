build:
	go build -o rateme main.go

postgres:
	docker run --name postgres_rateme -p 1234:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

migrateup:
	 migrate -path db/migrations -database "postgres://root:secret@localhost:1234/postgres?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migrations -database "postgres://root:secret@localhost:1234/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate
