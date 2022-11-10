.PHONY: help

env:
	export PG_URL=postgres://postgres:postgrespw@localhost:5432/postgres

rollback:
    migrate -source db/migrations -database '$(PG_URL)?sslmode=disable' down

drop:
    migrate -source db/migrations -database '$(PG_URL)?sslmode=disable' drop

migrate-create:  ### create new migration
	migrate create -ext sql -dir db/migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path db/migrations -database 'postgres://postgres:postgrespw@localhost:5432/postgres?sslmode=disable' up
.PHONY: migrate-up

force: ### migration up
	migrate -path db/migrations -database '$(PG_URL)?sslmode=disable' force 20221025181800
.PHONY: migrate-up

run:
	go run cmd/main.go
.PHONY: go