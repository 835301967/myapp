package apm

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (s *Span) ReportAndEnd(histogramEntity *prometheus.HistogramVec, summaryEntity *prometheus.SummaryVec, labels map[string]string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ended() {
		return
	}
	if s.Duration < 0 {
		s.Duration = time.Since(s.timestamp)
	}
	if s.dropped() {
		if s.tx == nil {
			droppedSpanDataPool.Put(s.SpanData)
		} else {
			duration := s.reportSelfTime()
			reportMetrics := s.tracer.reportMetrics(labels, duration)
			if histogramEntity != nil {
				reportMetrics.Histogram(histogramEntity)
			}
			if summaryEntity != nil {
				reportMetrics.Summary(summaryEntity)
			}
			s.reset(s.tx.tracer)
		}
		s.SpanData = nil
		return
	}
	if len(s.stacktrace) == 0 && s.Duration >= s.stackFramesMinDuration {
		s.setStacktrace(1)
	}
	if s.tx != nil {
		duration := s.reportSelfTime()
		reportMetrics := s.tracer.reportMetrics(labels, duration)
		if histogramEntity != nil {
			reportMetrics.Histogram(histogramEntity)
		}
		if summaryEntity != nil {
			reportMetrics.Summary(summaryEntity)
		}
	}
	s.enqueue()
	s.SpanData = nil
}
