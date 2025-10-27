package shared

import "github.com/prometheus/client_golang/prometheus"

var HttpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests received",
	},
	[]string{"path", "method"},
)

func SetupPrometheus() {
	prometheus.MustRegister(HttpRequests)
}
