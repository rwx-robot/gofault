# gofault

A Go backend framework inspired by NestJS architecture, implemented in pure Go idioms.

## Architecture

```
┌──────────────────────────────────────────────┐
│                  App                         │
│  ┌────────────────────────────────────────┐  │
│  │           HTTP Server (:9090)          │  │
│  └────────────────────────────────────────┘  │
│                     │                        │
│  ┌────────────────────────────────────────┐  │
│  │              Router                    │  │
│  │   /hello/greet/:name   GET             │  │
│  │   /hello/            GET               │  │
│  └────────────────────────────────────────┘  │
│                     │                        │
│  ┌────────────────────────────────────────┐  │
│  │         Controller (Greeting)          │  │
│  │   Prefix: /hello                       │  │
│  │   Routes: /greet/:name, /              │  │
│  └────────────────────────────────────────┘  │
│                     │                        │
│  ┌────────────────────────────────────────┐  │
│  │      IoC Container (Singleton)         │  │
│  │   GreetingService                       │  │
│  └────────────────────────────────────────┘  │
└──────────────────────────────────────────────┘
```

## Core Interfaces

- `Module`: The basic application unit with controllers and providers
- `Controller`: Implements `Routes()` and `Prefix()` to expose HTTP endpoints
- `Provider`: Implements `Provide()` to supply injectable dependencies
- `MiddlewareFunc`: `func(ctx *Ctx, next Handler) error` for request preprocessing

## Features (v0.1.0)

- IoC container with singleton scope
- Module-based application organization
- HTTP routing with path parameter extraction
- Controller + Provider layering
- Minimal middleware chain

# Commit: feat(di): implement module setup

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for middleware chain

# Commit: feat(http): implement middleware chain

# Commit: feat(ioc): implement module setup

# Commit: perf(container): optimize request injection

# Commit: fix(module): resolve header setting

# Commit: perf(controller): optimize module setup
