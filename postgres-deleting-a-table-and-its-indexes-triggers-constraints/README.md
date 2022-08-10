Code for [one of my blog posts](https://zauner.nllk.net/post/FIXME/).

# Details

* Execute `$ docker-compose up --build` to build the image and start the container.
* Then connect to the database using the following connection details:
  * host: `localhost`
	* port: `5432`
	* database: `postgres`
	* username: `postgres`
	* password: `postgrespwd`
* Finally test the behaviour when deleting tables in a PostgreSQL database
  using the code snippets below:

# Code snippets

## Creating the trigger function

```
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now(); 
   RETURN NEW;
END;
$$ language 'plpgsql';
```

## Creating the table and associated resources (indexes, ...)

```
CREATE TABLE users
(
    id          UUID PRIMARY KEY,
    nickname    TEXT NOT NULL,
    description TEXT, 
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'),
    CONSTRAINT  users_nickname_unique UNIQUE (nickname)
);

CREATE INDEX users_description_idx ON users (description);

CREATE TRIGGER users_updated_at_trigger
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
```

## Deleting the table

```
DROP TABLE users;
```

## Validating the result

### List relevant indexes

```
SELECT tablename, indexname, indexdef
FROM pg_indexes
WHERE schemaname = 'public'
ORDER BY tablename, indexname;
```

### List relevant triggers

```
SELECT event_object_table AS table_name ,trigger_name         
FROM information_schema.triggers  
GROUP BY table_name , trigger_name 
ORDER BY table_name ,trigger_name 
```

### List relevant constraints

```
SELECT * from pg_catalog.pg_constraint
WHERE conname like 'users_%';
```
