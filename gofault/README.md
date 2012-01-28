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

# Commit: test(core): add coverage for response writing

# Commit: docs(example): update documentation

# Commit: chore(hello): update license

# Commit: perf(test): optimize provider registration

# Commit: docs(docs): update documentation

# Commit: docs(middleware): update documentation

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize controller routing

# Commit: chore(route): update build script

# Commit: refactor(http): improve type safety

# Commit: perf(ioc): optimize middleware chain

# Commit: feat(container): implement provider registration

# Commit: docs(module): update documentation

# Commit: fix(controller): handle routing conflict case

# Commit: fix(router): resolve header setting

# Commit: fix(controller): handle pattern matching case

# Commit: refactor(router): restructure test coverage

# Commit: test(server): add coverage for request injection

# Commit: fix(provider): resolve nil pointer

# Commit: feat(core): add provider registration

# Commit: docs(example): update documentation

# Commit: feat(hello): add handler resolution

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement response writing

# Commit: fix(di): handle routing conflict case

# Commit: fix(di-container): resolve pattern matching

# Commit: refactor(route): restructure test coverage

# Commit: test(http): add coverage for error handling

# Commit: fix(ioc): handle body parsing case

# Commit: chore(container): update dependencies

# Commit: feat(module): add route matching

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve performance

# Commit: feat(server): add context propagation

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure test coverage

# Commit: test(provider): add coverage for route matching

# Commit: refactor(core): improve documentation

# Commit: fix(example): handle nil pointer case

# Commit: feat(hello): add route matching

# Commit: refactor(test): improve code structure

# Commit: fix(docs): resolve body parsing

# Commit: chore(middleware): update ci configuration

# Commit: fix(di): resolve scope resolution

# Commit: refactor(di-container): improve error messages

# Commit: fix(route): handle scope resolution case

# Commit: refactor(http): improve concurrency handling

# Commit: fix(ioc): handle header setting case

# Commit: perf(container): optimize response writing

# Commit: fix(module): resolve routing conflict

# Commit: fix(controller): resolve body parsing

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for route matching

# Commit: refactor(provider): improve error messages

# Commit: refactor(server): restructure code structure

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): resolve header setting

# Commit: perf(example): optimize error handling

# Commit: fix(hello): handle scope resolution case

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): restructure code structure

# Commit: fix(di): resolve path extraction

# Commit: fix(di-container): handle header setting case

# Commit: perf(route): optimize singleton scope

# Commit: feat(http): add error handling

# Commit: fix(ioc): handle path extraction case

# Commit: feat(container): implement request injection

# Commit: fix(module): resolve nil pointer

# Commit: feat(controller): add error handling

# Commit: fix(router): handle body parsing case

# Commit: fix(server): handle scope resolution case

# Commit: refactor(provider): improve performance

# Commit: feat(core): add handler resolution

# Commit: docs(provider): update documentation

# Commit: fix(core): handle path extraction case

# Commit: fix(example): handle routing conflict case

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize controller routing

# Commit: chore(docs): update go mod

# Commit: feat(middleware): implement route matching

# Commit: refactor(di): improve code structure

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): resolve header setting

# Commit: perf(http): optimize error handling

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: docs(module): update documentation

# Commit: refactor(controller): restructure concurrency handling

# Commit: feat(router): implement controller routing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement middleware chain

# Commit: feat(core): implement request injection

# Commit: fix(example): resolve nil pointer

# Commit: docs(core): update documentation

# Commit: fix(example): handle scope resolution case

# Commit: refactor(hello): improve error messages

# Commit: refactor(test): improve memory usage

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): resolve path extraction

# Commit: feat(di): implement module setup

# Commit: perf(di-container): optimize request injection

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve performance

# Commit: feat(controller): add controller routing

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for controller routing

# Commit: chore(provider): update build script

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize controller routing

# Commit: chore(hello): update test suite

# Commit: feat(example): implement context propagation

# Commit: refactor(hello): restructure error messages

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve scope resolution

# Commit: refactor(middleware): improve type safety

# Commit: perf(di): optimize context propagation

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: fix(http): handle body parsing case

# Commit: chore(ioc): update test suite

# Commit: test(container): add coverage for handler resolution

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize module setup

# Commit: perf(router): optimize response writing

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement middleware chain

# Commit: feat(core): implement middleware chain

# Commit: feat(example): implement singleton scope

# Commit: feat(hello): add route matching

# Commit: refactor(test): improve code structure

# Commit: refactor(hello): restructure error messages

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for module setup

# Commit: perf(middleware): optimize singleton scope

# Commit: feat(di): add middleware chain

# Commit: feat(di-container): implement route matching

# Commit: refactor(route): improve memory usage

# Commit: refactor(http): restructure documentation

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve code structure

# Commit: fix(module): resolve header setting

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add route matching

# Commit: refactor(server): improve memory usage

# Commit: chore(provider): update dependencies

# Commit: feat(core): add provider registration

