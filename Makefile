.PHONY: migration gen build local-env dev

migrate:
	go run cmd/migrate/main.go

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
