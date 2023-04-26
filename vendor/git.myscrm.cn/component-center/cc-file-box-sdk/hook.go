package file_box

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"sync"
	"time"
)

var (
	monitorHookOnce sync.Once
	mk              *monitorHook
)

type MethodName string

type CallInfo struct {
	Method       MethodName
	CloudFactory CloudFactoryType
	Bucket       bucket
	AppCode      string
	Err          error
}

type Hook interface {
	Before(context.Context, *CallInfo) context.Context
	After(context.Context, *CallInfo)
}

type monitorHook struct {
	timer      *prometheus.HistogramVec
	counter    *prometheus.CounterVec
	errCounter *prometheus.CounterVec
}

func newMonitorHook() *monitorHook {
	monitorHookOnce.Do(func() {
		mk = &monitorHook{}
		mk.timer = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "file_box_sdk_request_time",
			Help: "file_box_sdk_request_time",
		}, []string{"cloud_factory", "method", "bucket"})
		mk.counter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "file_box_sdk_request_total",
			Help: "file_box_sdk_request_total",
		}, []string{"cloud_factory", "method", "bucket"})
		mk.errCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "file_box_sdk_request_err_total",
			Help: "file_box_sdk_request_err_total",
		}, []string{"cloud_factory", "method", "bucket"})
		prometheus.MustRegister(mk.timer, mk.counter, mk.errCounter)
	})
	return mk
}

type startTime struct{}

func (m *monitorHook) Before(ctx context.Context, call *CallInfo) context.Context {
	ctx = context.WithValue(ctx, startTime{}, time.Now())
	return ctx
}

func (m *monitorHook) After(ctx context.Context, call *CallInfo) {
	cloudFactoryStr := strconv.Itoa(int(call.CloudFactory))
	v := ctx.Value(startTime{})
	if st, ok := v.(time.Time); ok {
		m.timer.WithLabelValues(cloudFactoryStr, string(call.Method), call.Bucket.name).Observe(time.Now().Sub(st).Seconds())
	}
	if call.Err != nil {
		m.errCounter.WithLabelValues(cloudFactoryStr, string(call.Method), call.Bucket.name).Inc()
	}
	m.counter.WithLabelValues(cloudFactoryStr, string(call.Method), call.Bucket.name).Inc()
}
