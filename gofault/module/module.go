// Package module provides the application root that bootstraps all registered modules.
package module

import (
	"fmt"

	"github.com/gofault/gofault/controller"
	"github.com/gofault/gofault/core"
	"github.com/gofault/gofault/ioc"
	"github.com/gofault/gofault/provider"
	"github.com/gofault/gofault/router"
	"github.com/gofault/gofault/server"
)

// App is the root application structure that ties everything together.
type App struct {
	container *ioc.Container
	rtr       *router.Router
	modules   []*core.Module
	server    *server.HTTP
	booted    bool
}

// New creates a new application instance.
func New() *App {
	return &App{
		container: ioc.New(),
		rtr:       router.New(),
		modules:   make([]*core.Module, 0),
	}
}

// SetRouter replaces the default router with a custom one (e.g., to add filters).
func (a *App) SetRouter(r *router.Router) {
	a.rtr = r
}

// RegisterModules registers one or more application modules.
func (a *App) RegisterModules(mods ...*core.Module) {
	for _, m := range mods {
		for _, ctrl := range m.Controllers {
			a.container.Register(func() core.Controller { return ctrl })
		}
		for _, prov := range m.Providers {
			a.container.Register(func() provider.Provider { return prov })
		}
		a.modules = append(a.modules, m)
	}
}

// makeHandler creates a handler closure that properly captures route values,
// avoiding the classic Go closure capture bug where loop variables are
// captured by reference instead of by value.
func makeHandler(ctrl core.Controller, route core.Route) core.Handler {
	return func(ctx *core.Ctx) error {
		return controller.InvokeHandler(ctrl, route.Method, route.Path, route.Handler, ctx)
	}
}

// Bootstrap finalizes the application setup, registering all routes.
func (a *App) Bootstrap() error {
	for _, mod := range a.modules {
		for _, ctrl := range mod.Controllers {
			for _, route := range ctrl.Routes() {
				fullPath := ctrl.Prefix() + route.Path
				// Pass route as parameter to avoid classic Go closure capture bug.
				a.rtr.Handle(route.Method, fullPath, makeHandler(ctrl, route), mod.Middleware...)
			}
		}
	}
	a.booted = true
	return nil
}

// Start launches the HTTP server on the given port.
func (a *App) Start(port int) error {
	if !a.booted {
		if err := a.Bootstrap(); err != nil {
			return fmt.Errorf("bootstrap failed: %w", err)
		}
	}

	// Run OnBoot hooks.
	for _, mod := range a.modules {
		for _, hook := range mod.OnBootHooks {
			if err := hook.OnBoot(); err != nil {
				return fmt.Errorf("OnBoot hook failed: %w", err)
			}
		}
	}

	a.server = server.New(a.rtr, port)
	return a.server.Start()
}

// Stop gracefully shuts down the server and calls OnShutdown hooks.
func (a *App) Stop() error {
	var errs []error

	// Call OnShutdown hooks in reverse order.
	for i := len(a.modules) - 1; i >= 0; i-- {
		mod := a.modules[i]
		for j := len(mod.OnShutdownHooks) - 1; j >= 0; j-- {
			hook := mod.OnShutdownHooks[j]
			if err := hook.OnShutdown(); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if a.server != nil {
		if err := a.server.Stop(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}
	return nil
}

// Container returns the IoC container (exposed for testing).
func (a *App) Container() *ioc.Container {
	return a.container
}
