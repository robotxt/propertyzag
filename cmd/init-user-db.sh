#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER zaguser;
    GRANT ALL PRIVILEGES ON DATABASE zagdb TO zaguser;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO zaguser;
EOSQL