# Commit: docs(example): update documentation

# Commit: refactor(hello): improve performance

# Commit: feat(test): add middleware chain

# Commit: feat(docs): implement error handling

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize param extraction

# Commit: fix(middleware): resolve type inference

# Commit: test(di): add coverage for response writing

# Commit: feat(di-container): add controller routing

# Commit: chore(route): update readme

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement param extraction

# Commit: perf(container): optimize handler resolution

# Commit: test(module): add coverage for controller routing

# Commit: chore(controller): update ci configuration

# Commit: fix(router): resolve scope resolution

# Commit: refactor(server): improve concurrency handling

# Commit: feat(provider): implement handler resolution

# Commit: test(core): add coverage for middleware chain

# Commit: feat(example): implement handler resolution

# Commit: test(hello): add coverage for provider registration

# Commit: docs(test): update documentation

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement response writing

# Commit: docs(docs): update documentation

# Commit: test(middleware): add coverage for provider registration

# Commit: docs(di): update documentation

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement response writing

# Commit: fix(http): handle nil pointer case

# Commit: feat(ioc): add module setup

# Commit: perf(container): optimize response writing

# Commit: feat(module): implement handler resolution

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle body parsing case

# Commit: chore(provider): update ci configuration

# Commit: fix(core): resolve routing conflict

# Commit: fix(example): resolve routing conflict

# Commit: fix(hello): resolve header setting

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): improve code structure

# Commit: fix(di): handle scope resolution case

# Commit: refactor(middleware): improve code structure

# Commit: fix(di): resolve nil pointer

# Commit: fix(di-container): handle routing conflict case

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): restructure error messages

# Commit: refactor(ioc): improve code structure

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for singleton scope

# Commit: feat(controller): add module setup

# Commit: perf(router): optimize param extraction

# Commit: refactor(server): improve code structure

# Commit: fix(provider): resolve scope resolution

# Commit: refactor(core): restructure type safety

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize module setup

# Commit: perf(test): optimize provider registration

# Commit: docs(docs): update documentation

# Commit: fix(middleware): handle path extraction case

# Commit: feat(di): implement module setup

# Commit: perf(di-container): optimize module setup

# Commit: feat(di): add context propagation

# Commit: refactor(di-container): restructure memory usage

# Commit: chore(route): update dependencies

# Commit: refactor(http): restructure type safety

# Commit: perf(ioc): optimize response writing

# Commit: refactor(container): improve error messages

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: feat(router): implement request injection

# Commit: fix(server): resolve header setting

# Commit: perf(provider): optimize error handling

# Commit: fix(core): handle routing conflict case

# Commit: fix(example): handle header setting case

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update license

# Commit: refactor(docs): restructure documentation

# Commit: docs(middleware): update documentation

# Commit: refactor(di): restructure code structure

# Commit: fix(di-container): resolve pattern matching

# Commit: docs(route): update documentation

# Commit: fix(di-container): handle body parsing case

# Commit: chore(route): update test suite

# Commit: fix(http): handle routing conflict case

# Commit: fix(ioc): resolve header setting

# Commit: fix(container): handle header setting case

# Commit: perf(module): optimize controller routing

# Commit: chore(controller): update ci configuration

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve error messages

# Commit: fix(provider): handle header setting case

# Commit: fix(core): handle body parsing case

# Commit: refactor(example): restructure documentation

# Commit: docs(hello): update documentation

# Commit: chore(test): update readme

# Commit: docs(docs): update documentation

# Commit: docs(middleware): update documentation

# Commit: perf(di): optimize context propagation

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: fix(http): resolve header setting

# Commit: chore(route): update readme

# Commit: docs(http): update documentation

# Commit: fix(ioc): resolve body parsing

# Commit: chore(container): update dependencies

# Commit: feat(module): add middleware chain

# Commit: feat(controller): implement controller routing

# Commit: chore(router): update test suite

# Commit: refactor(server): restructure concurrency handling

# Commit: feat(provider): implement controller routing

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve header setting

# Commit: perf(hello): optimize module setup

# Commit: perf(test): optimize request injection

# Commit: fix(docs): resolve nil pointer

# Commit: feat(middleware): add provider registration

# Commit: docs(di): update documentation

# Commit: docs(di-container): update documentation

# Commit: chore(route): update license

# Commit: perf(http): optimize singleton scope

# Commit: feat(ioc): add response writing

# Commit: fix(http): resolve pattern matching

# Commit: docs(ioc): update documentation

# Commit: feat(container): add param extraction

# Commit: feat(module): implement module setup

# Commit: perf(controller): optimize context propagation

# Commit: refactor(router): restructure code structure

# Commit: fix(server): resolve header setting

# Commit: perf(provider): optimize context propagation

# Commit: refactor(core): restructure code structure

# Commit: fix(example): resolve scope resolution

# Commit: refactor(hello): improve memory usage

# Commit: chore(test): update license

