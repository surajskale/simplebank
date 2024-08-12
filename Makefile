postgres:
	docker run --name firstpostgresforgo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres@123 -d postgres:16-alpine

startrds: 
	docker container start firstpostgresforgo

createdb:
	docker exec -it firstpostgresforgo createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it firstpostgresforgo drop db simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:postgres@123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:postgres@123@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

runserver:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test runserver