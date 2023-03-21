package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (h *Handler) ShortenUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			log.Printf("handle ShortenUrl: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		u, err := url.ParseRequestURI(string(body))
		if err != nil {
			log.Printf("handle ShortenUrl: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		shortUrl, err := h.s.Shorten(u.String())
		if err != nil {
			log.Printf("handle ShortenUrl: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, shortUrl)
	})
}
