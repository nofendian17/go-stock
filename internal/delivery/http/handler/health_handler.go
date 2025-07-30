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

// Healthz is a health check endpoint
// @Summary Health check
// @Description Health check
// @Tags Health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthz [get]
func (s *healthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
