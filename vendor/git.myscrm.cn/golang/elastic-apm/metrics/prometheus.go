package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// 访问记录（RPC方法、错误码）
	AccessTracing = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stark_requests_total",
		},
		[]string{"method", "code"},
	)

	// 内部RPC调用的访问记录（RPC方法、错误码）
	InternalAccessTracing = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stark_inside_requests_total",
		},
		[]string{"method", "code"},
	)

	// 异常捕获（RPC方法）
	RecoveryPanic = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stark_recovery_total",
		},
		[]string{"method"},
	)

	// 错误日志的记录次数
	ErrorLogging = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "stark_error_total",
		},
	)

	// SQL语句执行的耗时（直方图）
	SQLHistogramTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_sql_tracing_h_total",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"span_name"},
	)

	// SQL语句执行的耗时（摘要）
	SQLSummaryTracing = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stark_sql_tracing_s_seconds",
		},
		[]string{"span_name"},
	)

	// HTTP语句执行的耗时（直方图）
	HTTPHistogramTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_http_tracing_h_total",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"span_name"},
	)

	// HTTP语句执行的耗时（摘要）
	HTTPSummaryTracing = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stark_http_tracing_s_seconds",
		},
		[]string{"span_name"},
	)

	// 内部RPC调用的耗时（直方图）
	RPCHistogramTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_rpc_tracing_h_total",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"method"},
	)

	// 内部RPC调用的耗时（摘要）
	RPCSummaryTracing = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stark_rpc_tracing_s_seconds",
		},
		[]string{"method"},
	)

	// Redis执行的耗时（直方图）
	RedisHistogramTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_redis_tracing_h_total",
			Buckets: []float64{0.01, 0.03, 0.05, 0.1, 0.3, 0.5, 1},
		},
		[]string{"command_name"},
	)

	// 外部HTTP调用的耗时（服务端）
	HTTPHistogramTransactionTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_http_transaction_tracing_h_total",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"transaction_name"},
	)

	HTTPSummaryTransactionTracing = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stark_http_transaction_tracing_s_seconds",
		},
		[]string{"transaction_name"},
	)

	// 外部RPC调用的耗时（服务端）
	RPCHistogramTransactionTracing = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stark_rpc_transaction_tracing_h_total",
			Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
		},
		[]string{"method"},
	)

	RPCSummaryTransactionTracing = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "stark_rpc_transaction_tracing_s_seconds",
		},
		[]string{"method"},
	)

	CircuitBreaker = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stark_circuit_breaker_total",
		},
		[]string{"resource_name", "state"},
	)

	FlowControl = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "stark_flow_control_total",
		},
		[]string{"resource_name", "state"},
	)
)

func init() {
	prometheus.MustRegister(AccessTracing)
	prometheus.MustRegister(InternalAccessTracing)
	prometheus.MustRegister(RecoveryPanic)
	prometheus.MustRegister(ErrorLogging)

	prometheus.MustRegister(SQLHistogramTracing)
	prometheus.MustRegister(SQLSummaryTracing)

	prometheus.MustRegister(HTTPSummaryTracing)
	prometheus.MustRegister(HTTPHistogramTracing)

	prometheus.MustRegister(RPCHistogramTracing)
	prometheus.MustRegister(RPCSummaryTracing)

	prometheus.MustRegister(HTTPHistogramTransactionTracing)
	prometheus.MustRegister(HTTPSummaryTransactionTracing)

	prometheus.MustRegister(RPCHistogramTransactionTracing)
	prometheus.MustRegister(RPCSummaryTransactionTracing)

	prometheus.MustRegister(RedisHistogramTracing)

	prometheus.MustRegister(CircuitBreaker)
	prometheus.MustRegister(FlowControl)
}
