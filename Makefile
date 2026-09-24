.PHONY: postgres adminer migrate

up: 
	docker compose up -d

down: 
	docker compose down

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable up

