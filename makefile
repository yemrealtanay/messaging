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

rebuild: down up-dev