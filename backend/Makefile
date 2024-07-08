test:
	go test ./...

get-u:
	go get -u ./...

build:
	docker compose build

run: build
	docker compose up

rund: build
	docker compose up -d

migrate-up:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema up

migrate-down:
	go run cmd/migrate/main.go -e .env -f cmd/migrate/schema down

gen-model: migrate-up
	rm -rf ./internal/entities/gorm && rm -rf ./internal/entities/model && go run gorm.io/gen/tools/gentool@latest -c ./gen.tool.yaml
