ENV_FILE := .env

# Подгружаем переменные окружения из .env
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

MIGRATION_DIR :=./migrations
DB_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@lms_db:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

# Запуск docker-compose
up:
	docker-compose --env-file $(ENV_FILE) up --build -d

# Остановить и удалить контейнеры
down:
	docker-compose --env-file $(ENV_FILE) down

# Перезапуск
restart: down up

# Название новой миграции передаётся через `make create-migration name=init_course`
create-migration:
	@read -p "Enter name of migration: " migration \
	&& goose -dir $(MIGRATION_DIR) create $$migration sql

# Применить все миграции
migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up

# Откатить последнюю миграцию
migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down
