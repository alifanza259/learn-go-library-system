package main

import (
	"context"
	"log"
	"os"

	"github.com/alifanza259/learn-go-library-system/api"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
	}

	dbpool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatalf("cannot create database connection pool: %s", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	q := db.NewLibrary(dbpool)

	server, err := api.NewServer(q, config)
	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
