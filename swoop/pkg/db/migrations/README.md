# Migrations

## Creating Migrations

Run from `pkg/db/migrations`

```

migrate create  -ext sql -dir . create_user_table
```

## Run Migrations to local PostgreSQL

Run from `pkg/db/migrations`

```
migrate --source file:// -database "postgres://postgres:password1\!@localhost:5432/byrdi?sslmode=disable" up
```

## Run Migrations to CloudSQL PostgreSQL though CloudSQL Proxy

Run from `pkg/db/migrations`

```
migrate --source file:// -database "postgres://postgres:Byrdi123\!@127.0.0.1:5432/byrdi?sslmode=disable" up
```

## Fixing Dirty Migrations

```
error: Dirty database version 20220129043912. Fix and force version.
```

Grab the previous version from the migration before `20220129043912` and run the following command:

```
migrate --source file:// -database "postgres://postgres:Byrdi123\!@127.0.0.1:5432/byrdi?sslmode=disable" force $PREVIOUS_VERSION
```
