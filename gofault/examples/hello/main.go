// Package main is a minimal example demonstrating gofault application structure.
package main

import (
	"log"

	"github.com/gofault/gofault/controller"
	"github.com/gofault/gofault/core"
	"github.com/gofault/gofault/provider"
	"github.com/gofault/gofault/router"
	"github.com/gofault/gofault/server"
)

// GreetingService is a simple injectable provider.
type GreetingService struct {
	prefix string
}

func NewGreetingService(prefix string) *GreetingService {
	return &GreetingService{prefix: prefix}
}

func (s *GreetingService) Provide() any { return s }

func (s *GreetingService) Greet(name string) string {
	return s.prefix + ", " + name + "!"
}

// GreetingController is the HTTP entry point for greeting endpoints.
type GreetingController struct {
	controller.BaseController
	svc *GreetingService
}

func NewGreetingController(svc *GreetingService) *GreetingController {
	return &GreetingController{svc: svc}
}

func (c *GreetingController) Prefix() string { return "/hello" }

func (c *GreetingController) Routes() []core.Route {
	return []core.Route{
		{Method: "GET", Path: "/greet/:name"},
		{Method: "GET", Path: "/"},
	}
}

func (c *GreetingController) Greet(ctx *core.Ctx) error {
	name := controller.Param(ctx, "name")
	msg := c.svc.Greet(name)
	return controller.OK(ctx.Response, map[string]string{"message": msg})
}

func (c *GreetingController) Index(ctx *core.Ctx) error {
	return controller.OK(ctx.Response, map[string]string{"message": "gofault is running"})
}

func main() {
	// Create service and controller
	svc := NewGreetingService("Hello")
	ctrl := NewGreetingController(svc)
	_ = provider.ValueProvider{Value: svc}

	// Create router
	rtr := router.New()

	// Register all routes from the controller
	for _, route := range ctrl.Routes() {
		fullPath := ctrl.Prefix() + route.Path
		switch route.Path {
		case "/greet/:name":
			rtr.Handle(route.Method, fullPath, ctrl.Greet)
		case "/":
			rtr.Handle(route.Method, fullPath, ctrl.Index)
		}
	}

	// Create and start server
	httpServer := server.New(rtr, 9090)
	log.Println("gofault hello example running on :9090")
	log.Fatal(httpServer.Start())
}
