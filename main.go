package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ltrias/time-analytics/api"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(100 * time.Millisecond))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		te := api.TimeEvent{time.Now(), "1:1", "Trias", 60, "Programming", "Devel", false}

		respondWithJSON(w, http.StatusOK, te)
	})

	http.ListenAndServe(":8080", r)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
