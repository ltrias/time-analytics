package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

	r.Get("/suggest", loadSuggest)

	r.Route("/events", func(r chi.Router) {
		r.Get("/", loadAllEvents)

		r.Route("/{eventID:\\d+}", func(r chi.Router) {
			r.Get("/", loadEvent)
		})
	})

	http.ListenAndServe(":8080", r)
}

func loadSuggest(w http.ResponseWriter, r *http.Request) {
	suggest := api.Suggest{}

	suggest.Type = repo.LoadTypeSuggest()
	suggest.Duration = repo.LoadDurationSuggest()
	suggest.Who = repo.LoadWhoSuggest()
	suggest.Subject = repo.LoadSubjectSuggest()

	respondWithJSON(w, http.StatusOK, suggest)
}

func loadAllEvents(w http.ResponseWriter, r *http.Request) {
	te := repo.LoadAllEvents()

	respondWithJSON(w, http.StatusOK, te)
}

func loadEvent(w http.ResponseWriter, r *http.Request) {

	if eventID := chi.URLParam(r, "eventID"); eventID != "" {
		iEventId, _ := strconv.Atoi(eventID)
		respondWithJSON(w, http.StatusOK, repo.LoadEvent(iEventId))
	} else {
		respondWithError(w, http.StatusBadRequest, "")
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
