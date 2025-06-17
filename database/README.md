# Database

## Create a new PSQL database

This assumes you have PostgreSQL installed and running.
You can create a new PostgreSQL database and user for your application using the following commands, either by pasting them into your terminal or running the setup script (`./setup_db.sh`).

```bash
createdb ago_user_data
createuser ago_user_data
psql -U $USER -d ago_user_data
GRANT ALL PRIVILEGES ON DATABASE ago_user_data TO ago_user_data;
ALTER ROLE ago_user_data WITH LOGIN PASSWORD 'secure_password';

-- Connect to your database: psql -U ago_user_data -d ago_user_data
# maybe you need to install pgvector extension
# e.g. by: brew install pgvector
# then you can login as a SU:
psql -U $USER -d ago_user_data
-- and run the following command to create the pgvector extension
CREATE EXTENSION IF NOT EXISTS vector;
```

