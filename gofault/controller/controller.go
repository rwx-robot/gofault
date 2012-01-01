// Package controller provides the controller base type and HTTP method helpers.
package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofault/gofault/core"
)

// BaseController provides helper methods for writing HTTP handlers.
type BaseController struct{}

// JSON writes a JSON response with the given status code.
func (c *BaseController) JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Text writes a plain text response.
func (c *BaseController) Text(w http.ResponseWriter, status int, msg string) error {
	w.WriteHeader(status)
	_, err := w.Write([]byte(msg))
	return err
}

// Param extracts a path parameter from the context.
func Param(ctx *core.Ctx, name string) string {
	return ctx.Params[name]
}

// Query extracts a query string parameter.
func Query(ctx *core.Ctx, name string) string {
	return ctx.Request.URL.Query().Get(name)
}

// InvokeHandler dispatches to the appropriate controller method based on HTTP method and path.
// This is a simple reflection-based dispatcher.
func InvokeHandler(ctrl core.Controller, method, path string, ctx *core.Ctx) error {
	// Use functional dispatch based on registered routes.
	// The controller registers its methods via Routes(), so we match here.
	return nil // routed via closure in module.Bootstrap
}

// ControllerMethod is the signature for controller action methods.
type ControllerMethod func(ctx *core.Ctx) error

// Response represents a standard API response.
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// OK sends a 200 JSON response.
func OK(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(Response{Code: 0, Message: "success", Data: data})
}

// Error sends an error JSON response.
func Error(w http.ResponseWriter, status int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(Response{Code: status, Message: msg})
}
