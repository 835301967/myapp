package apm

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (tx *Transaction) ReportEnd(histogramEntity *prometheus.HistogramVec, summaryEntity *prometheus.SummaryVec, labels map[string]string) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.ended() {
		return
	}
	if tx.Duration < 0 {
		tx.Duration = time.Since(tx.timestamp)
	}
	tx.enqueue()
	endTime := tx.timestamp.Add(tx.Duration)
	duration := tx.Duration - tx.childrenTimer.finalDuration(endTime)
	reportMetrics := tx.tracer.reportMetrics(labels, duration)
	if histogramEntity != nil {
		reportMetrics.Histogram(histogramEntity)
	}
	if summaryEntity != nil {
		reportMetrics.Summary(summaryEntity)
	}
	tx.TransactionData = nil
}
