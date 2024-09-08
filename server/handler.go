package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bmonitor/database"
	"github.com/fermyon/spin/sdk/go/v2/variables"
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
	urlParam := r.URL.Query().Get("url")

	if err := validateUrl(urlParam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.New()
	defer db.Close()

	if err := database.InsertMetric(db, urlParam); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateUrl(urlParam string) error {
	if urlParam == "" {
		return errors.New("Bad param 'url'")
	}

	if _, err := url.ParseRequestURI(urlParam); err != nil {
		return errors.New(fmt.Sprintf("Invalid 'url' format '%s'", urlParam))
	}

	allowedOrigins, err := variables.Get("cors_allowed_origins")

	if err != nil {
		return err
	}

	if allowedOrigins == "*" {
		return nil
	}

	// NOTE: Assuming only one origin
	if !strings.Contains(urlParam, allowedOrigins) {
		return errors.New(fmt.Sprintf("Not allowed 'url' '%s'", urlParam))
	}

	return nil
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
