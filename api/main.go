package main

import (
	"log"

	"github.com/tuoitrevohoc/app-template/api/internal/server"
)

func main() {
	server, err := server.CreateServer()

	if err != nil {
		log.Fatal("Can't create server", err)
	}

	if err := server.Start(); err != nil {
		log.Fatal("Can't start server", err)
	}
}
