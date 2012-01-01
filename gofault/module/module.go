// Package module provides the application root that bootstraps all registered modules.
package module

import (
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
}

// New creates a new application instance.
func New() *App {
	return &App{
		container: ioc.New(),
		rtr:       router.New(),
		modules:   make([]*core.Module, 0),
	}
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

// Bootstrap finalizes the application setup, registering all routes.
func (a *App) Bootstrap() error {
	for _, mod := range a.modules {
		for _, ctrl := range mod.Controllers {
			for _, route := range ctrl.Routes() {
				fullPath := ctrl.Prefix() + route.Path
				a.rtr.Handle(route.Method, fullPath, func(ctx *core.Ctx) error {
					return controller.InvokeHandler(ctrl, route.Method, route.Path, ctx)
				}, mod.Middleware...)
			}
		}
	}
	return nil
}

// Start launches the HTTP server on the given port.
func (a *App) Start(port int) error {
	a.server = server.New(a.rtr, port)
	return a.server.Start()
}

// Stop gracefully shuts down the server.
func (a *App) Stop() error {
	if a.server != nil {
		return a.server.Stop()
	}
	return nil
}

// Container returns the IoC container (exposed for testing).
func (a *App) Container() *ioc.Container {
	return a.container
}
