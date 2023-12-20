package main

import (
	"log"
	"news/internal/config"
	"news/internal/http-server/server"
	"news/internal/repository/postgres"
)

func main() {
	cfg := config.MustLoad()

	db := postgres.New(cfg.PostgresURL)

	srv := server.New(db, cfg)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
