#!/bin/sh
set -e

# Собираем строку подключения из переменных окружения
DB_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@lms_db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

echo "Running database migrations..."
goose -dir ./migrations postgres "$DB_URL" up

echo "Starting application..."
exec ./main
