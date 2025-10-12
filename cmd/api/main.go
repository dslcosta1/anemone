package main

import (
	"log"
	"os"

	// "github.com/dslcosta1/social/internal/db"
	"github.com/dslcosta1/anemone/internal/env"
	// "github.com/dslcosta1/social/internal/env/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		/*
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},*/
	}

	/*db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool estabilished")

	store := store.NewStorage(db)
	*/
	app := &application{
		config: cfg,
		// store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
