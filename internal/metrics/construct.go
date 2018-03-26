package metrics

import "github.com/prometheus/client_golang/prometheus"

type dummyGauge struct {
	prometheus.Gauge
}

func (d *dummyGauge) Inc() {}
func (d *dummyGauge) Dec() {}

type dummyHistogram struct {
	prometheus.Histogram
}

const defaultNamespace string = "nats_bridge"

func (d *dummyHistogram) Observe(_ float64) {}

func constructHistogram(dummy bool, name string, help string) prometheus.Histogram {
	if dummy {
		return &dummyHistogram{}
	}
	histogram := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:      name,
			Namespace: defaultNamespace,
			Help:      help,
		},
	)
	prometheus.MustRegister(histogram)
	return histogram
}

func constructGauge(dummy bool, name string, help string) prometheus.Gauge {
	if dummy {
		return &dummyGauge{}
	}
	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:      name,
			Namespace: defaultNamespace,
			Help:      help,
		},
	)
	prometheus.MustRegister(gauge)
	return gauge
}

// ConstructSubscribePerformanceHistogram returns a metric for instrumenting subscription performance.
// If dummy is true, the returned metric is a no-op object.
func ConstructSubscribePerformanceHistogram(dummy bool) prometheus.Histogram {
	return constructHistogram(dummy, "from_queue", "Queue subscription metrics")
}

// ConstructPublishPerformanceHistogram returns a metric for instrumenting publishing performance.
// If dummy is true, the returned metric is a no-op object.
func ConstructPublishPerformanceHistogram(dummy bool) prometheus.Histogram {
	return constructHistogram(dummy, "to_queue", "Queue publish metrics")
}

// ConstructSubscriberCountGauge returns a metric for tracking the number of subscribers.
// If dummy is true, the returned metric is a no-op object.
func ConstructSubscriberCountGauge(dummy bool) prometheus.Gauge {
	return constructGauge(dummy, "subscribers_total", "Number of subscribers registered")
}
