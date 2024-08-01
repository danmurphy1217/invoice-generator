package api

import (
	"encoding/json"
	"net/http"

	"github.com/danmurphy1217/invoice-generator/gen"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	/*
	health check handler func that our containers use to
	check if the service is live and available for traffic
	*/
    w.Header().Set("Content-Type", "application/json")

    resp := &gen.HealthCheckResponse{Healthy: true}
    json.NewEncoder(w).Encode(resp)
}