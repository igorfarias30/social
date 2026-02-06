# Social App

## Build Application

1. Database
`$: docker compose up --build`

2. Go App
`$: go run cmd/api/*.go`

With air to use hot reload:
`$: air`

3. Run Migration
`$: migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users`

`$: migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up`