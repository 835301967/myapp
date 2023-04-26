package apm

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type promMetrics struct {
	Labels   prometheus.Labels
	Duration time.Duration
}

func (s *Tracer) reportMetrics(labels map[string]string, duration time.Duration) *promMetrics {
	pls := prometheus.Labels{}
	for key, value := range labels {
		pls[key] = value
	}
	metrics := &promMetrics{Labels: pls, Duration: duration}
	return metrics
}

func (m *promMetrics) Histogram(entity *prometheus.HistogramVec) {
	if entity != nil {
		entity.With(m.Labels).Observe(m.Duration.Seconds())
	}
}

func (m *promMetrics) Summary(entity *prometheus.SummaryVec) {
	if entity != nil {
		entity.With(m.Labels).Observe(m.Duration.Seconds())
	}
}
