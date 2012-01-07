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

# Commit: feat(module): implement singleton scope

# Commit: feat(controller): add param extraction

# Commit: fix(router): resolve routing conflict

# Commit: fix(server): resolve body parsing

# Commit: refactor(provider): restructure documentation

# Commit: docs(core): update documentation

# Commit: refactor(example): restructure type safety

# Commit: perf(hello): optimize module setup

# Commit: perf(test): optimize param extraction

# Commit: feat(docs): implement route matching

# Commit: refactor(middleware): improve performance

# Commit: feat(di): add response writing

# Commit: refactor(di-container): improve error messages

# Commit: refactor(route): improve error messages

# Commit: refactor(http): improve memory usage

# Commit: refactor(ioc): restructure code structure

# Commit: fix(http): handle nil pointer case

# Commit: feat(ioc): add controller routing

# Commit: chore(container): update dependencies

# Commit: feat(module): add error handling

# Commit: fix(controller): handle type inference case

# Commit: test(router): add coverage for controller routing

# Commit: chore(server): update gitignore

# Commit: chore(provider): update gitignore

# Commit: fix(core): handle body parsing case

# Commit: chore(example): update readme

# Commit: refactor(hello): restructure documentation

# Commit: docs(test): update documentation

# Commit: feat(docs): implement singleton scope

# Commit: feat(middleware): add middleware chain

# Commit: feat(di): implement param extraction

# Commit: feat(di-container): add request injection

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): improve documentation

# Commit: docs(ioc): update documentation

# Commit: feat(container): implement response writing

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update test suite

# Commit: test(module): add coverage for param extraction

# Commit: chore(controller): update readme

# Commit: docs(router): update documentation

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for error handling

# Commit: fix(core): handle routing conflict case

# Commit: fix(example): resolve type inference

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize controller routing

# Commit: chore(docs): update go mod

# Commit: feat(middleware): implement context propagation

# Commit: refactor(di): restructure documentation

# Commit: docs(di-container): update documentation

# Commit: feat(route): add controller routing

# Commit: chore(http): update readme

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize route matching

# Commit: refactor(module): improve error messages

# Commit: test(container): add coverage for controller routing

# Commit: chore(module): update ci configuration

# Commit: fix(controller): resolve nil pointer

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: feat(provider): add provider registration

# Commit: docs(core): update documentation

# Commit: refactor(example): improve test coverage

# Commit: refactor(hello): restructure concurrency handling

# Commit: feat(test): implement request injection

# Commit: fix(docs): resolve scope resolution

# Commit: refactor(middleware): improve error messages

# Commit: refactor(di): improve error messages

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement module setup

# Commit: perf(http): optimize param extraction

# Commit: perf(ioc): optimize param extraction

# Commit: docs(container): update documentation

# Commit: test(module): add coverage for singleton scope

# Commit: feat(controller): add singleton scope

# Commit: docs(module): update documentation

# Commit: feat(controller): implement response writing

# Commit: refactor(router): improve code structure

# Commit: fix(server): resolve scope resolution

# Commit: refactor(provider): improve test coverage
