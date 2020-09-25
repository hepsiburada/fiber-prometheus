package fiberprometheus

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware_Register(t *testing.T) {
	app := fiber.New()
	m := NewMiddleware("test_namespace", "test_http_subsystem", "/test-metrics")
	m.Register(app)

	app.Get("/testurl", func(c *fiber.Ctx) error {
		return c.SendString("Hello test")
	})

	req := httptest.NewRequest(
		"GET",
		"/testurl",
		nil,
	)
	resp, _ := app.Test(req, -1)

	assert.NotNil(t, resp)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Get metrics
	req = httptest.NewRequest(
		"GET",
		"/test-metrics",
		nil,
	)
	resp, _ = app.Test(req, -1)

	responseBody, _ := ioutil.ReadAll(resp.Body)

	reqDurationText := "test_namespace_test_http_subsystem_request_duration_seconds_bucket{handler=\"/testurl\",method=\"GET\",le=\"5\"} 1"
	reqCountText := "test_namespace_test_http_subsystem_requests_total{method=\"GET\",path=\"/testurl\",status_code=\"200\"} 1"

	assert.Contains(t, string(responseBody), reqDurationText)
	assert.Contains(t, string(responseBody), reqCountText)
}

func TestMiddleware_SetupPath(t *testing.T) {
	app := fiber.New()
	m := NewMiddleware("test_namespace", "test_http_subsystem", "/test-metrics")
	m.SetupPath(app)

	assert.Equal(t, "/test-metrics", m.MetricPath)
}

func TestNewMiddleware(t *testing.T) {
	m := NewMiddleware("test_namespace", "test_http_subsystem", "/test-metrics")
	assert.NotNil(t, m)
	assert.Equal(t, "test_namespace", m.Namespace)
	assert.Equal(t, "test_http_subsystem", m.Subsystem)
	assert.Equal(t, "/test-metrics", m.MetricPath)
}
