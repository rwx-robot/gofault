package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gofault/gofault/core"
)

// TimeoutConfig holds configuration for the timeout middleware.
type TimeoutConfig struct {
	// Enabled enables the timeout middleware.
	Enabled bool
	// Duration is the timeout duration.
	Duration time.Duration
	// ErrorMessage is the message to return on timeout.
	ErrorMessage string
}

// DefaultTimeoutConfig returns a default timeout configuration.
func DefaultTimeoutConfig() TimeoutConfig {
	return TimeoutConfig{
		Enabled:      true,
		Duration:     30 * time.Second,
		ErrorMessage: "Request timeout",
	}
}

// TimeoutMiddleware creates a middleware that cancels requests that exceed the timeout.
func TimeoutMiddleware(config TimeoutConfig) core.MiddlewareFunc {
	if !config.Enabled {
		return nil
	}

	if config.Duration <= 0 {
		config.Duration = 30 * time.Second
	}

	return func(ctx *core.Ctx, next core.Handler) error {
		ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), config.Duration)
		defer cancel()

		ctx.Request = ctx.Request.WithContext(ctxWithTimeout)

		errCh := make(chan error, 1)
		go func() {
			errCh <- next(ctx)
		}()

		select {
		case err := <-errCh:
			return err
		case <-ctxWithTimeout.Done():
			ctx.Response.Header().Set("Content-Type", "application/json")
			ctx.Response.WriteHeader(http.StatusGatewayTimeout)
			ctx.Response.Write([]byte(`{"status":"error","message":"` + config.ErrorMessage + `"}`))
			return ctxWithTimeout.Err()
		}
	}
}
