.PHONY: migrate-up migrate-down gen build local-env dev

migrate-up: build
	bin/migrate up

migrate-down: build
	bin/migrate down

gen:
	@sqlboiler --wipe psql

build:
	go build -o bin/migrate ./cmd/migrate
	go build -o bin/server ./cmd/server

local-env:
	docker-compose down
	docker-compose up -d

dev:
	go build -o bin/server ./cmd/server
	bin/server
