#!/bin/bash

ENV_FILE=".env"
DATABASE_URL=""
MIGRATE_COMMAND=migrate

# Check if the environment file exists
if [ -f "$ENV_FILE" ]; then
    # Read the value of DATABASE_URL from the environment file
    DATABASE_URL=$(grep -E "^DATABASE_URL=" "$ENV_FILE" | cut -d '=' -f2-)
else
    echo "Error: Environment file '$ENV_FILE' not found."
    exit 1
fi

if [ -f "./migrate" ]; then
    MIGRATE_COMMAND=./migrate
    echo ".migrate file exists, using migration locally."
fi


if [ "$1" = "create" ]; then
    echo "Creating migration"
    $MIGRATE_COMMAND create -ext sql -dir=database/migrations -seq init
elif [ "$1" = "run" ]; then 
    echo "Applying migration"
    $MIGRATE_COMMAND -path=database/migrations -database $DATABASE_URL -verbose up
fi
