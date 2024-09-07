package server

import (
	"fmt"
	"net/http"

	"github.com/fermyon/spin/sdk/go/v2/variables"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins, err := variables.Get("cors_allowed_origins")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler, private bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyVar := "public_api_key"

		if private {
			keyVar = "private_api_key"
		}

		key, err := variables.Get(keyVar)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if key != r.Header.Get("X-API-Key") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println(r.Header.Get("Origin"))

		next.ServeHTTP(w, r)
	})
}
