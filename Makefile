export PGUSER=dev
export PGPASSWORD=12345
export PGDATABASE=spy

up:
	chmod +x scripts/wait-for-postgres.sh

	docker compose up --build -d app-postgres
	./scripts/wait-for-postgres.sh app-postgres ./internal/repository/migrations/
	docker compose up --build -d app
