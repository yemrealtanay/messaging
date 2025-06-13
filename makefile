migrate:
	docker compose run --rm migrate

up:
	docker compose up --build -d db redis
	docker compose run --rm migrate
	docker compose up -d app

down:
	docker compose down -v --remove-orphans

rebuild: down up