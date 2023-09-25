postgres:
	docker run --name postgres12 -p 127.0.0.1:15432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it dropdb simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb sqlc test