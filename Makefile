DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

create-db:
	docker exec -it  postgres createdb --username=root --owner=root simple_bank

drop-db:
	docker exec -it  postgres dropdb simple_bank

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine

db-migrate:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

db-down:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down