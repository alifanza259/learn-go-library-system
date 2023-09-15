package main

import (
	"log"

	"github.com/alifanza259/learn-go-library-system/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("cannot create server: %s", err)
	}

	err = server.Start("127.0.0.1:3000")
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
