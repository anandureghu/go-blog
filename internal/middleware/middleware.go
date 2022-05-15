package middleware

import (
	"net/http"
)

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Add("Content-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
