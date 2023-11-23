package main

import (
	"context"
	"log"
	"os"

	"github.com/alifanza259/learn-go-library-system/api"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/util"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Load Environment Variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
	}

	// Create database connection pool
	dbpool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatalf("cannot create database connection pool: %s", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// Sort of attaching model: available methods, tx, etc
	q := db.NewLibrary(dbpool)

	// Create server instance (gin)
	server, err := api.NewServer(q, config)
	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}

	// Start listening
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
