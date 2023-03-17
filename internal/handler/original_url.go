package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handler) GetOriginalUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		key := mux.Vars(request)["key"]

		original, err := h.s.GetOriginal(key)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, original)
	})
}
