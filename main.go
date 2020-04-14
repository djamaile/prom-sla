package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Labels map[string]string

var (


	SERVICE_SLA_GAUGE = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "service_sla",
		Help:        "Exposes the SLA of the application",
		ConstLabels: prometheus.Labels{"application_name": "app1", "pod": "pod-1"},
	})

	SERVICE_SLA_GAUGE_2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "service_sla",
		Help:        "Exposes the SLA of the application",
		ConstLabels: prometheus.Labels{"application_name": "app2", "pod": "pod-1"},
	})

	SERVICE_SLA_GAUGE_3 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "service_sla",
		Help:        "Exposes the SLA of the application",
		ConstLabels: prometheus.Labels{"application_name": "app2", "pod": "pod-2"},
	})


)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(SERVICE_SLA_GAUGE)
	prometheus.MustRegister(SERVICE_SLA_GAUGE_2)
	prometheus.MustRegister(SERVICE_SLA_GAUGE_3)
}


func main() {
	SERVICE_SLA_GAUGE.Set(1)
	SERVICE_SLA_GAUGE_2.Set(2)
	SERVICE_SLA_GAUGE_3.Set(2)

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
