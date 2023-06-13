package middlewares

import (
	"freep.space/fsp/internals"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(internals.BLUE+"%s %s %s %d"+internals.RESET, r.URL.Path, r.Method, r.UserAgent(), r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
