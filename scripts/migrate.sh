#!/bin/bash

# Read environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

DB_USER=${DB_USER}
DB_PASS=${DB_PASS}
DB_HOST=${DB_HOST}
DB_PORT=${DB_PORT}
DB_NAME=${DB_NAME}

DB_URL="postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
MIGRATIONS_DIR="./migrations"

# Arguments: <command> [<version>]
COMMAND=$1
VERSION=$2

# Display usage
usage() {
    echo "Usage: $0 <command> [<version>]"
    echo "Commands:"
    echo "  up [<version>]     - Apply migrations (up to a specific version)"
    echo "  down [<version>]   - Rollback migrations (down to a specific version)"
    echo "  create <name>      - Create a new migration with the specified name"
    echo "  version            - Print the current migration version"
    echo "  force <version>    - Set the migration version without running migrations"
}

# Run migrate command
run_migrate() {
    if [ -n "$VERSION" ]; then
        migrate -database "${DB_URL}" -path ${MIGRATIONS_DIR} $1 $VERSION
    else
        migrate -database "${DB_URL}" -path ${MIGRATIONS_DIR} $1
    fi
}

# Check command
case $COMMAND in
    up)
        run_migrate "up"
        ;;
    down)
        run_migrate "down"
        ;;
    create)
        if [ -z "$VERSION" ]; then
            echo "Error: Migration name is required for create command."
            usage
            exit 1
        fi
        migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $VERSION
        ;;
    version)
        migrate -database "${DB_URL}" -path ${MIGRATIONS_DIR} version
        ;;
    force)
        if [ -z "$VERSION" ]; then
            echo "Error: Version is required for force command."
            usage
            exit 1
        fi
        migrate -database "${DB_URL}" -path ${MIGRATIONS_DIR} force $VERSION
        ;;
    *)
        usage
        exit 1
        ;;
esac

# Check if migration command failed
if [ $? -ne 0 ]; then
    echo "Error: Migration command failed."
    exit 1
fi

echo "Migration command completed successfully."
