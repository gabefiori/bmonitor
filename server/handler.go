package server

import (
	"encoding/json"
	"net/http"

	"github.com/bmonitor/database"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var h http.Handler

	switch r.Method {
	case http.MethodOptions:
		h = corsMiddleware(http.HandlerFunc(handleOptions))
	case http.MethodPost:
		h = withDefaultMiddlewares(http.HandlerFunc(handleInsertMetric), false)
	case http.MethodGet:
		h = withDefaultMiddlewares(http.HandlerFunc(handleRetriveMetrics), true)
	default:
		http.Error(w, "Invalid METHOD", http.StatusBadRequest)
	}

	h.ServeHTTP(w, r)
}

func withDefaultMiddlewares(h http.Handler, private bool) http.Handler {
	return corsMiddleware(authMiddleware(h, private))
}

func handleInsertMetric(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	if url == "" {
		http.Error(w, "Bad param 'url'", http.StatusBadRequest)
		return
	}

	db := database.New()
	defer db.Close()

	if err := database.InsertMetric(db, url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleRetriveMetrics(w http.ResponseWriter, r *http.Request) {
	db := database.New()
	defer db.Close()

	pm, err := database.RetriveMetrics(db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pm)
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
