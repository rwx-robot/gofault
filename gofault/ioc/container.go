// Package ioc provides the inversion-of-control container for dependency injection.
package ioc

import (
	"fmt"
	"reflect"
	"sync"
)

// Scope defines the lifecycle scope of a registered service.
type Scope int

const (
	ScopeSingleton Scope = iota
	ScopeRequest
	ScopeTransient
)

// Container is the publicly exposed IoC container type (alias of the internal container).
type Container = container

// container is the inversion-of-control container.
type container struct {
	mu         sync.RWMutex
	singletons map[reflect.Type]*entry
	transients map[reflect.Type]*entry
	request    map[reflect.Type]*entry
}

type entry struct {
	ctor any
	inst any // nil means not-yet-instantiated (for singletons)
}

// New creates a new container.
func New() *container {
	return &container{
		singletons: make(map[reflect.Type]*entry),
		transients: make(map[reflect.Type]*entry),
		request:    make(map[reflect.Type]*entry),
	}
}

// Register registers a singleton constructor. The constructor must be a function
// that returns a single value.
func (c *container) Register(ctor any) error {
	return c.register(ctor, ScopeSingleton)
}

// RegisterScoped registers a request-scoped constructor.
func (c *container) RegisterScoped(ctor any) error {
	return c.register(ctor, ScopeRequest)
}

// RegisterTransient registers a transient (every-resolve) constructor.
func (c *container) RegisterTransient(ctor any) error {
	return c.register(ctor, ScopeTransient)
}

func (c *container) register(ctor any, scope Scope) error {
	fn := reflect.ValueOf(ctor)
	if fn.Kind() != reflect.Func {
		return fmt.Errorf("ctor must be a function, got %T", ctor)
	}
	t := fn.Type()
	if t.NumOut() != 1 {
		return fmt.Errorf("ctor must return exactly one value, got %d", t.NumOut())
	}
	key := t.Out(0)
	c.mu.Lock()
	defer c.mu.Unlock()
	switch scope {
	case ScopeSingleton:
		c.singletons[key] = &entry{ctor: ctor}
	case ScopeTransient:
		c.transients[key] = &entry{ctor: ctor}
	case ScopeRequest:
		c.request[key] = &entry{ctor: ctor}
	}
	return nil
}

// Resolve instantiates or returns the cached instance for the given type.
// target must be a pointer to the desired type (e.g., &MyService{}).
func (c *container) Resolve(target any) (any, error) {
	t := reflect.TypeOf(target)
	if t == nil {
		return nil, fmt.Errorf("nil type")
	}
	if t.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("target must be a pointer, got %T", target)
	}

	// Check singleton: if registered and instantiated, return cached instance.
	c.mu.RLock()
	e, isSingleton := c.singletons[t]
	c.mu.RUnlock()
	if isSingleton {
		return c.resolveSingleton(e, t)
	}

	// Check transient.
	c.mu.RLock()
	e, ok := c.transients[t]
	c.mu.RUnlock()
	if ok {
		return c.callCtor(e.ctor)
	}

	// Check request.
	c.mu.RLock()
	e, ok = c.request[t]
	c.mu.RUnlock()
	if ok {
		return c.callCtor(e.ctor)
	}

	return nil, fmt.Errorf("no registered service for type %v", t)
}

// resolveSingleton returns the singleton instance, creating it lazily if needed.
func (c *container) resolveSingleton(e *entry, t reflect.Type) (any, error) {
	// Fast path: already instantiated.
	c.mu.RLock()
	inst := e.inst
	c.mu.RUnlock()
	if inst != nil {
		return inst, nil
	}

	// Slow path: lazily create the singleton instance.
	c.mu.Lock()
	defer c.mu.Unlock()
	// Double-check after acquiring write lock.
	if e.inst != nil {
		return e.inst, nil
	}
	inst, err := c.callCtor(e.ctor)
	if err != nil {
		return nil, err
	}
	e.inst = inst
	return inst, nil
}

func (c *container) callCtor(ctor any) (any, error) {
	fn := reflect.ValueOf(ctor)
	in := make([]reflect.Value, fn.Type().NumIn())
	for i := range in {
		argType := fn.Type().In(i)
		arg, err := c.resolveType(argType)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve arg %d (%v): %w", i, argType, err)
		}
		in[i] = reflect.ValueOf(arg)
	}
	out := fn.Call(in)
	if len(out) == 2 && !out[1].IsNil() {
		return out[0].Interface(), out[1].Interface().(error)
	}
	if len(out) == 1 {
		return out[0].Interface(), nil
	}
	return out[0].Interface(), nil
}

func (c *container) resolveType(t reflect.Type) (any, error) {
	// Check if it's a registered singleton.
	c.mu.RLock()
	e, isSingleton := c.singletons[t]
	c.mu.RUnlock()
	if isSingleton {
		return c.resolveSingleton(e, t)
	}

	// Check transients.
	c.mu.RLock()
	e, ok := c.transients[t]
	c.mu.RUnlock()
	if ok {
		return c.callCtor(e.ctor)
	}

	// Check request-scoped.
	c.mu.RLock()
	e, ok = c.request[t]
	c.mu.RUnlock()
	if ok {
		return c.callCtor(e.ctor)
	}

	return nil, fmt.Errorf("cannot resolve %v", t)
}
