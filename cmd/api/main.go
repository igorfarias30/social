package main

import (
	"log"

	"github.com/igorfarias30/social/internal/db"
	"github.com/igorfarias30/social/internal/env"
	"github.com/igorfarias30/social/internal/store"
)

const version = "0.0.1"

func main() {
	config := config{
		address: env.GetString("Address", ":4000"),
		dbConfig: dbConfig{
			address:            env.GetString("DB_ADDRESS", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 25),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 25),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		environment:        env.GetString("ENVIRONMENT", "development"),

	}

	db, err := db.New(
		config.dbConfig.address,
		config.dbConfig.maxOpenConnections,
		config.dbConfig.maxIdleConnections,
		config.dbConfig.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established.")

	store := store.NewPostgresStorage(db)
	app := &application{
		config: config,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
