package middleware

import "net/http"

// Middleware type allows chaining
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain chains multiple middleware functions
func Chain(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
