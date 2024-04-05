package middlewares

import (
	"net/http"
)

func Security(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("UserAgent", "FreeP WebServer/1.0.0")
		next.ServeHTTP(w, r)
	})
}
