package main

import (
	"log"
	"url-shortener/config"
	"url-shortener/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
