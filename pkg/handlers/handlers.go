package handlers

import "net/http"

// Login : provides handling logger
func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

	}
}
