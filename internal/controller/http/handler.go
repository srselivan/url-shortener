package http

import "github.com/gorilla/mux"

type Service interface {
	Shorten(url string) (string, error)
	GetOriginal(url string) (string, error)
}

type Handler struct {
	s Service
}

func New(s Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", h.ShortenUrl()).Methods("POST")
	router.Handle("/{key}", h.GetOriginalUrl()).Methods("GET")

	return router
}
