.PHONY: docker-build docker-up docker-down migrate-up migrate-down migrate-create

docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-down:
	docker-compose down

migrate-up:
	migrate -path ./database/migrations -database "postgres://postgres:postgres@localhost:5432/go-auth?sslmode=disable" up

migrate-down:
	migrate -path ./database/migrations -database "postgres://postgres:postgres@localhost:5432/go-auth?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir ./database/migrations -seq <имя миграции>