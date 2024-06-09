package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

var (
	apiCalledLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "server",
			Subsystem: "api",
			Name:      "latency_microseconds",
			Help:      "Latency of API calls in microseconds",
		},
		[]string{"api", "method", "status_code"},
	)
)

func init() {
	prometheus.MustRegister(apiCalledLatency)
}

func MetricsExportMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now() // 记录请求开始时间
		defer func() {
			cost := time.Since(start).Nanoseconds()     // 计算请求耗时
			api := ctx.Request.URL.Path                 // 获取请求的 URL 路径
			method := ctx.Request.Method                // 获取请求的方法（如 GET、POST 等）
			status := strconv.Itoa(ctx.Writer.Status()) // 获取响应状态码

			// 忽略对 "/metrics" 路径的请求
			if api == "/metrics" {
				return
			}

			// 记录请求的延迟时间
			apiCalledLatency.WithLabelValues(api, method, status).Observe(float64(cost))
		}()
		ctx.Next() // 继续处理请求
	}
}
