package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// Serve starts the HTTP server for Prometheus scrapers to use.
func Serve(dummy bool, port int) {
	go func() {
		if !dummy {
			http.Handle("/metrics", promhttp.Handler())
			bindAddr := fmt.Sprintf(":%d", port)
			log.Infof("Prometheus metrics can be scraped from this host on port %d.", port)
			log.Error(http.ListenAndServe(bindAddr, nil))
		} else {
			log.Info("Metrics scraping disabled.")
		}
	}()
}
