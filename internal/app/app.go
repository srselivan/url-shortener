package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"url-shortener/config"
	controller "url-shortener/internal/controller/http"
	"url-shortener/internal/service"
	"url-shortener/internal/service/repository"
	"url-shortener/pkg/postgres"
	"url-shortener/pkg/server"
)

const shutdownTimeout = 5 * time.Second

func Run(cfg *config.Config) {
	repos := make([]service.Repository, 0, 2)

	rMem := repository.NewInMemoryRepository()
	repos = append(repos, rMem)

	if cfg.Postgres.UsePostgres {
		dbPostgres, err := postgres.New(cfg.Postgres.URL)
		if err != nil {
			log.Fatalf("postgres db New(): %s", err)
		}
		defer dbPostgres.Close()

		rPg := repository.NewPostgresRepository(dbPostgres)
		repos = append(repos, rPg)
	}

	s := service.New(cfg.HTTP.Port, repos...)
	h := controller.New(s)

	srv := server.New(h.InitRoutes(), cfg)
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start server: %s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("stop server: %s", err)
	}
}
