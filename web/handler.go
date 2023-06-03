package web

import (
	"encoding/json"
	"net/http"

	"github.com/cloudberrybot/deeplinks"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store deeplinks.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Route("/", func(r chi.Router) {
		h.Get("/v1/apps", h.AppsList())
	})

	return h
}

type Handler struct {
	*chi.Mux
	store deeplinks.Store
}

func (h *Handler) AppsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apps, err := h.store.Apps()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(apps); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
