package proxy

import (
	"net/http"
)

// HealthCheckHandler handles the health check requests
// @Summary Health check
// @Description Returns the health status of the server
// @Produce  plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}