package main

import (
	"ecommerceGO/config"
	"ecommerceGO/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatal("Error loading environment variables: ", err)
	}
	api.StartServer(cfg)
}
