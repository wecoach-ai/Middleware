package middleware

import (
	"mime"
	"net/http"
)

func EnforceJSON(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		contentType := r.Header.Get("Content-Type")

		if contentType == "" {
			http.Error(w, "Content-Type not set. Update it to application/json", http.StatusBadRequest)
			return
		}

		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, "Malformed content-type", http.StatusBadRequest)
			return
		}

		if mt != "application/json" {
			http.Error(w, "Content-Type header must be application/json", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
