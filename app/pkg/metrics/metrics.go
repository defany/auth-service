package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	appName   string
	nameSpace string

	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

func Setup(_ context.Context, appName string, namespace string) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_requests_total",
				Help:      "Количество запросов к серверу",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_responses_total",
				Help:      "Количество ответов от сервера",
			},
			[]string{"status", "method"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_histogram_response_time_seconds",
				Help:      "Время ответа от сервера",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status", "method"},
		),
	}

	return nil
}

func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

func IncResponseCounter(status string, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

func HistogramResponseTimeObserve(status string, method string, t float64) {
	metrics.histogramResponseTime.WithLabelValues(status, method).Observe(t)
}
