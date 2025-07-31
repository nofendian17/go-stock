package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next(w, r)
			log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
		}
	}
}
