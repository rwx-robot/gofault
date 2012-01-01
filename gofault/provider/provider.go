// Package provider provides utilities for defining injectable providers.
package provider

// Provider is the interface for injectable services.
type Provider interface {
	Provide() any
}

// ProviderFunc is an adapter that turns a function into a Provider.
type ProviderFunc func() any

// Provide returns the result of calling the underlying function.
func (f ProviderFunc) Provide() any {
	return f()
}

// ValueProvider wraps a concrete value so it can be registered in the container.
type ValueProvider struct {
	Value any
}

// Provide returns the stored value.
func (v ValueProvider) Provide() any {
	return v.Value
}

// AsScope tags a provider function with a scope label.
// This is a marker type for future scope-based features.
type ScopeTag string

const (
	Singleton ScopeTag = "singleton"
	Request   ScopeTag = "request"
	Transient ScopeTag = "transient"
)
