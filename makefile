migrate:
	docker compose run --rm migrate

up:
	docker compose up --build -d db redis
	docker compose run --rm migrate
	docker compose up --build -d app

down:
	docker compose down -v --remove-orphans

up-dev:
	docker compose up --build -d db redis
	docker compose run --rm migrate
	docker compose up --build -d app-dev

faker:
	docker compose run --rm app go run ./cmd/faker/main.go 5000

rebuild: down up

rebuild-dev: down up-dev
