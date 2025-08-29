package inertia

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sheenazien8/inertia-go"
)

// FiberAdapter wraps the inertia manager for Fiber compatibility
type FiberAdapter struct {
	manager *inertia.Inertia
}

// NewFiberAdapter creates a new Fiber adapter for Inertia
func NewFiberAdapter(manager *inertia.Inertia) *FiberAdapter {
	return &FiberAdapter{manager: manager}
}

// Render renders an Inertia response in Fiber context
func (fa *FiberAdapter) Render(c *fiber.Ctx, component string, props inertia.Props) error {
	// Create a buffer to capture the response
	var buf bytes.Buffer

	// Create a mock http.ResponseWriter
	w := &mockResponseWriter{
		buffer: &buf,
		header: make(http.Header),
	}

	// Create a mock http.Request from Fiber context
	req, err := http.NewRequest(c.Method(), c.OriginalURL(), bytes.NewReader(c.Body()))
	if err != nil {
		return err
	}

	// Copy headers
	for key, values := range c.GetReqHeaders() {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	// Use the inertia manager to render
	err = fa.manager.Render(w, req, component, props)
	if err != nil {
		return err
	}

	// Set response headers from mock writer
	for key, values := range w.header {
		for _, value := range values {
			c.Response().Header.Set(key, value)
		}
	}

	// Set status code and send response
	if w.statusCode != 0 {
		c.Status(w.statusCode)
	}

	return c.Send(buf.Bytes())
}

// mockResponseWriter implements http.ResponseWriter for Fiber compatibility
type mockResponseWriter struct {
	buffer     *bytes.Buffer
	header     http.Header
	statusCode int
}

func (m *mockResponseWriter) Header() http.Header {
	return m.header
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	return m.buffer.Write(data)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}
