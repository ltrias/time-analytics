package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ltrias/time-analytics/api"
)

var repo *api.TimeEventRepository = api.NewTimeEventRepository()

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(100 * time.Millisecond))

	r.Route("/", func(r chi.Router) {
		r.Get("/", loadAllEvents)
	})

	http.ListenAndServe(":8080", r)
}

func loadAllEvents(w http.ResponseWriter, r *http.Request) {
	te := repo.LoadAllEvents()

	respondWithJSON(w, http.StatusOK, te)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
