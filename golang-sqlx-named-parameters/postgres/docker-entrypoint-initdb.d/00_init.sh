#!/bin/bash
set -e
echo "---> executing 00_init.sh script begin"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE books (
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	title TEXT
);
EOSQL
echo "---> executing 00_init.sh script end"
