// Package server provides the HTTP server implementation.
package server

import (
	"fmt"
	"net/http"
)

// HTTP wraps the standard library HTTP server.
type HTTP struct {
	server *http.Server
}

// New creates a new HTTP server that will listen on the given port.
func New(handler http.Handler, port int) *HTTP {
	return &HTTP{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		},
	}
}

// Start launches the server and blocks.
func (s *HTTP) Start() error {
	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the server.
func (s *HTTP) Stop() error {
	return s.server.Close()
}
