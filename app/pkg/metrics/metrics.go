package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	RequestCounterName  = "requests_total"
	ResponseCounterName = "responses_total"
)

var (
	ErrRequestCounterIsNotDefined = errors.New("request counter is not defined. Maybe you forgot to define it")
	ErrInvalidRequestCounterType  = errors.New("failed to cast counter to prometheus.Counter")
	ErrInvalidResponseCounterType = errors.New("failed to cast counter to prometheus.CounterVec")
)

type Metrics struct {
	appName string

	counters map[string]any
}

func NewMetrics(appName string) *Metrics {
	return &Metrics{
		appName:  appName,
		counters: make(map[string]any),
	}
}

func (m *Metrics) WithRequestCounter(namespace string) *Metrics {
	counter := NewCounter(namespace)

	counter.
		WithName(RequestCounterName).
		WithDescription("Total requests count to server").
		WithSubSystem("grpc")

	_, ok := m.counters[counter.name]
	if ok {
		return m
	}

	m.counters[counter.name] = promauto.NewCounter(counter.prometheusOptions(m.appName))

	return m
}

func (m *Metrics) WithResponseCounter(namespace string) *Metrics {
	counter := NewCounter(namespace)

	counter.
		WithName(ResponseCounterName).
		WithDescription("Total responses count from server").
		WithSubSystem("grpc")

	_, ok := m.counters[counter.name]
	if ok {
		return m
	}

	m.counters[counter.name] = promauto.NewCounterVec(counter.prometheusOptions(m.appName), []string{"status", "method"})

	return m
}

func (m *Metrics) IncRequestsCount() error {
	c, ok := m.counters[RequestCounterName]
	if !ok {
		return ErrRequestCounterIsNotDefined
	}

	counter, ok := c.(prometheus.Counter)
	if !ok {
		return ErrInvalidRequestCounterType
	}

	counter.Inc()

	return nil
}

func (m *Metrics) IncResponsesCount(status string, method string) error {
	c, ok := m.counters[ResponseCounterName]
	if !ok {
		return ErrRequestCounterIsNotDefined
	}

	counter, ok := c.(*prometheus.CounterVec)
	if !ok {
		return ErrInvalidResponseCounterType
	}

	counter.WithLabelValues(status, method).Inc()

	return nil
}
