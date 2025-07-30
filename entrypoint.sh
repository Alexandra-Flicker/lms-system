#!/bin/sh
set -e

echo "Ожидание базы данных на ${DB_HOST}:${DB_PORT}..."

while ! nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 1
done

echo "База данных доступна. Выполняю миграции..."
goose -dir ./migrations postgres \
  "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" up

echo "Запускаю приложение..."
exec ./main
