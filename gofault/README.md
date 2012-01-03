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

# Commit: perf(router): optimize response writing

# Commit: feat(server): add param extraction

# Commit: feat(provider): add param extraction

# Commit: refactor(core): restructure documentation

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve body parsing

# Commit: chore(test): update test suite

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: chore(di-container): update license

# Commit: perf(route): optimize singleton scope

# Commit: test(di-container): add coverage for context propagation

# Commit: refactor(route): restructure test coverage

# Commit: test(http): add coverage for route matching

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement middleware chain

# Commit: feat(module): implement controller routing

# Commit: chore(controller): update go mod

# Commit: feat(router): implement error handling

# Commit: fix(server): handle body parsing case

# Commit: chore(provider): update ci configuration

# Commit: fix(core): handle path extraction case

# Commit: fix(example): handle pattern matching case

# Commit: refactor(hello): restructure code structure

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): improve test coverage

# Commit: test(di): add coverage for route matching

# Commit: refactor(di-container): improve test coverage

# Commit: test(route): add coverage for response writing

# Commit: chore(http): update license

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for request injection

# Commit: fix(ioc): resolve routing conflict

# Commit: fix(container): resolve path extraction
