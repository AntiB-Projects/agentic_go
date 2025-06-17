#!/bin/bash
#
# This script orchestrates the database setup by executing a series of SQL files.
# It reads its configuration from a .env file located in the parent directory.

set -e

# --- Configuration Loading ---
echo "➡️ Loading configuration..."
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ENV_FILE="$SCRIPT_DIR/../.env"

if [ ! -f "$ENV_FILE" ]; then
    echo "❌ Error: Configuration file not found at $ENV_FILE"
    exit 1
fi
export $(grep -v '^#' "$ENV_FILE" | xargs)
if [ -z "$DATABASE_NAME" ] || [ -z "$DATABASE_USER" ] || [ -z "$DATABASE_PASSWORD" ]; then
    echo "❌ Error: Required database variables not set in $ENV_FILE"
    exit 1
fi
echo "✅ Configuration loaded."

PG_SUPERUSER=$(whoami)
SQL_DIR="$SCRIPT_DIR/sql"

# --- Step 1: Install Dependencies (if needed) ---
echo "️️➡️ Checking dependencies..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    if ! brew list pgvector &>/dev/null; then
        brew install pgvector
        brew services restart postgresql@16 || brew services restart postgresql@15 || brew services restart postgresql@14 || brew services restart postgresql
    fi
fi
echo "✅ Dependencies are satisfied."

# --- Step 2: Create Database and User (as Superuser) ---
echo "➡️ Executing 01_setup_db_and_user.sql as superuser..."
psql -v ON_ERROR_STOP=1 --host="$DATABASE_HOST" --port="$DATABASE_PORT" --username="$PG_SUPERUSER" --dbname="postgres" \
    -v DB_NAME="$DATABASE_NAME" \
    -v DB_USER="$DATABASE_USER" \
    -v DB_PASS="$DATABASE_PASSWORD" \
    -f "$SQL_DIR/01_setup_db_and_user.sql"
echo "✅ Database and user created."

# --- Step 3: Enable Extension (as Superuser) ---
# THIS IS THE NEW STEP
echo "➡️ Enabling vector extension in '$DATABASE_NAME' as superuser..."
psql -v ON_ERROR_STOP=1 --host="$DATABASE_HOST" --port="$DATABASE_PORT" --username="$PG_SUPERUSER" --dbname="$DATABASE_NAME" \
    -c "CREATE EXTENSION IF NOT EXISTS vector;"
echo "✅ Extension enabled."

# --- Step 4: Create Schema (as Application User) ---
echo "➡️ Executing 02_create_schema.sql as application user..."
export PGPASSWORD="$DATABASE_PASSWORD"
psql --host="$DATABASE_HOST" --port="$DATABASE_PORT" --username="$DATABASE_USER" --dbname="$DATABASE_NAME" \
    -f "$SQL_DIR/02_create_schema.sql"
unset PGPASSWORD
echo "✅ Schema created successfully."

echo "🎉 Setup complete! Database '$DATABASE_NAME' is ready to use."