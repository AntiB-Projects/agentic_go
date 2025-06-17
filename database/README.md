# Database

## Create a new PSQL database

```bash
createdb ago_user_data
createuser ago_user_data
psql -U $USER -d ago_user_data
GRANT ALL PRIVILEGES ON DATABASE ago_user_data TO ago_user_data;
ALTER ROLE ago_user_data WITH LOGIN PASSWORD 'secure_password';
```