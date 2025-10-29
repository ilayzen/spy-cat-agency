#!/bin/bash
set -euo pipefail

# Usage: ./scripts/wait-for-postgres.sh <container_name> <path_to_migrations_directory>

if [ -z "${1:-}" ] || [ -z "${2:-}" ]; then
    echo "Usage: $0 <container_name> <path_to_migrations_directory>"
    exit 1
fi

CONTAINER_NAME="$1"
MIGRATIONS_DIR="$2"

PGUSER="${PGUSER:-postgres}"
PGPASSWORD="${PGPASSWORD:-postgres}"
PGDATABASE="${PGDATABASE:-postgres}"
PGPORT="${PGPORT:-5432}"
PGHOST="localhost"

# Check if container is running
if ! docker ps --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Container $CONTAINER_NAME is not running."
    exit 1
fi

echo "Waiting for PostgreSQL to be ready in container $CONTAINER_NAME..."

MAX_RETRIES=30
RETRY_INTERVAL=5
retry_count=0

# Wait for Postgres to respond
while ! docker exec -e PGPASSWORD="$PGPASSWORD" "$CONTAINER_NAME" \
    pg_isready -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" >/dev/null 2>&1; do
    retry_count=$((retry_count + 1))
    if [ $retry_count -ge $MAX_RETRIES ]; then
        echo "PostgreSQL in $CONTAINER_NAME is not ready after $MAX_RETRIES attempts. Exiting."
        exit 1
    fi
    echo "PostgreSQL in $CONTAINER_NAME is not ready yet. Retrying in $RETRY_INTERVAL seconds..."
    sleep $RETRY_INTERVAL
done

echo "PostgreSQL in $CONTAINER_NAME is ready!"

# Apply migrations
echo "Applying up migrations from directory: $MIGRATIONS_DIR"

if [ ! -d "$MIGRATIONS_DIR" ]; then
    echo "Directory $MIGRATIONS_DIR does not exist."
    exit 1
fi

shopt -s nullglob
migrations=( "$MIGRATIONS_DIR"/*up.sql )
IFS=$'\n' migrations=( $(printf "%s\n" "${migrations[@]}" | sort) )
unset IFS

if [ ${#migrations[@]} -eq 0 ]; then
    echo "No *up.sql migrations found in $MIGRATIONS_DIR."
    exit 0
fi

for migration in "${migrations[@]}"; do
    echo "Applying $migration..."
    if ! cat "$migration" | docker exec -i -e PGPASSWORD="$PGPASSWORD" "$CONTAINER_NAME" \
        psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" \
        -v ON_ERROR_STOP=1 -f -; then
        echo "Failed to apply $migration"
        exit 1
    fi
done

echo "All up migrations applied successfully!"
