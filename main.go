package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alifanza259/learn-go-library-system/api"
	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/mail"
	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/alifanza259/learn-go-library-system/worker"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/hibiken/asynq"
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
	library := db.NewLibrary(dbpool)

	// Create asynq task distributor and processor
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(config, redisOpt, library)

	// Create server instance (gin)
	server, err := api.NewServer(library, config, taskDistributor)
	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}

	// Start listening
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, library db.Library) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, library, mailer)
	fmt.Println("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		fmt.Println("failed to start task processor")
	}
}
