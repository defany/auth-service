package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	RequestCounterName = "requests_total"
)

var (
	ErrRequestCounterIsNotDefined = errors.New("request counter is not defined. Maybe you forgot to define it")
)

type Metrics struct {
	appName string

	counters map[string]prometheus.Counter
}

func NewMetrics(appName string) *Metrics {
	return &Metrics{
		appName:  appName,
		counters: make(map[string]prometheus.Counter),
	}
}

func (m *Metrics) WithRequestCounter(namespace string) *Metrics {
	counter := NewCounter(namespace)

	counter.
		WithName(RequestCounterName).
		WithDescription("Requests total count to server").
		WithSubSystem("grpc")

	_, ok := m.counters[counter.name]
	if ok {
		return m
	}

	m.counters[counter.name] = promauto.NewCounter(counter.prometheusOptions(m.appName))

	return m
}

func (m *Metrics) IncRequestsCount() error {
	counter, ok := m.counters[RequestCounterName]
	if !ok {
		return ErrRequestCounterIsNotDefined
	}

	counter.Inc()

	return nil
}
