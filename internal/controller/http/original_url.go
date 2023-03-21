package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener/internal/service/repository"
)

func (h *Handler) GetOriginalUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		key := mux.Vars(request)["key"]

		original, err := h.s.GetOriginal(key)
		if err != nil {
			if err == repository.ErrorNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				log.Printf("handle GetOriginalUrl: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, original)
	})
}
