package main

import (
	"fmt"
	"log"
	"net"

	"hexes/server/pkg/api"
	"hexes/server/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	address := fmt.Sprintf(":%d", cfg.Server.Port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can't listen: %s", err)
	}

	log.Printf("starting server on '%s'", address)
	if err = api.NewServer().Serve(l); err != nil {
		log.Fatalf("server crashed: %s", err)
	}
}