# Commit: perf(docs): optimize error handling

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update ci configuration

# Commit: fix(di-container): resolve scope resolution

# Commit: refactor(route): improve test coverage

# Commit: test(http): add coverage for context propagation

# Commit: refactor(ioc): restructure error messages

# Commit: refactor(container): improve performance

# Commit: refactor(ioc): restructure error messages

# Commit: refactor(container): improve memory usage

# Commit: chore(module): update dependencies

# Commit: feat(controller): add singleton scope

# Commit: feat(router): add context propagation

# Commit: refactor(server): restructure code structure

# Commit: fix(provider): resolve pattern matching

# Commit: docs(core): update documentation

# Commit: feat(example): implement param extraction

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle scope resolution case

# Commit: refactor(docs): improve memory usage

# Commit: refactor(middleware): restructure memory usage

# Commit: chore(di): update build script

# Commit: refactor(di-container): improve error messages

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update license

# Commit: perf(ioc): optimize handler resolution

# Commit: test(container): add coverage for middleware chain

# Commit: feat(module): implement route matching

# Commit: test(container): add coverage for module setup

# Commit: perf(module): optimize module setup

# Commit: perf(controller): optimize singleton scope

# Commit: feat(router): add middleware chain

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement provider registration

# Commit: docs(core): update documentation

# Commit: perf(example): optimize route matching

# Commit: refactor(hello): improve memory usage

# Commit: chore(test): update go mod

# Commit: feat(docs): implement singleton scope

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize controller routing

# Commit: chore(di-container): update gitignore

# Commit: chore(route): update readme

# Commit: docs(http): update documentation

# Commit: docs(ioc): update documentation

# Commit: feat(container): add controller routing

# Commit: chore(module): update ci configuration

# Commit: refactor(controller): restructure performance

# Commit: chore(module): update gitignore

# Commit: chore(controller): update ci configuration

# Commit: fix(router): resolve pattern matching

# Commit: docs(server): update documentation

# Commit: fix(provider): handle body parsing case

# Commit: chore(core): update readme

# Commit: docs(example): update documentation

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize error handling

# Commit: fix(docs): handle path extraction case

# Commit: fix(middleware): handle scope resolution case

# Commit: refactor(di): improve type safety

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for controller routing

# Commit: chore(http): update build script

# Commit: refactor(ioc): restructure memory usage

# Commit: chore(container): update dependencies

# Commit: fix(module): handle type inference case

# Commit: test(controller): add coverage for module setup

# Commit: perf(router): optimize controller routing

# Commit: refactor(controller): restructure test coverage

# Commit: test(router): add coverage for handler resolution

# Commit: test(server): add coverage for controller routing

# Commit: chore(provider): update license

# Commit: perf(core): optimize request injection

# Commit: fix(example): resolve type inference

# Commit: test(hello): add coverage for response writing

# Commit: test(test): add coverage for param extraction

# Commit: perf(docs): optimize route matching

# Commit: refactor(middleware): improve documentation

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve scope resolution

# Commit: refactor(route): improve error messages

# Commit: fix(http): handle type inference case

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve header setting

# Commit: perf(module): optimize request injection

# Commit: fix(controller): resolve nil pointer

# Commit: feat(router): add error handling

# Commit: fix(server): handle body parsing case

# Commit: refactor(router): improve error messages

# Commit: refactor(server): improve performance

# Commit: fix(provider): handle pattern matching case

# Commit: docs(core): update documentation

# Commit: fix(example): handle routing conflict case

# Commit: fix(hello): resolve type inference

# Commit: test(test): add coverage for error handling

# Commit: fix(docs): handle type inference case

# Commit: fix(middleware): handle nil pointer case

# Commit: feat(di): add singleton scope

# Commit: feat(di-container): add param extraction

# Commit: fix(route): handle type inference case

# Commit: fix(http): handle type inference case

# Commit: test(ioc): add coverage for module setup

# Commit: perf(container): optimize context propagation

# Commit: refactor(module): restructure error messages

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for request injection

# Commit: fix(server): resolve path extraction

# Commit: feat(provider): implement provider registration

# Commit: docs(server): update documentation

# Commit: fix(provider): resolve path extraction

# Commit: feat(core): implement error handling

# Commit: fix(example): handle pattern matching case

# Commit: docs(hello): update documentation

# Commit: fix(test): resolve type inference

# Commit: test(docs): add coverage for context propagation

# Commit: refactor(middleware): restructure type safety

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add param extraction

# Commit: refactor(route): improve code structure

# Commit: fix(http): handle scope resolution case

# Commit: fix(ioc): handle routing conflict case

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: refactor(router): restructure performance

# Commit: fix(server): handle header setting case

# Commit: perf(provider): optimize middleware chain

# Commit: feat(core): implement provider registration

# Commit: feat(provider): add context propagation

# Commit: refactor(core): restructure documentation

# Commit: fix(example): handle path extraction case

# Commit: feat(hello): implement singleton scope
