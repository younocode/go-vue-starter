# Project github.com/younocode/go-vue-starter/server

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```


# database
- create migration
```shell
migrate create -ext sql -dir ./internal/database/migration -seq user_table
```
- migrate
```shell
migrate -database "postgres://user_dKwZY5:password_SxKEGb@pgsql.pi.local:5432/user_dKwZY5?sslmode=disable&search_path=public" -path ./internal/database/migration up
```
- sqlc generate
```shell
sqlc generate

```