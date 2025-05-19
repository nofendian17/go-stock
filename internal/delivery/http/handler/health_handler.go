package handler

import (
	"net/http"
)

type HealthHandler interface {
	Healthz(w http.ResponseWriter, r *http.Request)
}
type healthHandler struct {
}

func NewHealthHandler() HealthHandler {
	return &healthHandler{}
}

func (s *healthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
