package middlewares

import (
	"log"
	"net/http"

	"freep.space/fsp/internals/constant"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(constant.BLUE+"%s %s %s %d"+constant.RESET, r.URL.Path, r.Method, r.UserAgent(), r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
