package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/postgres"
	"url-shortener/internal/repository/pudge"
	"url-shortener/internal/service"
)

func main() {
	var usePostgres bool
	flag.BoolVar(&usePostgres, "d", false, "use postgres")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var (
		rMem *repository.InMemoryRepository
		rPg  *repository.PostgresRepository
		s    *service.Shortener
	)

	if usePostgres {
		dbPostgres, err := postgres.New(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}

		dbPudge, err := pudge.New("./storage/urls")
		if err != nil {
			log.Fatal(err)
		}

		rMem = repository.NewInMemoryRepository(dbPudge)
		rPg = repository.NewPostgresRepository(dbPostgres)

		s = service.New(rMem, rPg)

		defer dbPostgres.Close()
		defer dbPudge.Close()

	} else {
		dbPudge, err := pudge.New("./storage/urls")
		if err != nil {
			log.Fatal(err)
		}

		rMem = repository.NewInMemoryRepository(dbPudge)

		s = service.New(rMem)

		defer dbPudge.Close()
	}

	h := handler.New(s)

	router := mux.NewRouter()
	router.Handle("/", h.ShortenUrl()).Methods("POST")
	router.Handle("/{key}", h.GetOriginalUrl()).Methods("GET")

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
