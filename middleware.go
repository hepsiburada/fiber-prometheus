package fiberprometheus

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Middleware struct {
	Namespace   string
	Subsystem   string
	MetricPath  string
	reqCount    *prometheus.CounterVec
	reqDuration *prometheus.HistogramVec
}

func (m *Middleware) PrometheusHandler() fiber.Handler {
	return func(c *fiber.Ctx) {
		if c.Path() == m.MetricPath {
			c.Next()
			return
		}

		start := time.Now()

		c.Next()

		r := c.Route()

		statusCode := strconv.Itoa(c.Fasthttp.Response.StatusCode())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		m.reqCount.WithLabelValues(statusCode, c.Method(), r.Path).Inc()
		m.reqDuration.WithLabelValues(c.Method(), r.Path).Observe(elapsed)
	}
}

func (m *Middleware) Register(app *fiber.App) {
	m.registerDefaultMetrics()
	m.SetupPath(app)
	app.Use(m.PrometheusHandler())
}

func (m *Middleware) SetupPath(app *fiber.App) {
	app.Get(m.MetricPath, m.metricHandler)
}

func (m *Middleware) metricHandler(c *fiber.Ctx) {
	p := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	p(c.Fasthttp)
}

func (m *Middleware) registerDefaultMetrics() {
	m.reqCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "requests_total",
			Namespace: m.Namespace,
			Subsystem: m.Subsystem,
			Help:      "Number of HTTP requests",
		},
		[]string{"status_code", "method", "path"},
	)

	m.reqDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "request_duration_seconds",
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Help:      "Duration of HTTP requests",
	}, []string{"method", "handler"})
}

func NewMiddleware(namespace string, subsystem string, metricPath string) *Middleware {
	return &Middleware{
		Namespace:  namespace,
		Subsystem:  subsystem,
		MetricPath: metricPath,
	}
}
