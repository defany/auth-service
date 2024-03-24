package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Counter struct {
	name        string
	description string

	namespace string
	subSystem string
}

func NewCounter(namespace string) *Counter {
	return &Counter{
		namespace: namespace,
	}
}

func (c *Counter) WithSubSystem(ss string) *Counter {
	c.subSystem = ss

	return c
}

func (c *Counter) WithName(name string) *Counter {
	c.name = name

	return c
}

// WithDescription - Добавляет help для описания того, за что отвечает counter
func (c *Counter) WithDescription(description string) *Counter {
	c.description = description

	return c
}

func (c *Counter) prometheusOptions(appName string) prometheus.CounterOpts {
	return prometheus.CounterOpts{
		Namespace: c.namespace,
		Subsystem: c.subSystem,
		Name:      fmt.Sprintf("%s_%s", appName, c.name),
		Help:      c.description,
	}
}
