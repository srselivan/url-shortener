package main

import (
	"flag"
	"log"
	"url-shortener/config"
	"url-shortener/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	flag.BoolVar(&cfg.Postgres.UsePostgres, "d", false, "use postgres")
	flag.Parse()

	app.Run(cfg)
}
