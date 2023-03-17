package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (h *Handler) ShortenUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		url, err := io.ReadAll(request.Body)
		if err != nil {
			log.Println(err)
			return
		}

		shortUrl, err := h.s.Shorten(string(url))
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, shortUrl)
	})
}
