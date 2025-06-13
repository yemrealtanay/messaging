up:
	docker compose up --build -d

down:
	docker compose down -v --remove-orphans

rebuild: down up
