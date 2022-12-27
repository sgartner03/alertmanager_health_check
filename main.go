package main

import (
	"alertmanager_health/metrics"
	"alertmanager_health/webserver"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	mux := http.NewServeMux()
	var metrics metrics.Metrics 
	metrics.SetCounterVec(
		"alertmanager_status",
		"The status of the alertmanager",
		"gepardec_cluster",
	)
	var web webserver.IncrementEndpoint
	web.Metrics = metrics
	
	mux.Handle("/inc", web)
	mux.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", mux)
}
