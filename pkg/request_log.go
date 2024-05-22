package middleware

import (
	"log/slog"
	"net/http"
)

func LogRequestData(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("URI:- " + r.RequestURI + " | " + "Method:- " + r.Method)

		next.ServeHTTP(w, r)
	})
}
