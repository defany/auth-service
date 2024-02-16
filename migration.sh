#!/bin/bash
source .env

export MIGRATION_DSN="host=postgres port=5432 dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATIONS_DIR}" postgres "${MIGRATION_DSN}" up -v