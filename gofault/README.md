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

# Commit: feat(test): add singleton scope

# Commit: feat(docs): add controller routing

# Commit: chore(middleware): update license

# Commit: perf(di): optimize provider registration

# Commit: docs(di-container): update documentation

# Commit: feat(route): add module setup

# Commit: perf(http): optimize response writing

# Commit: fix(ioc): handle type inference case

# Commit: test(container): add coverage for controller routing

# Commit: chore(module): update gitignore

# Commit: chore(controller): update dependencies

# Commit: feat(router): add handler resolution

# Commit: test(server): add coverage for request injection

# Commit: fix(provider): resolve path extraction

# Commit: fix(core): handle nil pointer case

# Commit: feat(example): add provider registration

# Commit: feat(core): implement provider registration

# Commit: docs(example): update documentation

# Commit: chore(hello): update dependencies

# Commit: feat(test): add handler resolution

# Commit: test(docs): add coverage for controller routing

# Commit: chore(middleware): update test suite

# Commit: test(di): add coverage for param extraction

# Commit: feat(di-container): add response writing

# Commit: chore(route): update readme

# Commit: refactor(http): restructure type safety

# Commit: perf(ioc): optimize provider registration

# Commit: docs(container): update documentation

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): restructure concurrency handling

# Commit: refactor(router): restructure code structure

# Commit: fix(server): resolve pattern matching

# Commit: docs(provider): update documentation

# Commit: perf(core): optimize module setup

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for context propagation

# Commit: test(example): add coverage for provider registration

# Commit: docs(hello): update documentation

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve nil pointer

# Commit: feat(middleware): add param extraction

# Commit: chore(di): update build script

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement request injection

# Commit: fix(http): resolve body parsing

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement handler resolution

# Commit: test(module): add coverage for context propagation

# Commit: refactor(controller): restructure test coverage

# Commit: test(router): add coverage for request injection

# Commit: fix(server): resolve routing conflict

# Commit: fix(provider): resolve path extraction

# Commit: feat(core): implement context propagation

# Commit: refactor(example): restructure memory usage

# Commit: chore(hello): update test suite

# Commit: test(test): add coverage for request injection

# Commit: feat(hello): implement handler resolution

# Commit: test(test): add coverage for middleware chain

# Commit: feat(docs): implement singleton scope

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize context propagation

# Commit: refactor(di-container): restructure type safety

# Commit: fix(route): handle type inference case

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add route matching

# Commit: refactor(container): improve error messages

# Commit: refactor(module): improve performance

# Commit: feat(controller): add param extraction

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for context propagation

# Commit: refactor(provider): restructure concurrency handling

# Commit: feat(core): implement context propagation

# Commit: refactor(example): restructure performance

# Commit: fix(hello): handle scope resolution case

# Commit: fix(test): handle pattern matching case

# Commit: fix(docs): handle header setting case

# Commit: fix(test): resolve pattern matching

# Commit: docs(docs): update documentation

# Commit: chore(middleware): update go mod

# Commit: feat(di): implement route matching

# Commit: refactor(di-container): improve memory usage

# Commit: chore(route): update go mod

# Commit: feat(http): implement controller routing

# Commit: chore(ioc): update license

# Commit: perf(container): optimize route matching

# Commit: refactor(module): improve type safety

# Commit: perf(controller): optimize provider registration

# Commit: docs(router): update documentation

# Commit: fix(server): handle type inference case

# Commit: test(provider): add coverage for request injection

# Commit: fix(core): resolve routing conflict

# Commit: fix(example): resolve body parsing

# Commit: chore(hello): update build script

# Commit: refactor(test): improve performance

# Commit: feat(docs): add context propagation

# Commit: refactor(middleware): restructure documentation

# Commit: test(docs): add coverage for context propagation

# Commit: refactor(middleware): restructure type safety

# Commit: perf(di): optimize response writing

# Commit: chore(di-container): update license

# Commit: perf(route): optimize request injection

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add param extraction

# Commit: refactor(container): restructure code structure

# Commit: fix(module): resolve header setting

# Commit: perf(controller): optimize controller routing

# Commit: chore(router): update license

# Commit: perf(server): optimize module setup

# Commit: perf(provider): optimize route matching

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure type safety

# Commit: perf(test): optimize middleware chain

# Commit: feat(docs): implement handler resolution

# Commit: test(middleware): add coverage for handler resolution

# Commit: test(di): add coverage for param extraction

# Commit: feat(middleware): implement module setup

# Commit: perf(di): optimize context propagation

# Commit: refactor(di-container): restructure concurrency handling

# Commit: fix(route): handle nil pointer case

# Commit: feat(http): add response writing

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve test coverage

# Commit: test(module): add coverage for middleware chain

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update dependencies

# Commit: feat(provider): add context propagation

# Commit: refactor(core): restructure documentation

# Commit: docs(example): update documentation

# Commit: test(hello): add coverage for handler resolution

# Commit: test(test): add coverage for handler resolution

# Commit: test(docs): add coverage for route matching

# Commit: refactor(middleware): improve performance

# Commit: feat(di): add middleware chain

# Commit: feat(di-container): implement provider registration

# Commit: refactor(di): restructure memory usage

# Commit: refactor(di-container): restructure code structure

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for error handling

# Commit: fix(container): handle routing conflict case

# Commit: fix(module): resolve type inference

# Commit: test(controller): add coverage for singleton scope

# Commit: feat(router): add route matching

# Commit: refactor(server): improve error messages

# Commit: fix(provider): handle type inference case

# Commit: fix(core): handle nil pointer case

# Commit: feat(example): add response writing

# Commit: perf(hello): optimize request injection

# Commit: fix(test): resolve path extraction

# Commit: feat(docs): implement error handling

# Commit: fix(middleware): handle path extraction case

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): improve test coverage

# Commit: test(route): add coverage for singleton scope

# Commit: docs(di-container): update documentation

# Commit: refactor(route): improve type safety

# Commit: refactor(http): restructure code structure

# Commit: fix(ioc): resolve header setting

# Commit: perf(container): optimize handler resolution

# Commit: test(module): add coverage for singleton scope

# Commit: feat(controller): add route matching

# Commit: refactor(router): improve error messages

# Commit: refactor(server): improve error messages

# Commit: refactor(provider): improve concurrency handling

# Commit: feat(core): implement route matching

# Commit: refactor(example): improve concurrency handling

# Commit: refactor(hello): restructure test coverage

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve test coverage

# Commit: fix(middleware): handle routing conflict case

# Commit: fix(di): resolve header setting

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for controller routing

# Commit: chore(http): update license

# Commit: feat(route): implement route matching

# Commit: refactor(http): improve performance

# Commit: feat(ioc): add response writing

# Commit: test(container): add coverage for module setup

# Commit: perf(module): optimize context propagation

# Commit: refactor(controller): restructure error messages

# Commit: refactor(router): improve test coverage

# Commit: refactor(server): restructure test coverage

# Commit: test(provider): add coverage for request injection

# Commit: fix(core): resolve type inference

# Commit: fix(example): handle header setting case

# Commit: perf(hello): optimize singleton scope

# Commit: feat(test): add route matching

# Commit: refactor(docs): improve code structure

# Commit: fix(middleware): resolve body parsing

# Commit: refactor(di): restructure performance

# Commit: feat(di-container): add error handling

# Commit: fix(route): handle nil pointer case

# Commit: feat(http): add param extraction

# Commit: refactor(ioc): improve code structure

# Commit: refactor(http): restructure test coverage

# Commit: refactor(ioc): restructure performance

# Commit: refactor(container): restructure documentation

# Commit: docs(module): update documentation

# Commit: docs(controller): update documentation

# Commit: fix(router): handle type inference case

# Commit: test(server): add coverage for module setup

# Commit: perf(provider): optimize provider registration

# Commit: docs(core): update documentation

# Commit: chore(example): update test suite

# Commit: test(hello): add coverage for param extraction

# Commit: refactor(test): improve performance

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize route matching

# Commit: refactor(di): improve error messages

# Commit: refactor(di-container): improve performance

# Commit: feat(route): add context propagation

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add handler resolution

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve test coverage

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize response writing

# Commit: chore(router): update license

# Commit: refactor(server): restructure error messages

# Commit: refactor(provider): improve test coverage

# Commit: test(core): add coverage for context propagation

# Commit: refactor(example): restructure code structure

# Commit: refactor(hello): restructure concurrency handling

# Commit: feat(test): implement route matching

# Commit: refactor(docs): improve test coverage

# Commit: test(middleware): add coverage for request injection

# Commit: fix(di): resolve header setting

# Commit: perf(di-container): optimize singleton scope

# Commit: feat(route): add error handling

# Commit: fix(http): handle routing conflict case

# Commit: fix(ioc): handle scope resolution case

# Commit: refactor(container): improve type safety

# Commit: perf(module): optimize handler resolution

# Commit: refactor(container): restructure memory usage

# Commit: chore(module): update build script

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve performance

# Commit: feat(server): add handler resolution

# Commit: test(provider): add coverage for controller routing

# Commit: chore(core): update license

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure code structure

# Commit: refactor(test): restructure concurrency handling

# Commit: feat(docs): implement middleware chain

# Commit: feat(middleware): implement context propagation

# Commit: refactor(di): restructure memory usage

# Commit: fix(di-container): handle nil pointer case

# Commit: feat(route): add param extraction

# Commit: refactor(http): restructure documentation

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update go mod

# Commit: fix(controller): handle nil pointer case

# Commit: docs(module): update documentation

# Commit: feat(controller): implement response writing

# Commit: refactor(router): restructure type safety

# Commit: perf(server): optimize response writing

# Commit: test(provider): add coverage for response writing

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add request injection

# Commit: fix(hello): resolve scope resolution

# Commit: refactor(test): improve type safety

# Commit: perf(docs): optimize context propagation

# Commit: refactor(middleware): restructure error messages

# Commit: fix(di): handle routing conflict case

# Commit: fix(di-container): resolve type inference

# Commit: test(route): add coverage for response writing

# Commit: perf(http): optimize singleton scope

# Commit: feat(ioc): add request injection

# Commit: fix(container): resolve header setting

# Commit: perf(module): optimize request injection

# Commit: fix(controller): resolve header setting

# Commit: perf(router): optimize singleton scope

# Commit: perf(controller): optimize middleware chain

# Commit: feat(router): implement provider registration

# Commit: docs(server): update documentation

# Commit: fix(provider): handle type inference case

# Commit: test(core): add coverage for route matching

# Commit: refactor(example): improve performance

# Commit: feat(hello): add provider registration

# Commit: docs(test): update documentation

# Commit: feat(docs): add error handling

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: feat(di-container): implement request injection

# Commit: fix(route): resolve header setting

# Commit: perf(http): optimize handler resolution

# Commit: test(ioc): add coverage for route matching

# Commit: refactor(container): improve documentation

# Commit: docs(module): update documentation

# Commit: docs(controller): update documentation

# Commit: docs(router): update documentation

# Commit: refactor(server): improve performance

# Commit: fix(router): resolve header setting

# Commit: perf(server): optimize controller routing

# Commit: chore(provider): update go mod

# Commit: feat(core): implement error handling

# Commit: fix(example): handle header setting case

# Commit: perf(hello): optimize middleware chain

# Commit: feat(test): implement param extraction

# Commit: feat(docs): implement error handling

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: refactor(di-container): restructure type safety

# Commit: refactor(route): restructure type safety

# Commit: perf(http): optimize response writing

# Commit: chore(ioc): update build script

# Commit: refactor(container): improve test coverage

# Commit: test(module): add coverage for middleware chain

# Commit: feat(controller): implement middleware chain

# Commit: feat(router): implement context propagation

# Commit: refactor(server): restructure documentation

# Commit: docs(provider): update documentation

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for handler resolution

# Commit: test(core): add coverage for module setup

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize request injection

# Commit: fix(test): resolve scope resolution

# Commit: refactor(docs): improve performance

# Commit: feat(middleware): add singleton scope

# Commit: feat(di): add context propagation

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: feat(http): add request injection

# Commit: fix(ioc): resolve path extraction

# Commit: feat(container): implement route matching

# Commit: refactor(module): improve code structure

# Commit: refactor(controller): restructure documentation

# Commit: docs(router): update documentation

# Commit: fix(server): resolve pattern matching

# Commit: fix(provider): handle pattern matching case

# Commit: docs(core): update documentation

# Commit: fix(provider): handle type inference case

# Commit: test(core): add coverage for middleware chain

# Commit: feat(example): implement response writing

# Commit: feat(hello): implement handler resolution

# Commit: test(test): add coverage for response writing

# Commit: fix(docs): resolve pattern matching

# Commit: docs(middleware): update documentation

# Commit: docs(di): update documentation

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement param extraction

# Commit: test(http): add coverage for handler resolution

# Commit: test(ioc): add coverage for module setup

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update go mod

# Commit: feat(controller): implement error handling

# Commit: fix(router): handle header setting case

# Commit: perf(server): optimize middleware chain

# Commit: feat(provider): implement context propagation

# Commit: refactor(core): restructure performance

# Commit: feat(example): add middleware chain

# Commit: perf(core): optimize error handling

# Commit: fix(example): handle body parsing case

# Commit: fix(hello): handle header setting case

# Commit: perf(test): optimize singleton scope

# Commit: feat(docs): add request injection

# Commit: fix(middleware): resolve body parsing

# Commit: chore(di): update gitignore

# Commit: chore(di-container): update go mod

# Commit: feat(route): implement request injection

# Commit: fix(http): resolve body parsing

# Commit: refactor(ioc): restructure code structure

# Commit: fix(container): resolve path extraction

# Commit: feat(module): implement context propagation

# Commit: refactor(controller): restructure type safety

# Commit: perf(router): optimize error handling

# Commit: fix(server): handle header setting case

# Commit: perf(provider): optimize handler resolution

# Commit: test(core): add coverage for response writing

# Commit: perf(example): optimize param extraction

# Commit: refactor(hello): improve error messages

# Commit: test(example): add coverage for middleware chain

# Commit: feat(hello): implement module setup

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize handler resolution

# Commit: test(middleware): add coverage for request injection

# Commit: fix(di): resolve body parsing

# Commit: chore(di-container): update go mod

# Commit: feat(route): implement middleware chain

# Commit: feat(http): implement response writing

# Commit: docs(ioc): update documentation

# Commit: feat(container): implement controller routing

# Commit: chore(module): update test suite

# Commit: test(controller): add coverage for provider registration

# Commit: docs(router): update documentation

# Commit: refactor(server): restructure type safety

# Commit: perf(provider): optimize route matching

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize singleton scope

# Commit: feat(hello): add request injection

# Commit: fix(test): resolve nil pointer

# Commit: perf(hello): optimize param extraction

# Commit: feat(test): implement provider registration

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): improve documentation

# Commit: docs(di): update documentation

# Commit: refactor(di-container): improve memory usage

# Commit: refactor(route): restructure error messages

# Commit: refactor(http): improve concurrency handling

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add response writing

# Commit: chore(module): update dependencies

# Commit: feat(controller): add route matching

# Commit: refactor(router): improve documentation

# Commit: docs(server): update documentation

# Commit: feat(provider): add context propagation

# Commit: refactor(core): restructure test coverage

# Commit: refactor(example): restructure documentation

# Commit: docs(hello): update documentation

# Commit: feat(test): add route matching

# Commit: refactor(docs): improve documentation

# Commit: test(test): add coverage for response writing

# Commit: feat(docs): add middleware chain

# Commit: feat(middleware): implement param extraction

# Commit: refactor(di): restructure documentation

# Commit: fix(di-container): handle routing conflict case

# Commit: refactor(route): restructure test coverage

# Commit: test(http): add coverage for request injection

# Commit: fix(ioc): resolve type inference

# Commit: test(container): add coverage for request injection

# Commit: fix(module): resolve path extraction

# Commit: feat(controller): implement singleton scope

# Commit: feat(router): add error handling

# Commit: fix(server): handle path extraction case

# Commit: fix(provider): handle path extraction case

# Commit: feat(core): implement handler resolution

# Commit: test(example): add coverage for param extraction

# Commit: perf(hello): optimize handler resolution

# Commit: test(test): add coverage for controller routing

# Commit: chore(docs): update go mod

# Commit: feat(middleware): implement route matching

# Commit: fix(docs): handle body parsing case

# Commit: chore(middleware): update go mod

# Commit: feat(di): implement middleware chain

# Commit: feat(di-container): implement middleware chain

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add route matching

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update test suite

# Commit: test(router): add coverage for request injection

# Commit: fix(server): resolve routing conflict

# Commit: fix(provider): resolve pattern matching

# Commit: docs(core): update documentation

# Commit: feat(example): add route matching

# Commit: refactor(hello): improve type safety

# Commit: perf(test): optimize param extraction

# Commit: feat(docs): add singleton scope

# Commit: feat(middleware): add context propagation

# Commit: refactor(di): restructure type safety

# Commit: refactor(middleware): improve concurrency handling

# Commit: feat(di): implement middleware chain

# Commit: feat(di-container): implement context propagation

# Commit: refactor(route): restructure error messages

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for middleware chain

# Commit: feat(container): implement request injection

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update test suite

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle header setting case

# Commit: perf(provider): optimize provider registration

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: refactor(hello): improve test coverage

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize route matching

# Commit: refactor(middleware): improve concurrency handling

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize param extraction

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): restructure error messages

# Commit: fix(route): handle scope resolution case

# Commit: refactor(http): restructure type safety

# Commit: refactor(ioc): restructure error messages

# Commit: refactor(container): improve performance

# Commit: feat(module): add context propagation

# Commit: refactor(controller): restructure error messages

# Commit: refactor(router): improve concurrency handling

# Commit: refactor(server): restructure error messages

# Commit: refactor(provider): improve memory usage

# Commit: chore(core): update build script

# Commit: refactor(example): improve code structure

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize middleware chain

# Commit: feat(docs): implement module setup

# Commit: perf(middleware): optimize provider registration

# Commit: docs(di): update documentation

# Commit: feat(di-container): implement module setup

# Commit: perf(route): optimize module setup

# Commit: fix(di-container): handle scope resolution case

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update test suite

# Commit: test(ioc): add coverage for error handling

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve performance

# Commit: feat(controller): add middleware chain

# Commit: feat(router): implement error handling

# Commit: fix(server): handle scope resolution case

# Commit: refactor(provider): improve error messages

# Commit: refactor(core): improve performance

# Commit: feat(example): add provider registration

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for context propagation

# Commit: refactor(di): restructure code structure

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): resolve nil pointer

# Commit: feat(http): add response writing

# Commit: fix(route): resolve path extraction

# Commit: feat(http): implement param extraction

# Commit: fix(ioc): handle nil pointer case

# Commit: feat(container): add singleton scope

# Commit: feat(module): add context propagation

# Commit: refactor(controller): restructure memory usage

# Commit: chore(router): update readme

# Commit: docs(server): update documentation

# Commit: fix(provider): handle body parsing case

# Commit: chore(core): update build script

# Commit: refactor(example): improve performance

# Commit: feat(hello): add response writing

# Commit: refactor(test): improve documentation

# Commit: docs(docs): update documentation

# Commit: fix(middleware): resolve type inference

# Commit: test(di): add coverage for provider registration

# Commit: docs(di-container): update documentation

# Commit: perf(route): optimize param extraction

# Commit: test(http): add coverage for request injection

# Commit: fix(ioc): resolve routing conflict

# Commit: refactor(http): restructure error messages

# Commit: refactor(ioc): improve error messages

# Commit: refactor(container): improve code structure

# Commit: fix(module): resolve nil pointer

# Commit: feat(controller): add singleton scope

# Commit: feat(router): add param extraction

# Commit: test(server): add coverage for context propagation

# Commit: refactor(provider): restructure concurrency handling

# Commit: feat(core): implement controller routing

# Commit: chore(example): update license

# Commit: perf(hello): optimize request injection

# Commit: fix(test): resolve type inference

# Commit: test(docs): add coverage for module setup

# Commit: perf(middleware): optimize context propagation

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize provider registration

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure memory usage

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update readme

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement response writing

# Commit: feat(module): add context propagation

# Commit: refactor(controller): restructure documentation

# Commit: docs(router): update documentation

# Commit: refactor(server): improve error messages

# Commit: refactor(provider): improve error messages

# Commit: refactor(core): improve documentation

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure documentation

# Commit: docs(test): update documentation

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add response writing

# Commit: chore(di): update dependencies

# Commit: feat(di-container): add error handling

# Commit: fix(route): handle header setting case

# Commit: fix(http): handle type inference case

# Commit: test(ioc): add coverage for controller routing

# Commit: chore(container): update build script

# Commit: refactor(module): improve performance

# Commit: perf(container): optimize module setup

# Commit: perf(module): optimize handler resolution

# Commit: test(controller): add coverage for singleton scope

# Commit: feat(router): add middleware chain

# Commit: feat(server): implement handler resolution

# Commit: test(provider): add coverage for module setup

# Commit: perf(core): optimize response writing

# Commit: chore(example): update go mod

# Commit: feat(hello): implement module setup

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure type safety

# Commit: perf(middleware): optimize module setup

# Commit: perf(di): optimize provider registration

# Commit: docs(di-container): update documentation

# Commit: refactor(route): restructure concurrency handling

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure performance

# Commit: feat(container): add request injection

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve error messages

# Commit: feat(module): implement param extraction

# Commit: chore(controller): update ci configuration

# Commit: fix(router): handle nil pointer case

# Commit: feat(server): add param extraction

# Commit: test(provider): add coverage for route matching

# Commit: refactor(core): improve code structure

# Commit: fix(example): resolve body parsing

# Commit: fix(hello): handle header setting case

# Commit: perf(test): optimize param extraction

# Commit: fix(docs): handle nil pointer case

# Commit: feat(middleware): add context propagation

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): improve code structure

# Commit: fix(route): resolve routing conflict

# Commit: fix(http): resolve type inference

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add request injection

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update build script

# Commit: refactor(router): improve type safety

# Commit: fix(controller): handle body parsing case

# Commit: chore(router): update build script

# Commit: refactor(server): improve test coverage

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): improve memory usage

# Commit: chore(example): update license

# Commit: perf(hello): optimize route matching

# Commit: refactor(test): improve performance

# Commit: feat(docs): add controller routing

# Commit: chore(middleware): update dependencies

# Commit: feat(di): add provider registration

# Commit: docs(di-container): update documentation

# Commit: refactor(route): improve type safety

# Commit: perf(http): optimize route matching

# Commit: refactor(ioc): improve performance

# Commit: fix(container): handle type inference case

# Commit: test(module): add coverage for response writing

# Commit: perf(controller): optimize singleton scope

# Commit: feat(router): add singleton scope

# Commit: feat(server): add singleton scope

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve performance

# Commit: feat(provider): add param extraction

# Commit: refactor(core): restructure type safety

# Commit: perf(example): optimize error handling

# Commit: fix(hello): handle body parsing case

# Commit: chore(test): update test suite

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle pattern matching case

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize middleware chain

# Commit: feat(route): implement response writing

# Commit: feat(http): add middleware chain

# Commit: feat(ioc): implement error handling

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve code structure

# Commit: fix(controller): resolve header setting

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure error messages

# Commit: refactor(provider): improve error messages

# Commit: feat(server): add response writing

# Commit: refactor(provider): restructure concurrency handling

# Commit: feat(core): implement error handling

# Commit: fix(example): handle body parsing case

# Commit: chore(hello): update readme

# Commit: docs(test): update documentation

# Commit: feat(docs): implement context propagation

# Commit: refactor(middleware): restructure concurrency handling

# Commit: fix(di): handle nil pointer case

# Commit: feat(di-container): add provider registration

# Commit: docs(route): update documentation

# Commit: test(http): add coverage for error handling

# Commit: fix(ioc): handle scope resolution case

# Commit: fix(container): handle type inference case

# Commit: refactor(module): restructure error messages

# Commit: refactor(controller): restructure type safety

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure test coverage

# Commit: test(provider): add coverage for error handling

# Commit: fix(core): handle header setting case

# Commit: chore(provider): update test suite

# Commit: refactor(core): restructure code structure

# Commit: fix(example): resolve type inference

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize handler resolution

# Commit: test(docs): add coverage for response writing

# Commit: docs(middleware): update documentation

# Commit: fix(di): handle body parsing case

# Commit: fix(di-container): handle type inference case

# Commit: test(route): add coverage for handler resolution

# Commit: test(http): add coverage for route matching

# Commit: refactor(ioc): improve documentation

# Commit: docs(container): update documentation

# Commit: feat(module): implement module setup

# Commit: perf(controller): optimize param extraction

# Commit: docs(router): update documentation

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: feat(core): implement controller routing

# Commit: chore(example): update license

# Commit: feat(core): add provider registration

# Commit: docs(example): update documentation

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize route matching

# Commit: refactor(docs): improve concurrency handling

# Commit: feat(middleware): implement module setup

# Commit: perf(di): optimize param extraction

# Commit: fix(di-container): resolve header setting

# Commit: perf(route): optimize middleware chain

# Commit: feat(http): implement module setup

# Commit: perf(ioc): optimize context propagation

# Commit: refactor(container): restructure code structure

# Commit: fix(module): handle nil pointer case

# Commit: refactor(controller): restructure code structure

# Commit: fix(router): resolve body parsing

# Commit: chore(server): update go mod

# Commit: fix(provider): handle pattern matching case

# Commit: docs(core): update documentation

# Commit: refactor(example): improve concurrency handling

# Commit: feat(hello): implement module setup

# Commit: test(example): add coverage for handler resolution

# Commit: test(hello): add coverage for response writing

# Commit: refactor(test): restructure error messages

# Commit: refactor(docs): improve performance

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize response writing

# Commit: docs(di-container): update documentation

# Commit: refactor(route): improve code structure

# Commit: fix(http): resolve type inference

# Commit: test(ioc): add coverage for controller routing

# Commit: chore(container): update build script

# Commit: refactor(module): improve performance

# Commit: feat(controller): add context propagation

# Commit: refactor(router): restructure memory usage

# Commit: refactor(server): restructure performance

# Commit: feat(provider): add param extraction

# Commit: test(core): add coverage for route matching

# Commit: refactor(example): improve performance

# Commit: feat(hello): add controller routing

# Commit: chore(test): update readme

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve nil pointer

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize param extraction

# Commit: refactor(di): restructure test coverage

# Commit: test(di-container): add coverage for param extraction

# Commit: feat(route): implement handler resolution

# Commit: test(http): add coverage for middleware chain

# Commit: feat(ioc): implement request injection

# Commit: fix(container): resolve routing conflict

# Commit: refactor(module): restructure performance

# Commit: feat(controller): add module setup

# Commit: perf(router): optimize param extraction

# Commit: refactor(server): improve memory usage

# Commit: refactor(provider): restructure type safety

# Commit: fix(core): handle type inference case

# Commit: test(example): add coverage for singleton scope

# Commit: feat(hello): add middleware chain

# Commit: feat(test): implement provider registration

# Commit: docs(docs): update documentation

# Commit: refactor(test): restructure error messages

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle scope resolution case

# Commit: refactor(di-container): improve documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure performance

# Commit: refactor(ioc): restructure memory usage

# Commit: refactor(container): restructure memory usage

# Commit: fix(module): handle routing conflict case

# Commit: fix(controller): resolve routing conflict

# Commit: fix(router): resolve header setting

# Commit: perf(server): optimize middleware chain

# Commit: feat(provider): implement module setup

# Commit: perf(core): optimize param extraction

# Commit: refactor(example): improve error messages

# Commit: fix(hello): handle type inference case

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve scope resolution

# Commit: fix(middleware): handle routing conflict case

# Commit: fix(docs): handle path extraction case

# Commit: feat(middleware): implement provider registration

# Commit: docs(di): update documentation

# Commit: feat(di-container): implement response writing

# Commit: fix(route): handle pattern matching case

# Commit: docs(http): update documentation

# Commit: perf(ioc): optimize handler resolution

# Commit: test(container): add coverage for param extraction

# Commit: perf(module): optimize request injection

# Commit: fix(controller): resolve routing conflict

# Commit: fix(router): handle routing conflict case

# Commit: fix(server): resolve nil pointer

# Commit: feat(provider): add provider registration

# Commit: docs(core): update documentation

# Commit: refactor(example): restructure memory usage

# Commit: chore(hello): update readme

# Commit: docs(test): update documentation

# Commit: refactor(docs): restructure concurrency handling

# Commit: refactor(middleware): restructure code structure

# Commit: fix(di): resolve scope resolution

# Commit: chore(middleware): update readme

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): restructure code structure

# Commit: fix(ioc): resolve scope resolution

# Commit: refactor(container): improve test coverage

# Commit: refactor(module): restructure type safety

# Commit: perf(controller): optimize singleton scope

# Commit: feat(router): add middleware chain

# Commit: feat(server): implement error handling

# Commit: fix(provider): handle pattern matching case

# Commit: fix(core): handle nil pointer case

# Commit: feat(example): add controller routing

# Commit: chore(hello): update dependencies

# Commit: feat(test): add context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for handler resolution

# Commit: test(di): add coverage for error handling

# Commit: fix(di-container): handle nil pointer case

# Commit: chore(di): update ci configuration

# Commit: fix(di-container): resolve routing conflict

# Commit: refactor(route): restructure documentation

# Commit: fix(http): handle routing conflict case

# Commit: fix(ioc): resolve nil pointer

# Commit: fix(container): handle type inference case

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize middleware chain

# Commit: feat(router): implement response writing

# Commit: fix(server): handle path extraction case

# Commit: feat(provider): implement provider registration

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: fix(hello): handle path extraction case

# Commit: fix(test): handle pattern matching case

# Commit: refactor(docs): restructure documentation

# Commit: docs(middleware): update documentation

# Commit: fix(di): resolve nil pointer

# Commit: feat(di-container): add response writing

# Commit: feat(route): implement request injection

# Commit: docs(di-container): update documentation

# Commit: test(route): add coverage for context propagation

# Commit: refactor(http): restructure test coverage

# Commit: test(ioc): add coverage for controller routing

# Commit: chore(container): update build script

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update gitignore

# Commit: chore(router): update go mod

# Commit: refactor(server): restructure documentation

# Commit: docs(provider): update documentation

# Commit: feat(core): add param extraction

# Commit: feat(example): implement singleton scope

# Commit: feat(hello): add context propagation

# Commit: refactor(test): restructure error messages

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): resolve path extraction

# Commit: feat(di): implement handler resolution

# Commit: test(di-container): add coverage for singleton scope

# Commit: feat(route): add response writing

# Commit: chore(http): update ci configuration

# Commit: chore(route): update gitignore

# Commit: chore(http): update build script

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement response writing

# Commit: test(module): add coverage for route matching

# Commit: refactor(controller): improve code structure

# Commit: fix(router): resolve routing conflict

# Commit: fix(server): resolve path extraction

# Commit: feat(provider): implement route matching

# Commit: refactor(core): improve test coverage

# Commit: test(example): add coverage for middleware chain

# Commit: feat(hello): implement route matching

# Commit: refactor(test): improve performance

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize singleton scope

# Commit: feat(di): add provider registration

# Commit: docs(di-container): update documentation

# Commit: chore(route): update license

# Commit: perf(http): optimize middleware chain

# Commit: feat(ioc): implement provider registration

# Commit: fix(http): resolve type inference

# Commit: refactor(ioc): restructure code structure

# Commit: fix(container): handle routing conflict case

# Commit: fix(module): handle header setting case

# Commit: perf(controller): optimize provider registration

# Commit: docs(router): update documentation

# Commit: refactor(server): restructure performance

# Commit: feat(provider): add param extraction

# Commit: feat(core): add handler resolution

# Commit: test(example): add coverage for route matching

# Commit: refactor(hello): improve test coverage

# Commit: test(test): add coverage for context propagation

# Commit: refactor(docs): restructure type safety

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle nil pointer case

# Commit: feat(di-container): add singleton scope

# Commit: feat(route): add controller routing

# Commit: chore(http): update readme

# Commit: docs(ioc): update documentation

# Commit: refactor(container): restructure error messages

# Commit: docs(ioc): update documentation

# Commit: fix(container): resolve body parsing

# Commit: chore(module): update test suite

# Commit: test(controller): add coverage for module setup

# Commit: perf(router): optimize controller routing

# Commit: chore(server): update dependencies

# Commit: feat(provider): add context propagation

# Commit: refactor(core): restructure documentation

# Commit: docs(example): update documentation

# Commit: feat(hello): implement controller routing

# Commit: chore(test): update go mod

# Commit: feat(docs): implement handler resolution

# Commit: test(middleware): add coverage for singleton scope

# Commit: feat(di): add provider registration

# Commit: docs(di-container): update documentation

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update readme

# Commit: refactor(ioc): restructure performance

# Commit: feat(container): add controller routing

# Commit: chore(module): update license

# Commit: refactor(container): restructure concurrency handling

# Commit: feat(module): implement route matching

# Commit: refactor(controller): improve type safety

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: refactor(provider): improve performance

# Commit: feat(core): add controller routing

# Commit: chore(example): update ci configuration

# Commit: fix(hello): resolve path extraction

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure error messages

# Commit: refactor(middleware): improve code structure

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize provider registration

# Commit: docs(route): update documentation

# Commit: chore(http): update build script

# Commit: refactor(ioc): restructure test coverage

# Commit: refactor(container): restructure performance

# Commit: refactor(module): restructure memory usage

# Commit: chore(controller): update dependencies

# Commit: refactor(module): improve type safety

# Commit: perf(controller): optimize response writing

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve concurrency handling

# Commit: fix(provider): handle nil pointer case

# Commit: feat(core): add module setup

# Commit: perf(example): optimize controller routing

# Commit: chore(hello): update go mod

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure memory usage

# Commit: chore(middleware): update gitignore

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize module setup

# Commit: perf(route): optimize controller routing

# Commit: chore(http): update readme

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize handler resolution

# Commit: test(module): add coverage for route matching

# Commit: refactor(controller): improve concurrency handling

# Commit: feat(router): implement request injection

# Commit: perf(controller): optimize error handling

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve error messages

# Commit: refactor(provider): improve concurrency handling

# Commit: feat(core): implement controller routing

# Commit: chore(example): update dependencies

# Commit: feat(hello): add route matching

# Commit: refactor(test): improve performance

# Commit: feat(docs): add singleton scope

# Commit: feat(middleware): add context propagation

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): improve test coverage

# Commit: test(route): add coverage for param extraction

# Commit: feat(http): add response writing

# Commit: refactor(ioc): restructure error messages

# Commit: refactor(container): improve memory usage

# Commit: chore(module): update dependencies

# Commit: feat(controller): add controller routing

# Commit: chore(router): update go mod

# Commit: feat(server): implement middleware chain

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: test(core): add coverage for module setup

# Commit: perf(example): optimize provider registration

# Commit: docs(hello): update documentation

# Commit: fix(test): handle body parsing case

# Commit: chore(docs): update test suite

# Commit: fix(middleware): handle type inference case

# Commit: test(di): add coverage for module setup

# Commit: perf(di-container): optimize param extraction

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: chore(ioc): update build script

# Commit: refactor(container): restructure memory usage

# Commit: chore(module): update readme

# Commit: docs(controller): update documentation

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve memory usage

# Commit: chore(provider): update dependencies

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement middleware chain

# Commit: feat(core): implement route matching

# Commit: refactor(example): improve error messages

# Commit: refactor(hello): restructure performance

# Commit: feat(test): add error handling

# Commit: fix(docs): handle path extraction case

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve pattern matching

# Commit: docs(route): update documentation

# Commit: docs(http): update documentation

# Commit: feat(ioc): add middleware chain

# Commit: feat(container): implement error handling

# Commit: fix(module): handle routing conflict case

# Commit: fix(controller): resolve path extraction

# Commit: feat(router): implement module setup

# Commit: perf(server): optimize handler resolution

# Commit: test(provider): add coverage for error handling

# Commit: fix(core): handle pattern matching case

# Commit: refactor(provider): restructure performance

# Commit: feat(core): add param extraction

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for singleton scope

# Commit: feat(di): add module setup

# Commit: perf(di-container): optimize middleware chain

# Commit: feat(route): implement request injection

# Commit: fix(http): resolve path extraction

# Commit: fix(ioc): handle path extraction case

# Commit: feat(container): implement singleton scope

# Commit: feat(module): add handler resolution

# Commit: test(controller): add coverage for route matching

# Commit: refactor(router): improve performance

# Commit: feat(server): add error handling

# Commit: fix(provider): handle header setting case

# Commit: perf(core): optimize response writing

# Commit: chore(example): update license

# Commit: docs(core): update documentation

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update license

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement handler resolution

# Commit: test(di): add coverage for singleton scope

# Commit: feat(di-container): add singleton scope

# Commit: feat(route): add provider registration

# Commit: docs(http): update documentation

# Commit: fix(ioc): resolve path extraction

# Commit: feat(container): implement response writing

# Commit: refactor(module): restructure performance

# Commit: fix(controller): handle scope resolution case

# Commit: refactor(router): improve concurrency handling

# Commit: feat(server): implement handler resolution

# Commit: test(provider): add coverage for controller routing

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve routing conflict

# Commit: refactor(hello): restructure concurrency handling

# Commit: fix(example): handle body parsing case

# Commit: chore(hello): update dependencies

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve header setting

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle pattern matching case

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure code structure

# Commit: fix(ioc): resolve header setting

# Commit: refactor(container): restructure memory usage

# Commit: chore(module): update ci configuration

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add response writing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement singleton scope

# Commit: feat(core): add singleton scope

# Commit: feat(example): add request injection

# Commit: fix(hello): resolve type inference

# Commit: refactor(test): restructure performance

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize param extraction

# Commit: fix(docs): resolve header setting

# Commit: perf(middleware): optimize response writing

# Commit: chore(di): update go mod

# Commit: fix(di-container): handle nil pointer case

# Commit: feat(route): add module setup

# Commit: perf(http): optimize singleton scope

# Commit: feat(ioc): add module setup

# Commit: perf(container): optimize middleware chain

# Commit: feat(module): implement handler resolution

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for controller routing

# Commit: chore(server): update ci configuration

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): resolve scope resolution

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for singleton scope

# Commit: feat(test): add context propagation

# Commit: refactor(docs): restructure error messages

# Commit: fix(test): resolve routing conflict

# Commit: fix(docs): resolve pattern matching

# Commit: docs(middleware): update documentation

# Commit: feat(di): implement response writing

# Commit: chore(di-container): update gitignore

# Commit: chore(route): update test suite

# Commit: test(http): add coverage for handler resolution

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve body parsing

# Commit: chore(module): update go mod

# Commit: refactor(controller): restructure type safety

# Commit: perf(router): optimize request injection

# Commit: fix(server): resolve body parsing

# Commit: refactor(provider): restructure type safety

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve documentation

# Commit: docs(hello): update documentation

# Commit: refactor(test): restructure error messages

# Commit: refactor(docs): restructure test coverage

# Commit: refactor(middleware): restructure documentation

# Commit: docs(docs): update documentation

# Commit: docs(middleware): update documentation

# Commit: fix(di): resolve header setting

# Commit: fix(di-container): handle type inference case

# Commit: test(route): add coverage for provider registration

# Commit: docs(http): update documentation

# Commit: docs(ioc): update documentation

# Commit: feat(container): implement request injection

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve type safety

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve documentation

# Commit: docs(provider): update documentation

# Commit: perf(core): optimize param extraction

# Commit: feat(example): implement singleton scope

# Commit: feat(hello): add provider registration

# Commit: docs(test): update documentation

# Commit: fix(docs): resolve path extraction

# Commit: feat(middleware): implement middleware chain

# Commit: feat(di): implement middleware chain

# Commit: fix(middleware): resolve pattern matching

# Commit: docs(di): update documentation

# Commit: chore(di-container): update dependencies

# Commit: feat(route): add provider registration

# Commit: docs(http): update documentation

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: fix(module): resolve pattern matching

# Commit: fix(controller): handle scope resolution case

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update ci configuration

# Commit: fix(provider): resolve path extraction

# Commit: feat(core): implement handler resolution

# Commit: test(example): add coverage for controller routing

# Commit: chore(hello): update test suite

# Commit: test(test): add coverage for error handling

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: fix(di): resolve routing conflict

# Commit: fix(di-container): resolve body parsing

# Commit: fix(di): resolve pattern matching

# Commit: docs(di-container): update documentation

# Commit: refactor(route): restructure test coverage

# Commit: test(http): add coverage for controller routing

# Commit: chore(ioc): update ci configuration

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for middleware chain

# Commit: feat(controller): implement singleton scope

# Commit: feat(router): add error handling

# Commit: fix(server): handle path extraction case

# Commit: feat(provider): implement response writing

# Commit: docs(core): update documentation

# Commit: feat(example): add param extraction

# Commit: chore(hello): update ci configuration

# Commit: fix(test): resolve body parsing

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: chore(di): update gitignore

# Commit: refactor(di-container): restructure memory usage

# Commit: chore(route): update gitignore

# Commit: feat(di-container): implement module setup

# Commit: perf(route): optimize provider registration

# Commit: docs(http): update documentation

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement singleton scope

# Commit: feat(module): add singleton scope

# Commit: feat(controller): add error handling

# Commit: fix(router): handle nil pointer case

# Commit: feat(server): add response writing

# Commit: fix(provider): resolve path extraction

# Commit: feat(core): implement singleton scope

# Commit: feat(example): add request injection

# Commit: fix(hello): resolve type inference

# Commit: test(test): add coverage for request injection

# Commit: fix(docs): resolve routing conflict

# Commit: fix(middleware): resolve nil pointer

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): improve memory usage

# Commit: chore(route): update build script

# Commit: refactor(http): restructure error messages

# Commit: docs(route): update documentation

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add singleton scope

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for context propagation

# Commit: refactor(server): restructure code structure

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize param extraction

# Commit: refactor(test): improve code structure

# Commit: fix(docs): resolve path extraction

# Commit: feat(middleware): implement param extraction

# Commit: test(di): add coverage for route matching

# Commit: refactor(di-container): improve code structure

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): improve concurrency handling

# Commit: fix(ioc): handle type inference case

# Commit: fix(http): handle type inference case

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for singleton scope

# Commit: feat(controller): add param extraction

# Commit: chore(router): update gitignore

# Commit: chore(server): update go mod

# Commit: refactor(provider): restructure concurrency handling

# Commit: feat(core): implement middleware chain

# Commit: feat(example): implement context propagation

# Commit: refactor(hello): restructure documentation

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle type inference case

# Commit: test(di-container): add coverage for middleware chain

# Commit: feat(route): implement param extraction

# Commit: feat(http): add response writing

# Commit: refactor(ioc): restructure documentation

# Commit: docs(container): update documentation

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for response writing

# Commit: fix(module): resolve body parsing

# Commit: refactor(controller): restructure test coverage

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize param extraction

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add handler resolution

# Commit: test(middleware): add coverage for module setup

# Commit: perf(di): optimize handler resolution

# Commit: test(di-container): add coverage for singleton scope

# Commit: feat(route): add response writing

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement handler resolution

# Commit: test(module): add coverage for provider registration

# Commit: feat(container): add middleware chain

# Commit: feat(module): implement error handling

# Commit: fix(controller): handle type inference case

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle routing conflict case

# Commit: fix(provider): resolve type inference

# Commit: test(core): add coverage for middleware chain

# Commit: feat(example): implement error handling

# Commit: fix(hello): handle nil pointer case

# Commit: feat(test): add error handling

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: test(di): add coverage for context propagation

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement module setup

# Commit: perf(http): optimize param extraction

# Commit: feat(ioc): implement context propagation

# Commit: refactor(container): restructure type safety

# Commit: perf(module): optimize response writing

# Commit: perf(controller): optimize provider registration

# Commit: perf(module): optimize middleware chain

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure type safety

# Commit: refactor(server): restructure test coverage

# Commit: test(provider): add coverage for middleware chain

# Commit: feat(core): implement context propagation

# Commit: refactor(example): restructure error messages

# Commit: refactor(hello): improve performance

# Commit: feat(test): add request injection

# Commit: fix(docs): resolve type inference

# Commit: fix(middleware): handle scope resolution case

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize module setup

# Commit: perf(route): optimize singleton scope

# Commit: feat(http): add handler resolution

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: feat(module): implement singleton scope

# Commit: feat(controller): add singleton scope

# Commit: feat(router): add context propagation

# Commit: chore(controller): update ci configuration

# Commit: fix(router): resolve path extraction

# Commit: feat(server): implement handler resolution

# Commit: test(provider): add coverage for request injection

# Commit: fix(core): resolve header setting

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for context propagation

# Commit: refactor(test): restructure type safety

# Commit: perf(docs): optimize singleton scope

# Commit: feat(middleware): add singleton scope

# Commit: feat(di): add middleware chain

# Commit: feat(di-container): implement singleton scope

# Commit: feat(route): add module setup

# Commit: perf(http): optimize context propagation

# Commit: refactor(ioc): restructure memory usage

# Commit: fix(container): handle type inference case

# Commit: refactor(module): restructure concurrency handling

# Commit: refactor(controller): restructure code structure

# Commit: fix(router): resolve header setting

# Commit: perf(server): optimize module setup

# Commit: feat(router): implement route matching

# Commit: refactor(server): improve concurrency handling

# Commit: feat(provider): implement route matching

# Commit: refactor(core): improve performance

# Commit: refactor(example): restructure concurrency handling

# Commit: feat(hello): implement error handling

# Commit: fix(test): handle pattern matching case

# Commit: refactor(docs): restructure type safety

# Commit: perf(middleware): optimize route matching

# Commit: refactor(di): improve concurrency handling

# Commit: feat(di-container): implement context propagation

# Commit: refactor(route): restructure concurrency handling

# Commit: feat(http): implement route matching

# Commit: refactor(ioc): improve test coverage

# Commit: test(container): add coverage for param extraction

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: docs(router): update documentation

# Commit: fix(server): handle type inference case

# Commit: test(provider): add coverage for controller routing

# Commit: feat(server): add singleton scope

# Commit: feat(provider): add singleton scope

# Commit: feat(core): add module setup

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize provider registration

# Commit: docs(test): update documentation

# Commit: docs(docs): update documentation

# Commit: fix(middleware): resolve header setting

# Commit: perf(di): optimize handler resolution

# Commit: test(di-container): add coverage for response writing

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add context propagation

# Commit: refactor(container): restructure concurrency handling

# Commit: feat(module): implement provider registration

# Commit: docs(controller): update documentation

# Commit: perf(router): optimize route matching

# Commit: refactor(server): improve code structure

# Commit: fix(provider): handle path extraction case

# Commit: feat(core): implement provider registration

# Commit: docs(provider): update documentation

# Commit: test(core): add coverage for route matching

# Commit: refactor(example): improve code structure

# Commit: fix(hello): resolve type inference

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add param extraction

# Commit: chore(middleware): update gitignore

# Commit: refactor(di): restructure concurrency handling

# Commit: feat(di-container): implement controller routing

# Commit: chore(route): update license

# Commit: perf(http): optimize context propagation

# Commit: refactor(ioc): restructure memory usage

# Commit: refactor(container): restructure code structure

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update dependencies

# Commit: feat(router): add error handling

# Commit: fix(server): handle nil pointer case

# Commit: feat(provider): add request injection

# Commit: fix(core): resolve nil pointer

# Commit: feat(example): add route matching

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve routing conflict

# Commit: fix(hello): resolve scope resolution

# Commit: refactor(test): restructure documentation

# Commit: docs(docs): update documentation

# Commit: test(middleware): add coverage for provider registration

# Commit: docs(di): update documentation

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement response writing

# Commit: perf(http): optimize provider registration

# Commit: docs(ioc): update documentation

# Commit: feat(container): implement provider registration

# Commit: docs(module): update documentation

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve code structure

# Commit: fix(provider): resolve scope resolution

# Commit: refactor(core): improve documentation

# Commit: refactor(example): restructure error messages

# Commit: fix(hello): handle header setting case

# Commit: fix(example): handle header setting case

# Commit: fix(hello): handle body parsing case

# Commit: chore(test): update go mod

# Commit: refactor(docs): restructure memory usage

# Commit: chore(middleware): update dependencies

# Commit: fix(di): handle body parsing case

# Commit: chore(di-container): update license

# Commit: perf(route): optimize request injection

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add singleton scope

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: perf(controller): optimize controller routing

# Commit: chore(router): update readme

# Commit: refactor(server): restructure code structure

# Commit: fix(provider): resolve header setting

# Commit: fix(core): handle type inference case

# Commit: fix(example): handle header setting case

# Commit: perf(hello): optimize error handling

# Commit: fix(test): handle pattern matching case

# Commit: docs(hello): update documentation

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize provider registration

# Commit: docs(middleware): update documentation

# Commit: chore(di): update license

# Commit: perf(di-container): optimize error handling

# Commit: fix(route): handle type inference case

# Commit: test(http): add coverage for handler resolution

# Commit: test(ioc): add coverage for middleware chain

# Commit: feat(container): implement provider registration

# Commit: docs(module): update documentation

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for handler resolution

# Commit: test(server): add coverage for handler resolution

# Commit: test(provider): add coverage for error handling

# Commit: fix(core): handle path extraction case

# Commit: refactor(example): restructure performance

# Commit: feat(hello): add error handling

# Commit: fix(test): handle header setting case

# Commit: perf(docs): optimize controller routing

# Commit: chore(test): update readme

# Commit: docs(docs): update documentation

# Commit: docs(middleware): update documentation

# Commit: fix(di): resolve scope resolution

# Commit: refactor(di-container): improve code structure

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for provider registration

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve concurrency handling

# Commit: fix(module): handle scope resolution case

# Commit: refactor(controller): improve documentation

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: docs(core): update documentation

# Commit: perf(example): optimize param extraction

# Commit: test(hello): add coverage for param extraction

# Commit: chore(test): update ci configuration

# Commit: fix(docs): resolve pattern matching

# Commit: docs(middleware): update documentation

# Commit: refactor(docs): restructure memory usage

# Commit: chore(middleware): update go mod

# Commit: feat(di): implement response writing

# Commit: test(di-container): add coverage for route matching

# Commit: refactor(route): improve error messages

# Commit: refactor(http): improve error messages

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement controller routing

# Commit: chore(module): update license

# Commit: perf(controller): optimize context propagation

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update readme

# Commit: docs(provider): update documentation

# Commit: feat(core): add module setup

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure code structure

# Commit: fix(test): resolve routing conflict

# Commit: fix(docs): handle path extraction case

# Commit: feat(middleware): implement module setup

# Commit: perf(di): optimize error handling

# Commit: test(middleware): add coverage for controller routing

# Commit: chore(di): update go mod

# Commit: feat(di-container): implement request injection

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update license

# Commit: perf(ioc): optimize response writing

# Commit: docs(container): update documentation

# Commit: docs(module): update documentation

# Commit: feat(controller): implement request injection

# Commit: fix(router): resolve body parsing

# Commit: chore(server): update license

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve body parsing

# Commit: chore(example): update test suite

# Commit: test(hello): add coverage for middleware chain

# Commit: feat(test): implement controller routing

# Commit: chore(docs): update test suite

# Commit: test(middleware): add coverage for module setup

# Commit: perf(di): optimize error handling

# Commit: fix(di-container): handle header setting case

# Commit: fix(di): resolve nil pointer

# Commit: fix(di-container): handle header setting case

# Commit: perf(route): optimize error handling

# Commit: fix(http): handle nil pointer case

# Commit: feat(ioc): add param extraction

# Commit: fix(container): resolve routing conflict

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve performance

# Commit: feat(router): add singleton scope

# Commit: feat(server): add middleware chain

# Commit: feat(provider): implement handler resolution

# Commit: test(core): add coverage for context propagation

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): handle body parsing case

# Commit: chore(test): update test suite

# Commit: fix(docs): handle scope resolution case

# Commit: refactor(middleware): improve documentation

# Commit: docs(di): update documentation

# Commit: chore(di-container): update gitignore

# Commit: chore(route): update license

# Commit: docs(di-container): update documentation

# Commit: chore(route): update test suite

# Commit: test(http): add coverage for request injection

# Commit: fix(ioc): resolve header setting

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle header setting case

# Commit: perf(controller): optimize context propagation

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update dependencies

# Commit: feat(provider): add param extraction

# Commit: fix(core): handle nil pointer case

# Commit: feat(example): add request injection

# Commit: fix(hello): resolve nil pointer

# Commit: feat(test): add route matching

# Commit: refactor(docs): improve memory usage

# Commit: chore(middleware): update test suite

# Commit: test(di): add coverage for response writing

# Commit: fix(di-container): resolve scope resolution

# Commit: refactor(route): improve error messages

# Commit: refactor(http): improve type safety

# Commit: refactor(route): improve type safety

# Commit: refactor(http): restructure documentation

# Commit: docs(ioc): update documentation

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve test coverage

# Commit: test(controller): add coverage for request injection

# Commit: fix(router): resolve type inference

# Commit: test(server): add coverage for singleton scope

# Commit: feat(provider): add provider registration

# Commit: docs(core): update documentation

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for param extraction

# Commit: refactor(test): improve type safety

# Commit: perf(docs): optimize route matching

# Commit: refactor(middleware): improve type safety

# Commit: perf(di): optimize param extraction

# Commit: feat(di-container): add middleware chain

# Commit: feat(route): implement middleware chain

# Commit: feat(http): implement handler resolution

# Commit: test(ioc): add coverage for request injection

# Commit: chore(http): update license

# Commit: perf(ioc): optimize controller routing

# Commit: chore(container): update build script

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: refactor(router): improve error messages

# Commit: refactor(server): improve type safety

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update go mod

# Commit: fix(example): handle type inference case

# Commit: test(hello): add coverage for request injection

# Commit: fix(test): resolve scope resolution

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle pattern matching case

# Commit: refactor(di-container): restructure type safety

# Commit: perf(route): optimize response writing

# Commit: test(http): add coverage for module setup

# Commit: perf(ioc): optimize param extraction

# Commit: refactor(container): improve memory usage

# Commit: refactor(ioc): improve error messages

# Commit: refactor(container): improve error messages

# Commit: fix(module): handle nil pointer case

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize controller routing

# Commit: chore(server): update readme

# Commit: refactor(provider): restructure code structure

# Commit: fix(core): resolve type inference

# Commit: test(example): add coverage for context propagation

# Commit: refactor(hello): restructure memory usage

# Commit: chore(test): update readme

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): restructure type safety

# Commit: refactor(di): restructure concurrency handling

# Commit: feat(di-container): implement route matching

# Commit: refactor(route): improve documentation

# Commit: fix(http): handle pattern matching case

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize singleton scope

# Commit: feat(module): add error handling

# Commit: chore(container): update gitignore

# Commit: fix(module): handle path extraction case

# Commit: fix(controller): handle scope resolution case

# Commit: refactor(router): restructure type safety

# Commit: perf(server): optimize context propagation

# Commit: refactor(provider): restructure memory usage

# Commit: fix(core): handle body parsing case

# Commit: fix(example): handle routing conflict case

# Commit: fix(hello): resolve pattern matching

# Commit: fix(test): handle path extraction case

# Commit: feat(docs): implement route matching

# Commit: refactor(middleware): improve memory usage

# Commit: chore(di): update go mod

# Commit: feat(di-container): implement module setup

# Commit: perf(route): optimize context propagation

# Commit: refactor(http): restructure code structure

# Commit: fix(ioc): resolve pattern matching

# Commit: docs(container): update documentation

# Commit: feat(module): add param extraction

# Commit: chore(controller): update test suite

# Commit: docs(module): update documentation

# Commit: refactor(controller): restructure type safety

# Commit: refactor(router): restructure test coverage

# Commit: test(server): add coverage for handler resolution

# Commit: test(provider): add coverage for module setup

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add route matching

# Commit: refactor(hello): improve concurrency handling

# Commit: feat(test): implement param extraction

# Commit: fix(docs): resolve body parsing

# Commit: chore(middleware): update ci configuration

# Commit: fix(di): resolve pattern matching

# Commit: docs(di-container): update documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): improve documentation

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve documentation

# Commit: docs(module): update documentation

# Commit: refactor(controller): improve type safety

# Commit: fix(router): handle path extraction case

# Commit: fix(controller): handle pattern matching case

# Commit: docs(router): update documentation

# Commit: perf(server): optimize controller routing

# Commit: chore(provider): update build script

# Commit: refactor(core): improve test coverage

# Commit: test(example): add coverage for controller routing

# Commit: chore(hello): update dependencies

# Commit: feat(test): add provider registration

# Commit: docs(docs): update documentation

# Commit: feat(middleware): add context propagation

# Commit: refactor(di): restructure code structure

# Commit: fix(di-container): resolve body parsing

# Commit: chore(route): update ci configuration

# Commit: fix(http): resolve path extraction

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add context propagation

# Commit: refactor(module): restructure concurrency handling

# Commit: refactor(controller): restructure performance

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve concurrency handling

# Commit: docs(router): update documentation

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update gitignore

# Commit: chore(core): update dependencies

# Commit: feat(example): add module setup

# Commit: perf(hello): optimize request injection

# Commit: fix(test): resolve body parsing

# Commit: chore(docs): update build script

# Commit: refactor(middleware): improve documentation

# Commit: fix(di): handle type inference case

# Commit: test(di-container): add coverage for response writing

# Commit: perf(route): optimize response writing

# Commit: refactor(http): improve error messages

# Commit: fix(ioc): handle header setting case

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle path extraction case

# Commit: feat(controller): implement module setup

# Commit: perf(router): optimize handler resolution

# Commit: test(server): add coverage for singleton scope

# Commit: feat(provider): add error handling

# Commit: feat(server): implement provider registration

# Commit: docs(provider): update documentation

# Commit: perf(core): optimize provider registration

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve pattern matching

# Commit: docs(test): update documentation

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize module setup

# Commit: perf(di): optimize route matching

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement middleware chain

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for provider registration

# Commit: docs(module): update documentation

# Commit: docs(controller): update documentation

# Commit: docs(router): update documentation

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update build script

# Commit: refactor(core): improve error messages

# Commit: fix(provider): resolve type inference

# Commit: test(core): add coverage for singleton scope

# Commit: feat(example): add provider registration

# Commit: docs(hello): update documentation

# Commit: feat(test): implement handler resolution

# Commit: test(docs): add coverage for request injection

# Commit: fix(middleware): resolve pattern matching

# Commit: docs(di): update documentation

# Commit: feat(di-container): add response writing

# Commit: perf(route): optimize context propagation

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve concurrency handling

# Commit: feat(module): implement controller routing

# Commit: chore(controller): update ci configuration

# Commit: fix(router): resolve header setting

# Commit: perf(server): optimize context propagation

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize controller routing

# Commit: chore(example): update gitignore

# Commit: chore(core): update license

# Commit: perf(example): optimize provider registration

# Commit: docs(hello): update documentation

# Commit: fix(test): handle path extraction case

# Commit: feat(docs): implement context propagation

# Commit: refactor(middleware): restructure documentation

# Commit: docs(di): update documentation

# Commit: perf(di-container): optimize error handling

# Commit: fix(route): handle header setting case

# Commit: perf(http): optimize param extraction

# Commit: feat(ioc): implement response writing

# Commit: feat(container): implement context propagation

# Commit: refactor(module): restructure type safety

# Commit: perf(controller): optimize request injection

# Commit: fix(router): resolve path extraction

# Commit: feat(server): implement request injection

# Commit: fix(provider): resolve type inference

# Commit: test(core): add coverage for error handling

# Commit: fix(example): handle body parsing case

# Commit: chore(hello): update readme

# Commit: refactor(example): improve documentation

# Commit: docs(hello): update documentation

# Commit: fix(test): handle type inference case

# Commit: test(docs): add coverage for middleware chain

# Commit: feat(middleware): implement module setup

# Commit: perf(di): optimize module setup

# Commit: perf(di-container): optimize module setup

# Commit: perf(route): optimize singleton scope

# Commit: feat(http): add controller routing

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update gitignore

# Commit: chore(module): update readme

# Commit: docs(controller): update documentation

# Commit: fix(router): handle body parsing case

# Commit: chore(server): update readme

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize controller routing

# Commit: chore(example): update license

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update go mod

# Commit: perf(hello): optimize route matching

# Commit: refactor(test): improve type safety

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement response writing

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle type inference case

# Commit: refactor(route): restructure type safety

# Commit: perf(http): optimize middleware chain

# Commit: feat(ioc): implement handler resolution

# Commit: test(container): add coverage for response writing

# Commit: fix(module): handle scope resolution case

# Commit: refactor(controller): restructure error messages

# Commit: refactor(router): improve error messages

# Commit: refactor(server): improve performance

# Commit: feat(provider): add error handling

# Commit: fix(core): handle pattern matching case

# Commit: fix(example): handle path extraction case

# Commit: feat(hello): implement response writing

# Commit: test(test): add coverage for singleton scope

# Commit: feat(docs): add param extraction

# Commit: feat(test): implement handler resolution

# Commit: test(docs): add coverage for controller routing

# Commit: chore(middleware): update test suite

# Commit: test(di): add coverage for handler resolution

# Commit: test(di-container): add coverage for controller routing

# Commit: chore(route): update dependencies

# Commit: feat(http): add route matching

# Commit: refactor(ioc): improve memory usage

# Commit: fix(container): handle routing conflict case

# Commit: fix(module): resolve path extraction

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize param extraction

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: feat(core): implement module setup

# Commit: perf(example): optimize provider registration

# Commit: docs(hello): update documentation

# Commit: fix(test): handle type inference case

# Commit: test(docs): add coverage for singleton scope

# Commit: feat(middleware): add param extraction

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add error handling

# Commit: fix(di): handle nil pointer case

# Commit: feat(di-container): add controller routing

# Commit: chore(route): update dependencies

# Commit: feat(http): add module setup

# Commit: perf(ioc): optimize singleton scope

# Commit: feat(container): add error handling

# Commit: fix(module): handle routing conflict case

# Commit: fix(controller): resolve nil pointer

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update license

# Commit: perf(provider): optimize param extraction

# Commit: feat(core): implement module setup

# Commit: perf(example): optimize request injection

# Commit: fix(hello): resolve type inference

# Commit: test(test): add coverage for provider registration

# Commit: docs(docs): update documentation

# Commit: perf(middleware): optimize request injection

# Commit: fix(di): resolve scope resolution

# Commit: refactor(middleware): improve memory usage

# Commit: chore(di): update license

# Commit: perf(di-container): optimize module setup

# Commit: perf(route): optimize request injection

# Commit: fix(http): resolve type inference

# Commit: test(ioc): add coverage for response writing

# Commit: feat(container): implement param extraction

# Commit: refactor(module): restructure type safety

# Commit: refactor(controller): restructure error messages

# Commit: refactor(router): improve performance

# Commit: feat(server): add context propagation

# Commit: refactor(provider): restructure performance

# Commit: refactor(core): restructure performance

# Commit: feat(example): add param extraction

# Commit: fix(hello): handle nil pointer case

# Commit: fix(test): handle type inference case

# Commit: test(docs): add coverage for middleware chain

# Commit: feat(middleware): implement controller routing

# Commit: chore(di): update readme

# Commit: refactor(di-container): restructure code structure

# Commit: fix(di): resolve path extraction

# Commit: feat(di-container): implement provider registration

# Commit: docs(route): update documentation

# Commit: fix(http): handle body parsing case

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: refactor(module): restructure type safety

# Commit: refactor(controller): restructure concurrency handling

# Commit: feat(router): implement response writing

# Commit: feat(server): add context propagation

# Commit: refactor(provider): restructure test coverage

# Commit: test(core): add coverage for request injection

# Commit: fix(example): resolve scope resolution

# Commit: refactor(hello): improve test coverage

# Commit: test(test): add coverage for middleware chain

# Commit: feat(docs): implement singleton scope

# Commit: feat(middleware): add middleware chain

# Commit: feat(di): implement module setup

# Commit: perf(di-container): optimize error handling

# Commit: fix(route): handle pattern matching case

# Commit: perf(di-container): optimize middleware chain

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: fix(controller): resolve path extraction

# Commit: feat(router): implement route matching

# Commit: refactor(server): improve performance

# Commit: feat(provider): add context propagation

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve error messages

# Commit: refactor(hello): improve documentation

# Commit: docs(test): update documentation

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add singleton scope

# Commit: feat(di): add response writing

# Commit: feat(di-container): implement error handling

# Commit: fix(route): handle body parsing case

# Commit: chore(http): update dependencies

# Commit: test(route): add coverage for param extraction

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add controller routing

# Commit: chore(container): update gitignore

# Commit: chore(module): update go mod

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure performance

# Commit: feat(server): add route matching

# Commit: refactor(provider): improve test coverage

# Commit: test(core): add coverage for context propagation

# Commit: refactor(example): restructure performance

# Commit: feat(hello): add param extraction

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure concurrency handling

# Commit: feat(middleware): implement error handling

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize middleware chain

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: test(ioc): add coverage for route matching

# Commit: chore(http): update build script

# Commit: refactor(ioc): improve performance

# Commit: feat(container): add context propagation

# Commit: refactor(module): restructure performance

# Commit: feat(controller): add middleware chain

# Commit: feat(router): implement context propagation

# Commit: refactor(server): restructure documentation

# Commit: refactor(provider): restructure code structure

# Commit: fix(core): resolve type inference

# Commit: test(example): add coverage for handler resolution

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle header setting case

# Commit: refactor(docs): restructure concurrency handling

# Commit: feat(middleware): implement singleton scope

# Commit: feat(di): add response writing

# Commit: fix(di-container): resolve body parsing

# Commit: chore(route): update test suite

# Commit: test(http): add coverage for middleware chain

# Commit: feat(ioc): implement context propagation

# Commit: refactor(container): restructure code structure

# Commit: refactor(ioc): restructure error messages

# Commit: fix(container): handle body parsing case

# Commit: chore(module): update dependencies

# Commit: feat(controller): add response writing

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure error messages

# Commit: refactor(provider): improve documentation

# Commit: fix(core): handle routing conflict case

# Commit: fix(example): resolve header setting

# Commit: perf(hello): optimize route matching

# Commit: refactor(test): improve performance

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for error handling

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize controller routing

# Commit: chore(route): update dependencies

# Commit: feat(http): add error handling

# Commit: fix(ioc): handle type inference case

# Commit: test(container): add coverage for controller routing

# Commit: chore(module): update build script

# Commit: feat(container): implement route matching

# Commit: refactor(module): improve concurrency handling

# Commit: feat(controller): implement controller routing

# Commit: chore(router): update dependencies

# Commit: feat(server): add handler resolution

# Commit: test(provider): add coverage for provider registration

# Commit: docs(core): update documentation

# Commit: fix(example): resolve nil pointer

# Commit: feat(hello): add error handling

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add middleware chain

# Commit: feat(middleware): implement middleware chain

# Commit: feat(di): implement context propagation

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: test(http): add coverage for middleware chain

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve concurrency handling

# Commit: feat(module): implement singleton scope

# Commit: feat(controller): add middleware chain

# Commit: perf(module): optimize request injection

# Commit: fix(controller): resolve routing conflict

# Commit: fix(router): resolve type inference

# Commit: test(server): add coverage for error handling

# Commit: fix(provider): handle pattern matching case

# Commit: docs(core): update documentation

# Commit: fix(example): resolve nil pointer

# Commit: feat(hello): add param extraction

# Commit: fix(test): handle nil pointer case

# Commit: fix(docs): handle path extraction case

# Commit: fix(middleware): handle routing conflict case

# Commit: fix(di): resolve path extraction

# Commit: feat(di-container): implement request injection

# Commit: fix(route): resolve scope resolution

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update dependencies

# Commit: feat(container): add module setup

# Commit: perf(module): optimize route matching

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for controller routing

# Commit: test(controller): add coverage for context propagation

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve memory usage

# Commit: chore(provider): update ci configuration

# Commit: fix(core): resolve header setting

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure type safety

# Commit: perf(test): optimize singleton scope

# Commit: feat(docs): add param extraction

# Commit: perf(middleware): optimize request injection

# Commit: fix(di): resolve pattern matching

# Commit: docs(di-container): update documentation

# Commit: perf(route): optimize handler resolution

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add controller routing

# Commit: chore(container): update ci configuration

# Commit: fix(module): resolve pattern matching

# Commit: fix(controller): handle routing conflict case

# Commit: fix(router): resolve routing conflict

# Commit: fix(server): resolve type inference

# Commit: fix(router): resolve type inference

# Commit: test(server): add coverage for response writing

# Commit: chore(provider): update ci configuration

# Commit: fix(core): resolve scope resolution

# Commit: refactor(example): improve memory usage

# Commit: chore(hello): update go mod

# Commit: feat(test): implement route matching

# Commit: refactor(docs): improve documentation

# Commit: docs(middleware): update documentation

# Commit: perf(di): optimize handler resolution

# Commit: test(di-container): add coverage for handler resolution

# Commit: test(route): add coverage for middleware chain

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure memory usage

# Commit: chore(container): update gitignore

# Commit: chore(module): update gitignore

# Commit: chore(controller): update build script

# Commit: refactor(router): improve memory usage

# Commit: chore(server): update dependencies

# Commit: feat(provider): add handler resolution

# Commit: feat(server): implement param extraction

# Commit: perf(provider): optimize controller routing

# Commit: chore(core): update ci configuration

# Commit: refactor(example): restructure documentation

# Commit: docs(hello): update documentation

# Commit: fix(test): resolve scope resolution

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize provider registration

# Commit: docs(di): update documentation

# Commit: fix(di-container): handle nil pointer case

# Commit: refactor(route): restructure performance

# Commit: refactor(http): restructure type safety

# Commit: perf(ioc): optimize middleware chain

# Commit: feat(container): implement handler resolution

# Commit: test(module): add coverage for middleware chain

# Commit: feat(controller): implement param extraction

# Commit: fix(router): handle nil pointer case

# Commit: feat(server): add handler resolution

# Commit: test(provider): add coverage for response writing

# Commit: test(core): add coverage for controller routing

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve error messages

# Commit: refactor(example): restructure memory usage

# Commit: chore(hello): update gitignore

# Commit: chore(test): update readme

# Commit: docs(docs): update documentation

# Commit: feat(middleware): implement request injection

# Commit: fix(di): resolve type inference

# Commit: test(di-container): add coverage for error handling

# Commit: fix(route): handle type inference case

# Commit: test(http): add coverage for param extraction

# Commit: test(ioc): add coverage for middleware chain

# Commit: feat(container): implement route matching

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: fix(router): handle type inference case

# Commit: test(server): add coverage for request injection

# Commit: fix(provider): resolve nil pointer

# Commit: feat(core): add context propagation

# Commit: refactor(example): restructure concurrency handling

# Commit: feat(core): implement param extraction

# Commit: feat(example): implement route matching

# Commit: refactor(hello): improve documentation

# Commit: docs(test): update documentation

# Commit: feat(docs): implement module setup

# Commit: perf(middleware): optimize singleton scope

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle path extraction case

# Commit: fix(route): handle path extraction case

# Commit: fix(http): handle scope resolution case

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement context propagation

# Commit: refactor(module): restructure error messages

# Commit: refactor(controller): improve type safety

# Commit: perf(router): optimize singleton scope

# Commit: feat(server): add module setup

# Commit: perf(provider): optimize error handling

# Commit: fix(core): handle body parsing case

# Commit: chore(example): update gitignore

# Commit: chore(hello): update readme

# Commit: fix(example): handle type inference case

# Commit: test(hello): add coverage for context propagation

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve routing conflict

# Commit: refactor(middleware): restructure test coverage

# Commit: test(di): add coverage for middleware chain

# Commit: feat(di-container): implement controller routing

# Commit: chore(route): update gitignore

# Commit: chore(http): update dependencies

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle header setting case

# Commit: perf(module): optimize param extraction

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure concurrency handling

# Commit: feat(provider): implement route matching

# Commit: refactor(core): improve type safety

# Commit: perf(example): optimize route matching

# Commit: refactor(hello): improve error messages

# Commit: fix(test): handle scope resolution case

# Commit: feat(hello): implement module setup

# Commit: perf(test): optimize singleton scope

# Commit: feat(docs): add error handling

# Commit: fix(middleware): handle path extraction case

# Commit: feat(di): implement provider registration

# Commit: docs(di-container): update documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update ci configuration

# Commit: fix(container): resolve nil pointer

# Commit: fix(module): handle pattern matching case

# Commit: docs(controller): update documentation

# Commit: chore(router): update ci configuration

# Commit: refactor(server): restructure performance

# Commit: feat(provider): add response writing

# Commit: fix(core): resolve routing conflict

# Commit: fix(example): resolve nil pointer

# Commit: feat(hello): add error handling

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add middleware chain

# Commit: feat(test): add module setup

# Commit: perf(docs): optimize module setup

# Commit: perf(middleware): optimize handler resolution

# Commit: test(di): add coverage for module setup

# Commit: perf(di-container): optimize controller routing

# Commit: chore(route): update dependencies

# Commit: refactor(http): restructure memory usage

# Commit: chore(ioc): update build script

# Commit: refactor(container): improve code structure

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update license

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure memory usage

# Commit: chore(provider): update ci configuration

# Commit: fix(core): resolve body parsing

# Commit: chore(example): update test suite

# Commit: test(hello): add coverage for route matching

# Commit: refactor(test): improve test coverage

# Commit: refactor(docs): restructure memory usage

# Commit: chore(middleware): update build script

# Commit: perf(docs): optimize module setup

# Commit: perf(middleware): optimize module setup

# Commit: perf(di): optimize handler resolution

# Commit: test(di-container): add coverage for param extraction

# Commit: perf(route): optimize response writing

# Commit: fix(http): resolve nil pointer

# Commit: fix(ioc): handle header setting case

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update readme

# Commit: docs(controller): update documentation

# Commit: fix(router): resolve pattern matching

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for module setup

# Commit: perf(core): optimize module setup

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize middleware chain

# Commit: feat(test): implement controller routing

# Commit: chore(docs): update build script

# Commit: refactor(middleware): improve documentation

# Commit: docs(di): update documentation

# Commit: fix(middleware): resolve body parsing

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize singleton scope

# Commit: feat(route): add error handling

# Commit: fix(http): handle header setting case

# Commit: perf(ioc): optimize request injection

# Commit: fix(container): resolve header setting

# Commit: refactor(module): restructure error messages

# Commit: refactor(controller): improve type safety

# Commit: perf(router): optimize middleware chain

# Commit: feat(server): implement singleton scope

# Commit: feat(provider): add route matching

# Commit: refactor(core): improve documentation

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize response writing

# Commit: chore(docs): update ci configuration

# Commit: fix(middleware): resolve type inference

# Commit: test(di): add coverage for error handling

# Commit: fix(di-container): handle body parsing case

# Commit: perf(di): optimize request injection

# Commit: fix(di-container): resolve header setting

# Commit: perf(route): optimize request injection

# Commit: fix(http): resolve scope resolution

# Commit: fix(ioc): handle path extraction case

# Commit: feat(container): implement controller routing

# Commit: chore(module): update go mod

# Commit: fix(controller): handle path extraction case

# Commit: feat(router): implement provider registration

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for handler resolution

# Commit: test(core): add coverage for handler resolution

# Commit: test(example): add coverage for request injection

# Commit: fix(hello): resolve scope resolution

# Commit: refactor(test): improve type safety

# Commit: refactor(docs): restructure concurrency handling

# Commit: feat(middleware): implement middleware chain

# Commit: feat(di): implement middleware chain

# Commit: feat(di-container): implement singleton scope

# Commit: feat(route): add provider registration

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for request injection

# Commit: fix(http): resolve header setting

# Commit: perf(ioc): optimize response writing

# Commit: feat(container): add handler resolution

# Commit: test(module): add coverage for route matching

# Commit: refactor(controller): improve concurrency handling

# Commit: feat(router): implement middleware chain

# Commit: feat(server): implement handler resolution

# Commit: test(provider): add coverage for error handling

# Commit: fix(core): handle header setting case

# Commit: perf(example): optimize param extraction

# Commit: feat(hello): add route matching

# Commit: refactor(test): improve type safety

# Commit: perf(docs): optimize handler resolution

# Commit: test(middleware): add coverage for module setup

# Commit: perf(di): optimize controller routing

# Commit: chore(di-container): update ci configuration

# Commit: fix(route): resolve nil pointer

# Commit: feat(http): add response writing

# Commit: fix(route): handle pattern matching case

# Commit: docs(http): update documentation

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle type inference case

# Commit: test(module): add coverage for route matching

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for param extraction

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure documentation

# Commit: refactor(core): restructure memory usage

# Commit: chore(example): update gitignore

# Commit: chore(hello): update license

# Commit: perf(test): optimize controller routing

# Commit: chore(docs): update test suite

# Commit: test(middleware): add coverage for route matching

# Commit: refactor(di): improve memory usage

# Commit: chore(di-container): update ci configuration

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for route matching

# Commit: refactor(ioc): improve documentation

# Commit: perf(http): optimize route matching

# Commit: refactor(ioc): improve test coverage

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add controller routing

# Commit: chore(controller): update build script

# Commit: refactor(router): improve concurrency handling

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement handler resolution

# Commit: test(core): add coverage for param extraction

# Commit: feat(example): add route matching

# Commit: refactor(hello): improve type safety

# Commit: perf(test): optimize singleton scope

# Commit: feat(docs): add singleton scope

# Commit: feat(middleware): add context propagation

# Commit: refactor(di): restructure documentation

# Commit: docs(di-container): update documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): improve error messages

# Commit: refactor(ioc): improve memory usage

# Commit: chore(container): update build script

# Commit: docs(ioc): update documentation

# Commit: refactor(container): restructure concurrency handling

# Commit: feat(module): implement context propagation

# Commit: refactor(controller): restructure documentation

# Commit: docs(router): update documentation

# Commit: feat(server): implement param extraction

# Commit: refactor(provider): improve documentation

# Commit: fix(core): handle path extraction case

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize singleton scope

# Commit: feat(middleware): add handler resolution

# Commit: test(di): add coverage for response writing

# Commit: fix(di-container): resolve body parsing

# Commit: chore(route): update readme

# Commit: docs(http): update documentation

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve documentation

# Commit: docs(module): update documentation

# Commit: test(container): add coverage for response writing

# Commit: chore(module): update build script

# Commit: refactor(controller): improve documentation

# Commit: docs(router): update documentation

# Commit: refactor(server): restructure code structure

# Commit: fix(provider): resolve pattern matching

# Commit: docs(core): update documentation

# Commit: refactor(example): improve concurrency handling

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve body parsing

# Commit: fix(docs): handle routing conflict case

# Commit: fix(middleware): resolve body parsing

# Commit: chore(di): update gitignore

# Commit: fix(di-container): handle pattern matching case

# Commit: refactor(route): restructure type safety

# Commit: fix(http): handle path extraction case

# Commit: feat(ioc): implement param extraction

# Commit: perf(container): optimize response writing

# Commit: fix(module): resolve routing conflict

# Commit: fix(controller): resolve body parsing

# Commit: perf(module): optimize provider registration

# Commit: docs(controller): update documentation

# Commit: feat(router): implement route matching

# Commit: refactor(server): improve test coverage

# Commit: test(provider): add coverage for provider registration

# Commit: docs(core): update documentation

# Commit: refactor(example): restructure error messages

# Commit: refactor(hello): improve type safety

# Commit: perf(test): optimize middleware chain

# Commit: feat(docs): implement param extraction

# Commit: perf(middleware): optimize singleton scope

# Commit: feat(di): add handler resolution

# Commit: test(di-container): add coverage for param extraction

# Commit: refactor(route): restructure documentation

# Commit: refactor(http): restructure performance

# Commit: feat(ioc): add param extraction

# Commit: test(container): add coverage for middleware chain

# Commit: feat(module): implement param extraction

# Commit: refactor(controller): improve type safety

# Commit: perf(router): optimize param extraction

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve code structure

# Commit: fix(server): resolve routing conflict

# Commit: fix(provider): resolve body parsing

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve concurrency handling

# Commit: feat(hello): implement param extraction

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize error handling

# Commit: fix(middleware): handle header setting case

# Commit: perf(di): optimize module setup

# Commit: perf(di-container): optimize singleton scope

# Commit: feat(route): add response writing

# Commit: fix(http): resolve header setting

# Commit: perf(ioc): optimize error handling

# Commit: fix(container): handle body parsing case

# Commit: refactor(module): restructure test coverage

# Commit: test(controller): add coverage for error handling

# Commit: fix(router): handle path extraction case

# Commit: feat(server): implement route matching

# Commit: test(router): add coverage for handler resolution

# Commit: test(server): add coverage for context propagation

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize response writing

# Commit: refactor(example): restructure error messages

# Commit: fix(hello): handle path extraction case

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure concurrency handling

# Commit: feat(middleware): implement handler resolution

# Commit: test(di): add coverage for route matching

# Commit: refactor(di-container): improve documentation

# Commit: docs(route): update documentation

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add request injection

# Commit: fix(container): resolve body parsing

# Commit: chore(module): update license

# Commit: perf(controller): optimize provider registration

# Commit: docs(router): update documentation

# Commit: feat(server): implement param extraction

# Commit: perf(provider): optimize controller routing

# Commit: perf(server): optimize param extraction

# Commit: chore(provider): update ci configuration

# Commit: fix(core): resolve type inference

# Commit: test(example): add coverage for error handling

# Commit: fix(hello): handle path extraction case

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add route matching

# Commit: refactor(middleware): improve type safety

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add singleton scope

# Commit: feat(route): add response writing

# Commit: docs(http): update documentation

# Commit: perf(ioc): optimize singleton scope

# Commit: feat(container): add singleton scope

# Commit: feat(module): add param extraction

# Commit: refactor(controller): restructure test coverage

# Commit: test(router): add coverage for request injection

# Commit: fix(server): resolve pattern matching

# Commit: refactor(provider): restructure documentation

# Commit: docs(core): update documentation

# Commit: feat(provider): implement context propagation

# Commit: refactor(core): restructure type safety

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize provider registration

# Commit: docs(test): update documentation

# Commit: perf(docs): optimize handler resolution

# Commit: test(middleware): add coverage for response writing

# Commit: chore(di): update dependencies

# Commit: feat(di-container): add error handling

# Commit: fix(route): handle header setting case

# Commit: perf(http): optimize middleware chain

# Commit: feat(ioc): implement provider registration

# Commit: docs(container): update documentation

# Commit: perf(module): optimize route matching

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve error messages

# Commit: fix(server): handle pattern matching case

# Commit: docs(provider): update documentation

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: chore(core): update build script

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for handler resolution

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize handler resolution

# Commit: test(middleware): add coverage for singleton scope

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle body parsing case

# Commit: chore(route): update ci configuration

# Commit: fix(http): resolve nil pointer

# Commit: feat(ioc): add middleware chain

# Commit: feat(container): implement middleware chain

# Commit: feat(module): implement error handling

# Commit: fix(controller): handle routing conflict case

# Commit: refactor(router): restructure type safety

# Commit: perf(server): optimize param extraction

# Commit: perf(provider): optimize error handling

# Commit: fix(core): handle pattern matching case

# Commit: docs(example): update documentation

# Commit: feat(hello): implement controller routing

# Commit: chore(example): update dependencies

# Commit: feat(hello): add module setup

# Commit: perf(test): optimize route matching

# Commit: refactor(docs): improve performance

# Commit: feat(middleware): add error handling

# Commit: fix(di): handle pattern matching case

# Commit: refactor(di-container): restructure performance

# Commit: feat(route): add request injection

# Commit: fix(http): resolve header setting

# Commit: perf(ioc): optimize controller routing

# Commit: chore(container): update ci configuration

# Commit: fix(module): resolve pattern matching

# Commit: docs(controller): update documentation

# Commit: refactor(router): improve test coverage

# Commit: test(server): add coverage for request injection

# Commit: fix(provider): resolve pattern matching

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: perf(hello): optimize module setup

# Commit: perf(test): optimize response writing

# Commit: refactor(hello): restructure type safety

# Commit: perf(test): optimize request injection

# Commit: fix(docs): resolve type inference

# Commit: test(middleware): add coverage for provider registration

# Commit: docs(di): update documentation

# Commit: feat(di-container): add request injection

# Commit: fix(route): resolve nil pointer

# Commit: feat(http): add singleton scope

# Commit: feat(ioc): add param extraction

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle pattern matching case

# Commit: docs(controller): update documentation

# Commit: feat(router): implement middleware chain

# Commit: feat(server): implement module setup

# Commit: perf(provider): optimize handler resolution

# Commit: test(core): add coverage for module setup

# Commit: perf(example): optimize module setup

# Commit: perf(hello): optimize param extraction

# Commit: docs(test): update documentation

# Commit: fix(docs): resolve pattern matching

# Commit: chore(test): update test suite

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle type inference case

# Commit: test(di): add coverage for context propagation

# Commit: refactor(di-container): restructure memory usage

# Commit: chore(route): update go mod

# Commit: fix(http): handle type inference case

# Commit: test(ioc): add coverage for controller routing

# Commit: chore(container): update dependencies

# Commit: feat(module): add request injection

# Commit: fix(controller): resolve type inference

# Commit: test(router): add coverage for singleton scope

# Commit: feat(server): add route matching

# Commit: refactor(provider): improve test coverage

# Commit: test(core): add coverage for module setup

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for controller routing

# Commit: chore(test): update build script

# Commit: refactor(docs): improve test coverage

# Commit: test(middleware): add coverage for module setup

# Commit: test(docs): add coverage for controller routing

# Commit: chore(middleware): update build script

# Commit: refactor(di): improve type safety

# Commit: perf(di-container): optimize module setup

# Commit: perf(route): optimize controller routing

# Commit: chore(http): update go mod

# Commit: feat(ioc): implement middleware chain

# Commit: feat(container): implement singleton scope

# Commit: feat(module): add route matching

# Commit: refactor(controller): improve type safety

# Commit: perf(router): optimize response writing

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: feat(core): add controller routing

# Commit: chore(example): update build script

# Commit: fix(hello): handle scope resolution case

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize param extraction

# Commit: perf(di): optimize provider registration

# Commit: feat(middleware): add singleton scope

# Commit: feat(di): add param extraction

# Commit: refactor(di-container): improve performance

# Commit: feat(route): add request injection

# Commit: fix(http): resolve pattern matching

# Commit: refactor(ioc): restructure type safety

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle pattern matching case

# Commit: docs(controller): update documentation

# Commit: feat(router): add controller routing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement error handling

# Commit: fix(core): handle nil pointer case

# Commit: refactor(example): restructure test coverage

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve type inference

# Commit: refactor(middleware): restructure concurrency handling

# Commit: feat(di): implement middleware chain

# Commit: feat(di-container): implement middleware chain

# Commit: refactor(di): restructure error messages

# Commit: refactor(di-container): improve type safety

# Commit: fix(route): handle body parsing case

# Commit: chore(http): update build script

# Commit: fix(ioc): handle type inference case

# Commit: test(container): add coverage for middleware chain

# Commit: feat(module): implement middleware chain

# Commit: feat(controller): implement singleton scope

# Commit: feat(router): add context propagation

# Commit: refactor(server): restructure type safety

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve path extraction

# Commit: feat(example): implement route matching

# Commit: refactor(hello): improve documentation

# Commit: fix(test): handle header setting case

# Commit: perf(docs): optimize route matching

# Commit: refactor(middleware): improve error messages

# Commit: refactor(di): improve documentation

# Commit: docs(di-container): update documentation

# Commit: fix(route): resolve scope resolution

# Commit: feat(di-container): add route matching

# Commit: refactor(route): improve type safety

# Commit: perf(http): optimize module setup

# Commit: perf(ioc): optimize response writing

# Commit: test(container): add coverage for route matching

# Commit: refactor(module): improve memory usage

# Commit: fix(controller): handle path extraction case

# Commit: feat(router): implement singleton scope

# Commit: feat(server): add singleton scope

# Commit: feat(provider): add route matching

# Commit: refactor(core): improve documentation

# Commit: docs(example): update documentation

# Commit: feat(hello): implement module setup

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize handler resolution

# Commit: test(middleware): add coverage for response writing

# Commit: fix(di): handle routing conflict case

# Commit: fix(di-container): resolve body parsing

# Commit: refactor(route): restructure performance

# Commit: feat(http): add route matching

# Commit: refactor(route): restructure code structure

# Commit: fix(http): resolve path extraction

# Commit: feat(ioc): implement provider registration

# Commit: docs(container): update documentation

# Commit: docs(module): update documentation

# Commit: fix(controller): handle nil pointer case

# Commit: fix(router): handle body parsing case

# Commit: chore(server): update gitignore

# Commit: refactor(provider): restructure performance

# Commit: feat(core): add handler resolution

# Commit: test(example): add coverage for middleware chain

# Commit: feat(hello): implement handler resolution

# Commit: test(test): add coverage for param extraction

# Commit: fix(docs): handle type inference case

# Commit: fix(middleware): handle nil pointer case

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle type inference case

# Commit: test(route): add coverage for request injection

# Commit: fix(http): resolve type inference

# Commit: fix(ioc): handle path extraction case

# Commit: refactor(http): restructure documentation

# Commit: docs(ioc): update documentation

# Commit: test(container): add coverage for route matching

# Commit: refactor(module): improve type safety

# Commit: fix(controller): handle type inference case

# Commit: test(router): add coverage for handler resolution

# Commit: test(server): add coverage for error handling

# Commit: fix(provider): handle routing conflict case

# Commit: fix(core): resolve pattern matching

# Commit: docs(example): update documentation

# Commit: refactor(hello): improve documentation

# Commit: docs(test): update documentation

# Commit: feat(docs): implement response writing

# Commit: chore(middleware): update dependencies

# Commit: feat(di): add handler resolution

# Commit: test(di-container): add coverage for provider registration

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure performance

# Commit: feat(ioc): add singleton scope

# Commit: feat(container): add singleton scope

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: refactor(module): restructure concurrency handling

# Commit: feat(controller): implement provider registration

# Commit: docs(router): update documentation

# Commit: chore(server): update ci configuration

# Commit: fix(provider): resolve nil pointer

# Commit: feat(core): add provider registration

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure documentation

# Commit: docs(test): update documentation

# Commit: docs(docs): update documentation

# Commit: perf(middleware): optimize route matching

# Commit: refactor(di): improve error messages

# Commit: refactor(di-container): improve performance

# Commit: feat(route): add singleton scope

# Commit: feat(http): add module setup

# Commit: perf(ioc): optimize route matching

# Commit: refactor(container): improve concurrency handling

# Commit: feat(module): implement response writing

# Commit: feat(container): add controller routing

# Commit: chore(module): update build script

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve memory usage

# Commit: chore(server): update ci configuration

# Commit: fix(provider): resolve header setting

# Commit: perf(core): optimize middleware chain

# Commit: feat(example): implement param extraction

# Commit: chore(hello): update go mod

# Commit: feat(test): implement module setup

# Commit: perf(docs): optimize module setup

# Commit: perf(middleware): optimize param extraction

# Commit: feat(di): add context propagation

# Commit: refactor(di-container): restructure memory usage

# Commit: chore(route): update build script

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for controller routing

# Commit: chore(controller): update license

# Commit: docs(module): update documentation

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle routing conflict case

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): resolve scope resolution

# Commit: fix(example): handle scope resolution case

# Commit: fix(hello): handle path extraction case

# Commit: feat(test): implement response writing

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize controller routing

# Commit: chore(di): update build script

# Commit: refactor(di-container): improve error messages

# Commit: refactor(route): restructure code structure

# Commit: fix(http): resolve scope resolution

# Commit: refactor(ioc): improve documentation

# Commit: docs(container): update documentation

# Commit: feat(module): add controller routing

# Commit: chore(controller): update dependencies

# Commit: feat(router): add response writing

# Commit: fix(controller): handle body parsing case

# Commit: chore(router): update readme

# Commit: docs(server): update documentation

# Commit: chore(provider): update license

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add controller routing

# Commit: chore(hello): update go mod

# Commit: feat(test): implement response writing

# Commit: fix(docs): resolve type inference

# Commit: test(middleware): add coverage for handler resolution

# Commit: test(di): add coverage for response writing

# Commit: perf(di-container): optimize singleton scope

# Commit: feat(route): add error handling

# Commit: fix(http): handle nil pointer case

# Commit: feat(ioc): add param extraction

# Commit: chore(container): update build script

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update license

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure type safety

# Commit: fix(router): handle header setting case

# Commit: perf(server): optimize provider registration

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve performance

# Commit: feat(example): add error handling

# Commit: fix(hello): handle nil pointer case

# Commit: feat(test): add handler resolution

# Commit: test(docs): add coverage for module setup

# Commit: perf(middleware): optimize controller routing

# Commit: chore(di): update dependencies

# Commit: refactor(di-container): restructure type safety

# Commit: perf(route): optimize param extraction

# Commit: refactor(http): improve type safety

# Commit: refactor(ioc): restructure memory usage

# Commit: chore(container): update readme

# Commit: docs(module): update documentation

# Commit: refactor(controller): improve documentation

# Commit: docs(router): update documentation

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: feat(server): implement singleton scope

# Commit: feat(provider): add middleware chain

# Commit: feat(core): implement middleware chain

# Commit: feat(example): implement singleton scope

# Commit: feat(hello): add context propagation

# Commit: refactor(test): restructure test coverage

# Commit: refactor(docs): restructure error messages

# Commit: refactor(middleware): improve performance

# Commit: feat(di): add route matching

# Commit: refactor(di-container): improve error messages

# Commit: refactor(route): improve documentation

# Commit: docs(http): update documentation

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement handler resolution

# Commit: test(module): add coverage for request injection

# Commit: fix(controller): resolve pattern matching

# Commit: docs(router): update documentation

# Commit: fix(server): handle path extraction case

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update dependencies

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve path extraction

# Commit: feat(example): implement error handling

# Commit: fix(hello): handle scope resolution case

# Commit: refactor(test): improve code structure

# Commit: fix(docs): resolve scope resolution

# Commit: refactor(middleware): improve documentation

# Commit: docs(di): update documentation

# Commit: docs(di-container): update documentation

# Commit: feat(route): add provider registration

# Commit: docs(http): update documentation

# Commit: refactor(ioc): restructure test coverage

# Commit: refactor(container): restructure code structure

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve memory usage

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for response writing

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update license

# Commit: perf(example): optimize context propagation

# Commit: refactor(core): restructure documentation

# Commit: fix(example): handle header setting case

# Commit: perf(hello): optimize handler resolution

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize param extraction

# Commit: docs(di): update documentation

# Commit: feat(di-container): implement request injection

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update license

# Commit: perf(ioc): optimize provider registration

# Commit: docs(container): update documentation

# Commit: docs(module): update documentation

# Commit: feat(controller): implement controller routing

# Commit: chore(router): update build script

# Commit: refactor(server): improve error messages

# Commit: refactor(provider): improve code structure

# Commit: fix(core): resolve routing conflict

# Commit: refactor(example): restructure error messages

# Commit: refactor(hello): improve type safety

# Commit: feat(example): implement request injection

# Commit: fix(hello): resolve header setting

# Commit: fix(test): handle scope resolution case

# Commit: refactor(docs): improve performance

# Commit: refactor(middleware): restructure concurrency handling

# Commit: feat(di): implement error handling

# Commit: fix(di-container): handle routing conflict case

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update gitignore

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add module setup

# Commit: perf(controller): optimize response writing

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for response writing

# Commit: fix(core): resolve path extraction

# Commit: feat(example): implement route matching

# Commit: refactor(hello): improve concurrency handling

# Commit: feat(test): implement param extraction

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize handler resolution

# Commit: test(docs): add coverage for provider registration

# Commit: docs(middleware): update documentation

# Commit: refactor(di): improve performance

# Commit: feat(di-container): add request injection

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update gitignore

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement param extraction

# Commit: feat(module): implement error handling

# Commit: fix(controller): handle pattern matching case

# Commit: docs(router): update documentation

# Commit: chore(server): update license

# Commit: perf(provider): optimize module setup

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add controller routing

# Commit: chore(hello): update go mod

# Commit: feat(test): implement request injection

# Commit: fix(docs): resolve nil pointer

# Commit: feat(test): add request injection

# Commit: fix(docs): resolve nil pointer

# Commit: feat(middleware): add handler resolution

# Commit: test(di): add coverage for provider registration

# Commit: docs(di-container): update documentation

# Commit: perf(route): optimize context propagation

# Commit: refactor(http): restructure concurrency handling

# Commit: fix(ioc): handle body parsing case

# Commit: chore(container): update go mod

# Commit: feat(module): implement handler resolution

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for route matching

# Commit: refactor(server): improve documentation

# Commit: docs(provider): update documentation

# Commit: test(core): add coverage for provider registration

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure code structure

# Commit: fix(test): resolve nil pointer

# Commit: feat(docs): add response writing

# Commit: fix(middleware): resolve scope resolution

# Commit: feat(docs): implement middleware chain

# Commit: feat(middleware): implement error handling

# Commit: fix(di): handle type inference case

# Commit: test(di-container): add coverage for controller routing

# Commit: chore(route): update build script

# Commit: refactor(http): improve error messages

# Commit: refactor(ioc): improve code structure

# Commit: fix(container): resolve nil pointer

# Commit: feat(module): add handler resolution

# Commit: test(controller): add coverage for response writing

# Commit: feat(router): add controller routing

# Commit: chore(server): update readme

# Commit: docs(provider): update documentation

# Commit: chore(core): update test suite

# Commit: test(example): add coverage for request injection

# Commit: fix(hello): resolve type inference

# Commit: test(test): add coverage for response writing

# Commit: refactor(docs): restructure documentation

# Commit: docs(middleware): update documentation

# Commit: test(di): add coverage for error handling

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update dependencies

# Commit: refactor(di-container): restructure error messages

# Commit: refactor(route): improve error messages

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add route matching

# Commit: refactor(module): improve type safety

# Commit: perf(controller): optimize context propagation

# Commit: refactor(router): restructure code structure

# Commit: fix(server): handle body parsing case

# Commit: fix(provider): handle type inference case

# Commit: fix(core): handle scope resolution case

# Commit: refactor(example): improve memory usage

# Commit: refactor(hello): restructure type safety

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve pattern matching

# Commit: docs(middleware): update documentation

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize request injection

# Commit: fix(di): handle path extraction case

# Commit: feat(di-container): implement middleware chain

# Commit: feat(route): implement context propagation

# Commit: refactor(http): restructure test coverage

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add singleton scope

# Commit: feat(module): add request injection

# Commit: fix(controller): resolve routing conflict

# Commit: fix(router): resolve path extraction

# Commit: feat(server): implement route matching

# Commit: refactor(provider): improve test coverage

# Commit: test(core): add coverage for context propagation

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): resolve nil pointer

# Commit: fix(test): handle body parsing case

# Commit: chore(docs): update readme

# Commit: docs(middleware): update documentation

# Commit: feat(di): add handler resolution

# Commit: test(di-container): add coverage for param extraction

# Commit: fix(route): handle scope resolution case

# Commit: docs(di-container): update documentation

# Commit: feat(route): add provider registration

# Commit: docs(http): update documentation

# Commit: fix(ioc): resolve scope resolution

# Commit: refactor(container): restructure type safety

# Commit: perf(module): optimize handler resolution

# Commit: test(controller): add coverage for module setup

# Commit: perf(router): optimize handler resolution

# Commit: test(server): add coverage for singleton scope

# Commit: feat(provider): add middleware chain

# Commit: feat(core): implement singleton scope

# Commit: feat(example): add context propagation

# Commit: refactor(hello): restructure performance

# Commit: feat(test): add request injection

# Commit: fix(docs): resolve pattern matching

# Commit: docs(middleware): update documentation

# Commit: fix(di): resolve type inference

# Commit: test(di-container): add coverage for route matching

# Commit: refactor(route): improve concurrency handling

# Commit: feat(http): implement controller routing

# Commit: fix(route): handle path extraction case

# Commit: feat(http): implement error handling

# Commit: fix(ioc): handle header setting case

# Commit: refactor(container): restructure concurrency handling

# Commit: feat(module): implement handler resolution

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for singleton scope

# Commit: feat(server): add module setup

# Commit: perf(provider): optimize context propagation

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve error messages

# Commit: refactor(hello): improve performance

# Commit: feat(test): add route matching

# Commit: refactor(docs): improve memory usage

# Commit: chore(middleware): update build script

# Commit: refactor(di): improve documentation

# Commit: docs(di-container): update documentation

# Commit: fix(route): resolve routing conflict

# Commit: fix(http): resolve routing conflict

# Commit: fix(ioc): resolve header setting

# Commit: perf(http): optimize context propagation

# Commit: refactor(ioc): restructure concurrency handling

# Commit: feat(container): implement handler resolution

# Commit: test(module): add coverage for context propagation

# Commit: refactor(controller): restructure concurrency handling

# Commit: feat(router): implement singleton scope

# Commit: feat(server): add param extraction

# Commit: perf(provider): optimize context propagation

# Commit: refactor(core): restructure test coverage

# Commit: test(example): add coverage for middleware chain

# Commit: feat(hello): implement route matching

# Commit: refactor(test): improve memory usage

# Commit: chore(docs): update build script

# Commit: refactor(middleware): improve test coverage

# Commit: refactor(di): restructure test coverage

# Commit: test(di-container): add coverage for handler resolution

# Commit: test(route): add coverage for singleton scope

# Commit: feat(http): add provider registration

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve performance

# Commit: feat(ioc): add handler resolution

# Commit: test(container): add coverage for context propagation

# Commit: refactor(module): restructure error messages

# Commit: refactor(controller): improve type safety

# Commit: perf(router): optimize middleware chain

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure code structure

# Commit: fix(core): resolve body parsing

# Commit: chore(example): update dependencies

# Commit: feat(hello): add request injection

# Commit: fix(test): resolve nil pointer

# Commit: feat(docs): add context propagation

# Commit: refactor(middleware): restructure performance

# Commit: fix(di): handle pattern matching case

# Commit: docs(di-container): update documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): improve error messages

# Commit: refactor(ioc): improve code structure

# Commit: fix(container): resolve type inference

# Commit: refactor(module): restructure type safety

# Commit: chore(container): update gitignore

# Commit: chore(module): update dependencies

# Commit: feat(controller): add middleware chain

# Commit: feat(router): implement error handling

# Commit: fix(server): handle routing conflict case

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): resolve routing conflict

# Commit: fix(example): resolve type inference

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure error messages

# Commit: refactor(middleware): improve memory usage

# Commit: chore(di): update dependencies

# Commit: feat(di-container): add route matching

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update test suite

# Commit: fix(ioc): handle scope resolution case

# Commit: refactor(container): restructure memory usage

# Commit: chore(module): update ci configuration

# Commit: fix(controller): resolve nil pointer

# Commit: perf(module): optimize controller routing

# Commit: chore(controller): update gitignore

# Commit: chore(router): update dependencies

# Commit: feat(server): add singleton scope

# Commit: feat(provider): add middleware chain

# Commit: feat(core): implement context propagation

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): resolve nil pointer

# Commit: feat(test): add context propagation

# Commit: refactor(docs): restructure memory usage

# Commit: chore(middleware): update test suite

# Commit: test(di): add coverage for context propagation

# Commit: refactor(di-container): restructure performance

# Commit: feat(route): add provider registration

# Commit: docs(http): update documentation

# Commit: chore(ioc): update build script

# Commit: refactor(container): improve performance

# Commit: fix(module): handle path extraction case

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure error messages

# Commit: perf(controller): optimize controller routing

# Commit: chore(router): update license

# Commit: perf(server): optimize singleton scope

# Commit: feat(provider): add route matching

# Commit: refactor(core): improve memory usage

# Commit: chore(example): update test suite

# Commit: test(hello): add coverage for controller routing

# Commit: chore(test): update build script

# Commit: refactor(docs): improve documentation

# Commit: docs(middleware): update documentation

# Commit: feat(di): add context propagation

# Commit: refactor(di-container): restructure documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): improve concurrency handling

# Commit: feat(ioc): implement middleware chain

# Commit: feat(container): implement middleware chain

# Commit: feat(module): implement module setup

# Commit: perf(controller): optimize context propagation

# Commit: refactor(router): restructure concurrency handling

# Commit: feat(server): implement context propagation

# Commit: chore(router): update go mod

# Commit: refactor(server): restructure concurrency handling

# Commit: feat(provider): implement response writing

# Commit: docs(core): update documentation

# Commit: refactor(example): improve documentation

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for handler resolution

# Commit: test(di): add coverage for route matching

# Commit: refactor(di-container): improve memory usage

# Commit: chore(route): update gitignore

# Commit: chore(http): update build script

# Commit: refactor(ioc): improve memory usage

# Commit: refactor(container): restructure test coverage

# Commit: refactor(module): restructure memory usage

# Commit: chore(controller): update readme

# Commit: docs(router): update documentation

# Commit: docs(server): update documentation

# Commit: feat(provider): implement error handling

# Commit: fix(server): handle header setting case

# Commit: perf(provider): optimize middleware chain

# Commit: feat(core): implement handler resolution

# Commit: test(example): add coverage for module setup

# Commit: perf(hello): optimize middleware chain

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize route matching

# Commit: refactor(di-container): improve concurrency handling

# Commit: refactor(route): restructure error messages

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update readme

# Commit: docs(container): update documentation

# Commit: feat(module): implement error handling

# Commit: fix(controller): handle nil pointer case

# Commit: feat(router): add singleton scope

# Commit: feat(server): add response writing

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): resolve scope resolution

# Commit: docs(provider): update documentation

# Commit: fix(core): resolve header setting

# Commit: perf(example): optimize param extraction

# Commit: feat(hello): add provider registration

# Commit: docs(test): update documentation

# Commit: chore(docs): update gitignore

# Commit: fix(middleware): handle scope resolution case

# Commit: refactor(di): improve type safety

# Commit: perf(di-container): optimize param extraction

# Commit: feat(route): add middleware chain

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure error messages

# Commit: refactor(container): improve code structure

# Commit: fix(module): resolve type inference

# Commit: test(controller): add coverage for param extraction

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: refactor(provider): restructure type safety

# Commit: refactor(core): restructure memory usage

# Commit: chore(example): update readme

# Commit: docs(core): update documentation

# Commit: feat(example): add response writing

# Commit: refactor(hello): improve type safety

# Commit: refactor(test): restructure test coverage

# Commit: test(docs): add coverage for request injection

# Commit: fix(middleware): resolve scope resolution

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for provider registration

# Commit: docs(http): update documentation

# Commit: chore(ioc): update readme

# Commit: fix(container): handle pattern matching case

# Commit: docs(module): update documentation

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize error handling

# Commit: fix(server): handle body parsing case

# Commit: fix(provider): handle scope resolution case

# Commit: refactor(core): improve error messages

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for param extraction

# Commit: test(example): add coverage for error handling

# Commit: fix(hello): handle body parsing case

# Commit: chore(test): update gitignore

# Commit: chore(docs): update readme

# Commit: docs(middleware): update documentation

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize request injection

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add context propagation

# Commit: refactor(container): restructure concurrency handling

# Commit: feat(module): implement singleton scope

# Commit: feat(controller): add param extraction

# Commit: test(router): add coverage for response writing

# Commit: refactor(server): restructure test coverage

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize context propagation

# Commit: refactor(example): restructure documentation

# Commit: docs(hello): update documentation

# Commit: fix(test): handle header setting case

# Commit: refactor(hello): improve type safety

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize route matching

# Commit: refactor(middleware): improve performance

# Commit: fix(di): handle pattern matching case

# Commit: docs(di-container): update documentation

# Commit: perf(route): optimize singleton scope

# Commit: feat(http): add context propagation

# Commit: refactor(ioc): restructure concurrency handling

# Commit: feat(container): implement route matching

# Commit: refactor(module): improve code structure

# Commit: fix(controller): resolve type inference

# Commit: test(router): add coverage for module setup

# Commit: perf(server): optimize controller routing

# Commit: chore(provider): update gitignore

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve header setting

# Commit: perf(hello): optimize singleton scope

# Commit: feat(test): add module setup

# Commit: perf(docs): optimize handler resolution

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize response writing

# Commit: test(middleware): add coverage for controller routing

# Commit: chore(di): update ci configuration

# Commit: refactor(di-container): restructure code structure

# Commit: fix(route): resolve routing conflict

# Commit: fix(http): handle routing conflict case

# Commit: fix(ioc): resolve scope resolution

# Commit: refactor(container): improve performance

# Commit: feat(module): add provider registration

# Commit: docs(controller): update documentation

# Commit: chore(router): update ci configuration

# Commit: fix(server): resolve pattern matching

# Commit: docs(provider): update documentation

# Commit: chore(core): update gitignore

# Commit: fix(example): handle body parsing case

# Commit: chore(hello): update ci configuration

# Commit: fix(test): resolve path extraction

# Commit: feat(docs): implement route matching

# Commit: refactor(middleware): improve concurrency handling

# Commit: fix(docs): resolve routing conflict

# Commit: fix(middleware): resolve path extraction

# Commit: feat(di): implement response writing

# Commit: fix(di-container): resolve path extraction

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: refactor(ioc): improve type safety

# Commit: perf(container): optimize provider registration

# Commit: docs(module): update documentation

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve test coverage

# Commit: test(server): add coverage for error handling

# Commit: fix(provider): handle header setting case

# Commit: perf(core): optimize response writing

# Commit: chore(example): update gitignore

# Commit: fix(hello): handle path extraction case

# Commit: feat(test): implement error handling

# Commit: fix(docs): handle routing conflict case

# Commit: fix(middleware): resolve scope resolution

# Commit: fix(di): handle routing conflict case

# Commit: refactor(middleware): improve code structure

# Commit: fix(di): resolve routing conflict

# Commit: fix(di-container): handle nil pointer case

# Commit: feat(route): add request injection

# Commit: fix(http): resolve body parsing

# Commit: refactor(ioc): restructure concurrency handling

# Commit: feat(container): implement route matching

# Commit: refactor(module): improve documentation

# Commit: refactor(controller): restructure test coverage

# Commit: test(router): add coverage for controller routing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement context propagation

# Commit: refactor(core): restructure concurrency handling

# Commit: feat(example): implement response writing

# Commit: test(hello): add coverage for handler resolution

# Commit: test(test): add coverage for response writing

# Commit: perf(docs): optimize controller routing

# Commit: chore(middleware): update readme

# Commit: docs(di): update documentation

# Commit: fix(di-container): handle header setting case

# Commit: test(di): add coverage for context propagation

# Commit: refactor(di-container): restructure code structure

# Commit: refactor(route): restructure test coverage

# Commit: test(http): add coverage for error handling

# Commit: fix(ioc): handle routing conflict case

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize singleton scope

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: perf(provider): optimize response writing

# Commit: fix(core): resolve header setting

# Commit: perf(example): optimize param extraction

# Commit: docs(hello): update documentation

# Commit: feat(test): implement request injection

# Commit: fix(docs): resolve header setting

# Commit: fix(middleware): handle header setting case

# Commit: perf(di): optimize handler resolution

# Commit: test(di-container): add coverage for module setup

# Commit: perf(route): optimize request injection

# Commit: feat(di-container): implement handler resolution

# Commit: test(route): add coverage for route matching

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: perf(module): optimize controller routing

# Commit: chore(controller): update test suite

# Commit: test(router): add coverage for module setup

# Commit: perf(server): optimize response writing

# Commit: test(provider): add coverage for response writing

# Commit: test(core): add coverage for error handling

# Commit: fix(example): handle nil pointer case

# Commit: feat(hello): add response writing

# Commit: test(test): add coverage for middleware chain

# Commit: feat(docs): implement route matching

# Commit: refactor(middleware): improve type safety

# Commit: perf(di): optimize route matching

# Commit: refactor(di-container): improve code structure

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update go mod

# Commit: fix(route): handle scope resolution case

# Commit: refactor(http): improve code structure

# Commit: fix(ioc): resolve header setting

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle pattern matching case

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize error handling

# Commit: fix(server): handle type inference case

# Commit: test(provider): add coverage for request injection

# Commit: fix(core): resolve scope resolution

# Commit: refactor(example): improve test coverage

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve pattern matching

# Commit: docs(docs): update documentation

# Commit: chore(middleware): update test suite

# Commit: test(di): add coverage for provider registration

# Commit: docs(di-container): update documentation

# Commit: refactor(route): restructure error messages

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for response writing

# Commit: docs(http): update documentation

# Commit: fix(ioc): resolve body parsing

# Commit: chore(container): update build script

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: chore(router): update test suite

# Commit: fix(server): handle nil pointer case

# Commit: refactor(provider): restructure performance

# Commit: feat(core): add route matching

# Commit: refactor(example): improve memory usage

# Commit: chore(hello): update go mod

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve body parsing

# Commit: chore(middleware): update ci configuration

# Commit: fix(di): resolve body parsing

# Commit: chore(di-container): update build script

# Commit: refactor(route): improve documentation

# Commit: fix(http): handle routing conflict case

# Commit: fix(ioc): resolve path extraction

# Commit: feat(container): implement provider registration

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize param extraction

# Commit: perf(module): optimize middleware chain

# Commit: feat(controller): implement response writing

# Commit: refactor(router): improve concurrency handling

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update go mod

# Commit: feat(core): implement module setup

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle type inference case

# Commit: test(docs): add coverage for route matching

# Commit: refactor(middleware): improve documentation

# Commit: refactor(di): restructure memory usage

# Commit: chore(di-container): update gitignore

# Commit: chore(route): update readme

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add handler resolution

# Commit: test(module): add coverage for provider registration

# Commit: feat(container): implement request injection

# Commit: fix(module): resolve nil pointer

# Commit: feat(controller): add provider registration

# Commit: docs(router): update documentation

# Commit: refactor(server): restructure test coverage

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): restructure code structure

# Commit: fix(example): handle scope resolution case

# Commit: refactor(hello): improve performance

# Commit: feat(test): add singleton scope

# Commit: feat(docs): add middleware chain

# Commit: feat(middleware): implement request injection

# Commit: fix(di): resolve nil pointer

# Commit: feat(di-container): add middleware chain

# Commit: feat(route): implement param extraction

# Commit: feat(http): add singleton scope

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle body parsing case

# Commit: chore(module): update license

# Commit: perf(controller): optimize route matching

# Commit: fix(module): resolve routing conflict

# Commit: fix(controller): resolve nil pointer

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: fix(provider): resolve path extraction

# Commit: feat(core): implement provider registration

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure memory usage

# Commit: chore(test): update dependencies

# Commit: feat(docs): add singleton scope

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize response writing

# Commit: feat(di-container): add controller routing

# Commit: chore(route): update ci configuration

# Commit: fix(http): resolve routing conflict

# Commit: refactor(ioc): restructure code structure

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve concurrency handling

# Commit: feat(controller): implement middleware chain

# Commit: feat(router): implement singleton scope

# Commit: docs(controller): update documentation

# Commit: fix(router): resolve path extraction

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement middleware chain

# Commit: feat(core): implement param extraction

# Commit: fix(example): handle nil pointer case

# Commit: feat(hello): add handler resolution

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve code structure

# Commit: refactor(middleware): restructure type safety

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add error handling

# Commit: fix(route): handle path extraction case

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure code structure

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve type safety

# Commit: perf(controller): optimize handler resolution

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle body parsing case

# Commit: perf(router): optimize middleware chain

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for route matching

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle type inference case

# Commit: test(di): add coverage for error handling

# Commit: fix(di-container): handle routing conflict case

# Commit: fix(route): handle routing conflict case

# Commit: fix(http): resolve body parsing

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: test(module): add coverage for handler resolution

# Commit: test(controller): add coverage for route matching

# Commit: refactor(router): improve performance

# Commit: feat(server): add request injection

# Commit: fix(provider): resolve type inference

# Commit: test(server): add coverage for middleware chain

# Commit: feat(provider): implement module setup

# Commit: perf(core): optimize module setup

# Commit: perf(example): optimize handler resolution

# Commit: test(hello): add coverage for request injection

# Commit: fix(test): resolve pattern matching

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): restructure error messages

# Commit: refactor(di): improve concurrency handling

# Commit: feat(di-container): implement param extraction

# Commit: refactor(route): improve performance

# Commit: feat(http): add middleware chain

# Commit: feat(ioc): implement module setup

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update gitignore

# Commit: fix(controller): handle nil pointer case

# Commit: feat(router): add route matching

# Commit: refactor(server): improve concurrency handling

# Commit: feat(provider): implement param extraction

# Commit: refactor(core): restructure memory usage

# Commit: refactor(provider): restructure concurrency handling

# Commit: refactor(core): restructure concurrency handling

# Commit: refactor(example): restructure type safety

# Commit: refactor(hello): restructure performance

# Commit: feat(test): add singleton scope

# Commit: feat(docs): add handler resolution

# Commit: test(middleware): add coverage for module setup

# Commit: perf(di): optimize controller routing

# Commit: chore(di-container): update ci configuration

# Commit: fix(route): resolve routing conflict

# Commit: fix(http): resolve header setting

# Commit: perf(ioc): optimize singleton scope

# Commit: feat(container): add route matching

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: fix(router): handle path extraction case

# Commit: feat(server): implement module setup

# Commit: perf(provider): optimize context propagation

# Commit: refactor(core): restructure memory usage

# Commit: chore(example): update build script

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add request injection

# Commit: fix(hello): resolve body parsing

# Commit: chore(test): update gitignore

# Commit: chore(docs): update gitignore

# Commit: fix(middleware): handle type inference case

# Commit: fix(di): handle type inference case

# Commit: test(di-container): add coverage for context propagation

# Commit: refactor(route): restructure concurrency handling

# Commit: feat(http): implement module setup

# Commit: perf(ioc): optimize module setup

# Commit: perf(container): optimize handler resolution

# Commit: test(module): add coverage for handler resolution

# Commit: test(controller): add coverage for route matching

# Commit: refactor(router): improve concurrency handling

# Commit: feat(server): implement route matching

# Commit: refactor(provider): improve memory usage

# Commit: fix(core): handle body parsing case

# Commit: chore(example): update license

# Commit: perf(hello): optimize request injection

# Commit: refactor(example): restructure type safety

# Commit: perf(hello): optimize middleware chain

# Commit: feat(test): implement controller routing

# Commit: chore(docs): update test suite

# Commit: test(middleware): add coverage for middleware chain

# Commit: feat(di): implement response writing

# Commit: fix(di-container): handle path extraction case

# Commit: feat(route): implement singleton scope

# Commit: feat(http): add context propagation

# Commit: refactor(ioc): restructure concurrency handling

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: test(router): add coverage for controller routing

# Commit: chore(server): update license

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve scope resolution

# Commit: refactor(example): improve code structure

# Commit: fix(hello): resolve header setting

# Commit: perf(test): optimize context propagation

# Commit: test(hello): add coverage for param extraction

# Commit: fix(test): resolve pattern matching

# Commit: fix(docs): handle body parsing case

# Commit: chore(middleware): update readme

# Commit: fix(di): handle scope resolution case

# Commit: refactor(di-container): improve code structure

# Commit: fix(route): resolve routing conflict

# Commit: fix(http): resolve nil pointer

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: docs(router): update documentation

# Commit: fix(server): resolve routing conflict

# Commit: fix(provider): resolve type inference

# Commit: test(core): add coverage for error handling

# Commit: fix(example): handle scope resolution case

# Commit: refactor(hello): improve error messages

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for route matching

# Commit: refactor(test): improve performance

# Commit: feat(docs): add controller routing

# Commit: chore(middleware): update go mod

# Commit: feat(di): implement provider registration

# Commit: docs(di-container): update documentation

# Commit: chore(route): update build script

# Commit: refactor(http): improve test coverage

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve header setting

# Commit: perf(module): optimize controller routing

# Commit: chore(controller): update gitignore

# Commit: chore(router): update go mod

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update test suite

# Commit: refactor(core): restructure memory usage

# Commit: chore(example): update dependencies

# Commit: feat(hello): add handler resolution

# Commit: test(test): add coverage for request injection

# Commit: fix(docs): resolve body parsing

# Commit: chore(middleware): update build script

# Commit: docs(docs): update documentation

# Commit: feat(middleware): implement route matching

# Commit: refactor(di): improve performance

# Commit: feat(di-container): add singleton scope

# Commit: feat(route): add module setup

# Commit: perf(http): optimize error handling

# Commit: fix(ioc): handle path extraction case

# Commit: feat(container): implement middleware chain

# Commit: feat(module): implement response writing

# Commit: feat(controller): add module setup

# Commit: perf(router): optimize request injection

# Commit: fix(server): resolve body parsing

# Commit: fix(provider): handle header setting case

# Commit: perf(core): optimize middleware chain

# Commit: feat(example): implement module setup

# Commit: perf(hello): optimize middleware chain

# Commit: feat(test): implement route matching

# Commit: refactor(docs): improve memory usage

# Commit: fix(middleware): handle type inference case

# Commit: fix(di): handle type inference case

# Commit: fix(middleware): resolve header setting

# Commit: perf(di): optimize route matching

# Commit: refactor(di-container): improve error messages

# Commit: refactor(route): improve performance

# Commit: feat(http): add param extraction

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update readme

# Commit: docs(router): update documentation

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve memory usage

# Commit: chore(example): update build script

# Commit: refactor(hello): improve documentation

# Commit: fix(test): handle type inference case

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): resolve routing conflict

# Commit: fix(di): resolve pattern matching

# Commit: fix(di-container): handle routing conflict case

# Commit: refactor(di): improve memory usage

# Commit: chore(di-container): update test suite

# Commit: fix(route): handle header setting case

# Commit: perf(http): optimize controller routing

# Commit: chore(ioc): update readme

# Commit: refactor(container): restructure documentation

# Commit: docs(module): update documentation

# Commit: refactor(controller): restructure code structure

# Commit: fix(router): resolve pattern matching

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve documentation

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): resolve pattern matching

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve header setting

# Commit: refactor(middleware): restructure performance

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle path extraction case

# Commit: feat(route): implement route matching

# Commit: refactor(di-container): improve error messages

# Commit: fix(route): handle body parsing case

# Commit: chore(http): update dependencies

# Commit: feat(ioc): add handler resolution

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add context propagation

# Commit: refactor(controller): restructure code structure

# Commit: refactor(router): restructure concurrency handling

# Commit: feat(server): implement param extraction

# Commit: feat(provider): implement context propagation

# Commit: refactor(core): restructure memory usage

# Commit: chore(example): update gitignore

# Commit: chore(hello): update license

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: fix(middleware): handle nil pointer case

# Commit: feat(di): add middleware chain

# Commit: feat(di-container): implement param extraction

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update test suite

# Commit: feat(route): implement error handling

# Commit: fix(http): handle pattern matching case

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for response writing

# Commit: test(module): add coverage for param extraction

# Commit: docs(controller): update documentation

# Commit: fix(router): handle path extraction case

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure type safety

# Commit: fix(core): handle body parsing case

# Commit: chore(example): update test suite

# Commit: test(hello): add coverage for error handling

# Commit: fix(test): handle type inference case

# Commit: test(docs): add coverage for request injection

# Commit: fix(middleware): resolve pattern matching

# Commit: docs(di): update documentation

# Commit: docs(di-container): update documentation

# Commit: docs(route): update documentation

# Commit: feat(http): add provider registration

# Commit: docs(ioc): update documentation

# Commit: feat(http): implement param extraction

# Commit: perf(ioc): optimize response writing

# Commit: perf(container): optimize route matching

# Commit: refactor(module): improve error messages

# Commit: refactor(controller): restructure code structure

# Commit: fix(router): handle routing conflict case

# Commit: fix(server): resolve type inference

# Commit: test(provider): add coverage for handler resolution

# Commit: test(core): add coverage for param extraction

# Commit: feat(example): add context propagation

# Commit: refactor(hello): restructure concurrency handling

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve nil pointer

# Commit: fix(middleware): handle routing conflict case

# Commit: refactor(di): restructure code structure

# Commit: fix(di-container): resolve type inference

# Commit: test(route): add coverage for controller routing

# Commit: chore(http): update build script

# Commit: refactor(ioc): improve concurrency handling

# Commit: feat(container): implement provider registration

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve documentation

# Commit: docs(module): update documentation

# Commit: fix(controller): handle path extraction case

# Commit: feat(router): implement middleware chain

# Commit: feat(server): implement provider registration

# Commit: docs(provider): update documentation

# Commit: fix(core): handle path extraction case

# Commit: refactor(example): restructure error messages

# Commit: refactor(hello): improve documentation

# Commit: docs(test): update documentation

# Commit: refactor(docs): improve documentation

# Commit: docs(middleware): update documentation

# Commit: feat(di): implement param extraction

# Commit: chore(di-container): update dependencies

# Commit: feat(route): add controller routing

# Commit: chore(http): update dependencies

# Commit: feat(ioc): add controller routing

# Commit: chore(container): update build script

# Commit: refactor(module): improve performance

# Commit: feat(container): add context propagation

# Commit: refactor(module): restructure test coverage

# Commit: test(controller): add coverage for error handling

# Commit: fix(router): handle routing conflict case

# Commit: fix(server): resolve header setting

# Commit: perf(provider): optimize singleton scope

# Commit: feat(core): add middleware chain

# Commit: feat(example): implement provider registration

# Commit: docs(hello): update documentation

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for module setup

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add response writing

# Commit: docs(route): update documentation

# Commit: feat(http): implement middleware chain

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve type safety

# Commit: perf(module): optimize request injection

# Commit: fix(controller): resolve body parsing

# Commit: chore(module): update gitignore

# Commit: refactor(controller): restructure error messages

# Commit: refactor(router): improve memory usage

# Commit: chore(server): update test suite

# Commit: test(provider): add coverage for context propagation

# Commit: refactor(core): restructure documentation

# Commit: docs(example): update documentation

# Commit: feat(hello): implement middleware chain

# Commit: feat(test): implement provider registration

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): restructure documentation

# Commit: docs(di): update documentation

# Commit: chore(di-container): update build script

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update dependencies

# Commit: feat(ioc): add route matching

# Commit: refactor(container): improve error messages

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update gitignore

# Commit: chore(router): update gitignore

# Commit: feat(controller): implement module setup

# Commit: perf(router): optimize provider registration

# Commit: docs(server): update documentation

# Commit: chore(provider): update dependencies

# Commit: feat(core): add route matching

# Commit: refactor(example): improve error messages

# Commit: refactor(hello): improve performance

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve path extraction

# Commit: feat(middleware): implement controller routing

# Commit: chore(di): update gitignore

# Commit: chore(di-container): update build script

# Commit: refactor(route): restructure error messages

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update license

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update dependencies

# Commit: fix(controller): handle type inference case

# Commit: test(router): add coverage for singleton scope

# Commit: feat(server): add controller routing

# Commit: docs(router): update documentation

# Commit: feat(server): implement param extraction

# Commit: chore(provider): update dependencies

# Commit: feat(core): add singleton scope

# Commit: feat(example): add context propagation

# Commit: refactor(hello): restructure code structure

# Commit: fix(test): resolve scope resolution

# Commit: refactor(docs): improve code structure

# Commit: fix(middleware): resolve path extraction

# Commit: feat(di): implement handler resolution

# Commit: test(di-container): add coverage for module setup

# Commit: perf(route): optimize route matching

# Commit: refactor(http): improve code structure

# Commit: fix(ioc): resolve type inference

# Commit: test(container): add coverage for controller routing

# Commit: chore(module): update ci configuration

# Commit: fix(controller): resolve pattern matching

# Commit: docs(router): update documentation

# Commit: chore(server): update readme

# Commit: fix(provider): handle nil pointer case

# Commit: test(server): add coverage for provider registration

# Commit: docs(provider): update documentation

# Commit: fix(core): resolve pattern matching

# Commit: docs(example): update documentation

# Commit: fix(hello): handle scope resolution case

# Commit: refactor(test): improve error messages

# Commit: refactor(docs): improve memory usage

# Commit: chore(middleware): update dependencies

# Commit: fix(di): handle nil pointer case

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement handler resolution

# Commit: test(http): add coverage for singleton scope

# Commit: feat(ioc): add context propagation

# Commit: refactor(container): restructure memory usage

# Commit: refactor(module): restructure concurrency handling

# Commit: feat(controller): implement middleware chain

# Commit: feat(router): implement middleware chain

# Commit: feat(server): implement response writing

# Commit: fix(provider): resolve pattern matching

# Commit: docs(core): update documentation

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve nil pointer

# Commit: feat(hello): add middleware chain

# Commit: feat(test): implement route matching

# Commit: refactor(docs): improve code structure

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update gitignore

# Commit: chore(di-container): update dependencies

# Commit: feat(route): add route matching

# Commit: refactor(http): improve documentation

# Commit: fix(ioc): handle type inference case

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add route matching

# Commit: refactor(controller): improve concurrency handling

# Commit: feat(router): implement route matching

# Commit: refactor(server): improve documentation

# Commit: docs(provider): update documentation

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add error handling

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve body parsing

# Commit: chore(docs): update test suite

# Commit: test(middleware): add coverage for context propagation

# Commit: refactor(di): restructure code structure

# Commit: fix(di-container): resolve nil pointer

# Commit: feat(route): add context propagation

# Commit: refactor(http): restructure error messages

# Commit: refactor(ioc): improve performance

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve documentation

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve performance

# Commit: feat(example): add provider registration

# Commit: docs(hello): update documentation

# Commit: test(example): add coverage for response writing

# Commit: perf(hello): optimize context propagation

# Commit: refactor(test): restructure performance

# Commit: fix(docs): handle nil pointer case

# Commit: fix(middleware): handle header setting case

# Commit: perf(di): optimize response writing

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement middleware chain

# Commit: feat(http): implement module setup

# Commit: perf(ioc): optimize module setup

# Commit: perf(container): optimize provider registration

# Commit: docs(module): update documentation

# Commit: feat(controller): implement context propagation

# Commit: refactor(router): restructure error messages

# Commit: refactor(server): improve memory usage

# Commit: chore(provider): update gitignore

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve scope resolution

# Commit: refactor(hello): improve error messages

# Commit: refactor(test): improve performance

# Commit: perf(hello): optimize module setup

# Commit: perf(test): optimize route matching

# Commit: refactor(docs): improve documentation

# Commit: fix(middleware): handle path extraction case

# Commit: feat(di): implement response writing

# Commit: fix(di-container): handle type inference case

# Commit: test(route): add coverage for singleton scope

# Commit: feat(http): add handler resolution

# Commit: test(ioc): add coverage for request injection

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): improve test coverage

# Commit: test(controller): add coverage for module setup

# Commit: perf(router): optimize provider registration

# Commit: docs(server): update documentation

# Commit: perf(provider): optimize singleton scope

# Commit: feat(core): add request injection

# Commit: fix(example): resolve path extraction

# Commit: feat(hello): implement singleton scope

# Commit: feat(test): add module setup

# Commit: perf(docs): optimize context propagation

# Commit: fix(test): handle routing conflict case

# Commit: fix(docs): resolve scope resolution

# Commit: fix(middleware): handle header setting case

# Commit: perf(di): optimize request injection

# Commit: fix(di-container): resolve nil pointer

# Commit: feat(route): add context propagation

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement context propagation

# Commit: refactor(container): restructure performance

# Commit: refactor(module): restructure test coverage

# Commit: test(controller): add coverage for singleton scope

# Commit: feat(router): add controller routing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement response writing

# Commit: refactor(core): improve performance

# Commit: feat(example): add error handling

# Commit: fix(hello): handle nil pointer case

# Commit: feat(test): add route matching

# Commit: refactor(docs): improve test coverage

# Commit: test(middleware): add coverage for context propagation

# Commit: fix(docs): handle nil pointer case

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update dependencies

# Commit: refactor(di-container): restructure concurrency handling

# Commit: fix(route): handle nil pointer case

# Commit: feat(http): add controller routing

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update license

# Commit: perf(module): optimize error handling

# Commit: fix(controller): handle header setting case

# Commit: perf(router): optimize error handling

# Commit: fix(server): handle scope resolution case

# Commit: refactor(provider): improve code structure

# Commit: fix(core): resolve type inference

# Commit: fix(example): handle type inference case

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve scope resolution

# Commit: refactor(docs): improve error messages

# Commit: refactor(middleware): restructure type safety

# Commit: perf(di): optimize middleware chain

# Commit: fix(middleware): handle type inference case

# Commit: test(di): add coverage for provider registration

# Commit: docs(di-container): update documentation

# Commit: feat(route): implement singleton scope

# Commit: feat(http): add error handling

# Commit: fix(ioc): handle type inference case

# Commit: test(container): add coverage for error handling

# Commit: fix(module): handle routing conflict case

# Commit: fix(controller): resolve routing conflict

# Commit: fix(router): resolve path extraction

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update dependencies

# Commit: feat(example): add response writing

# Commit: feat(hello): add handler resolution

# Commit: test(test): add coverage for module setup

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement controller routing

# Commit: chore(di): update dependencies

# Commit: fix(di-container): handle routing conflict case

# Commit: perf(di): optimize middleware chain

# Commit: feat(di-container): implement param extraction

# Commit: feat(route): implement provider registration

# Commit: docs(http): update documentation

# Commit: fix(ioc): handle path extraction case

# Commit: feat(container): implement error handling

# Commit: fix(module): handle pattern matching case

# Commit: refactor(controller): restructure documentation

# Commit: docs(router): update documentation

# Commit: feat(server): add singleton scope

# Commit: feat(provider): add request injection

# Commit: fix(core): resolve scope resolution

# Commit: refactor(example): improve concurrency handling

# Commit: feat(hello): implement param extraction

# Commit: docs(test): update documentation

# Commit: chore(docs): update license

# Commit: refactor(middleware): restructure performance

# Commit: feat(di): add singleton scope

# Commit: feat(di-container): add error handling

# Commit: fix(route): handle header setting case

# Commit: test(di-container): add coverage for module setup

# Commit: perf(route): optimize error handling

# Commit: fix(http): handle header setting case

# Commit: perf(ioc): optimize response writing

# Commit: fix(container): resolve body parsing

# Commit: refactor(module): restructure type safety

# Commit: perf(controller): optimize error handling

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve error messages

# Commit: refactor(provider): improve concurrency handling

# Commit: feat(core): implement route matching

# Commit: refactor(example): improve documentation

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize request injection

# Commit: fix(docs): resolve pattern matching

# Commit: refactor(middleware): restructure memory usage

# Commit: chore(di): update ci configuration

# Commit: fix(di-container): resolve nil pointer

# Commit: fix(route): handle pattern matching case

# Commit: docs(http): update documentation

# Commit: fix(route): handle pattern matching case

# Commit: docs(http): update documentation

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize handler resolution

# Commit: test(router): add coverage for middleware chain

# Commit: feat(server): implement handler resolution

# Commit: test(provider): add coverage for module setup

# Commit: perf(core): optimize route matching

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for context propagation

# Commit: refactor(test): restructure code structure

# Commit: fix(docs): resolve header setting

# Commit: perf(middleware): optimize handler resolution

# Commit: test(di): add coverage for module setup

# Commit: perf(di-container): optimize provider registration

# Commit: docs(route): update documentation

# Commit: perf(http): optimize singleton scope

# Commit: feat(ioc): add controller routing

# Commit: feat(http): implement route matching

# Commit: refactor(ioc): improve documentation

# Commit: docs(container): update documentation

# Commit: chore(module): update dependencies

# Commit: fix(controller): handle routing conflict case

# Commit: fix(router): resolve body parsing

# Commit: chore(server): update go mod

# Commit: feat(provider): implement controller routing

# Commit: chore(core): update gitignore

# Commit: refactor(example): restructure memory usage

# Commit: chore(hello): update build script

# Commit: refactor(test): improve documentation

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): improve error messages

# Commit: refactor(di): improve error messages

# Commit: refactor(di-container): improve test coverage

# Commit: test(route): add coverage for error handling

# Commit: fix(http): handle path extraction case

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add route matching

# Commit: refactor(ioc): improve type safety

# Commit: perf(container): optimize param extraction

# Commit: perf(module): optimize param extraction

# Commit: chore(controller): update readme

# Commit: docs(router): update documentation

# Commit: chore(server): update test suite

# Commit: test(provider): add coverage for context propagation

# Commit: refactor(core): restructure type safety

# Commit: perf(example): optimize error handling

# Commit: fix(hello): handle nil pointer case

# Commit: feat(test): add error handling

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: fix(di): handle pattern matching case

# Commit: refactor(di-container): restructure memory usage

# Commit: chore(route): update readme

# Commit: docs(http): update documentation

# Commit: test(ioc): add coverage for module setup

# Commit: perf(container): optimize error handling

# Commit: fix(module): handle path extraction case

# Commit: docs(container): update documentation

# Commit: fix(module): resolve nil pointer

# Commit: feat(controller): add provider registration

# Commit: docs(router): update documentation

# Commit: fix(server): resolve header setting

# Commit: perf(provider): optimize handler resolution

# Commit: test(core): add coverage for controller routing

# Commit: chore(example): update readme

# Commit: docs(hello): update documentation

# Commit: feat(test): implement error handling

# Commit: fix(docs): handle routing conflict case

# Commit: fix(middleware): resolve body parsing

# Commit: chore(di): update go mod

# Commit: feat(di-container): implement controller routing

# Commit: chore(route): update ci configuration

# Commit: fix(http): resolve pattern matching

# Commit: fix(ioc): handle routing conflict case

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): restructure type safety

# Commit: fix(controller): handle body parsing case

# Commit: test(module): add coverage for module setup

# Commit: perf(controller): optimize response writing

# Commit: refactor(router): restructure documentation

# Commit: docs(server): update documentation

# Commit: feat(provider): add response writing

# Commit: refactor(core): restructure concurrency handling

# Commit: feat(example): implement route matching

# Commit: refactor(hello): improve documentation

# Commit: docs(test): update documentation

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add provider registration

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve body parsing

# Commit: fix(route): handle routing conflict case

# Commit: fix(http): resolve path extraction

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for route matching

# Commit: refactor(module): improve memory usage

# Commit: chore(controller): update license

# Commit: fix(router): handle body parsing case

# Commit: feat(controller): add handler resolution

# Commit: test(router): add coverage for response writing

# Commit: feat(server): add middleware chain

# Commit: feat(provider): implement request injection

# Commit: fix(core): resolve nil pointer

# Commit: feat(example): add module setup

# Commit: perf(hello): optimize singleton scope

# Commit: feat(test): add handler resolution

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle routing conflict case

# Commit: fix(di): resolve scope resolution

# Commit: refactor(di-container): improve type safety

# Commit: perf(route): optimize module setup

# Commit: perf(http): optimize route matching

# Commit: refactor(ioc): improve type safety

# Commit: fix(container): handle pattern matching case

# Commit: docs(module): update documentation

# Commit: chore(controller): update go mod

# Commit: fix(router): handle type inference case

# Commit: fix(server): handle path extraction case

# Commit: feat(router): add handler resolution

# Commit: test(server): add coverage for response writing

# Commit: fix(provider): handle header setting case

# Commit: perf(core): optimize controller routing

# Commit: chore(example): update build script

# Commit: refactor(hello): improve test coverage

# Commit: fix(test): handle scope resolution case

# Commit: fix(docs): handle path extraction case

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: refactor(di-container): improve documentation

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure error messages

# Commit: refactor(ioc): improve type safety

# Commit: perf(container): optimize handler resolution

# Commit: test(module): add coverage for request injection

# Commit: fix(controller): resolve path extraction

# Commit: feat(router): implement provider registration

# Commit: docs(server): update documentation

# Commit: chore(provider): update dependencies

# Commit: perf(server): optimize error handling

# Commit: fix(provider): handle body parsing case

# Commit: fix(core): handle pattern matching case

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve scope resolution

# Commit: refactor(test): restructure test coverage

# Commit: test(docs): add coverage for module setup

# Commit: perf(middleware): optimize context propagation

# Commit: refactor(di): restructure test coverage

# Commit: test(di-container): add coverage for error handling

# Commit: fix(route): handle routing conflict case

# Commit: fix(http): resolve path extraction

# Commit: feat(ioc): implement singleton scope

# Commit: feat(container): add route matching

# Commit: refactor(module): improve test coverage

# Commit: test(controller): add coverage for singleton scope

# Commit: feat(router): add middleware chain

# Commit: feat(server): implement request injection

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): handle nil pointer case

# Commit: fix(provider): resolve nil pointer

# Commit: feat(core): add response writing

# Commit: fix(example): resolve routing conflict

# Commit: refactor(hello): restructure error messages

# Commit: refactor(test): improve type safety

# Commit: perf(docs): optimize module setup

# Commit: perf(middleware): optimize error handling

# Commit: fix(di): handle routing conflict case

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): resolve header setting

# Commit: fix(http): handle scope resolution case

# Commit: refactor(ioc): improve test coverage

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add request injection

# Commit: fix(controller): resolve pattern matching

# Commit: docs(router): update documentation

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for context propagation

# Commit: refactor(core): restructure code structure

# Commit: fix(example): resolve path extraction

# Commit: chore(core): update license

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure test coverage

# Commit: test(test): add coverage for request injection

# Commit: fix(docs): resolve nil pointer

# Commit: feat(middleware): add error handling

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for singleton scope

# Commit: feat(http): add route matching

# Commit: refactor(ioc): improve test coverage

# Commit: test(container): add coverage for controller routing

# Commit: chore(module): update readme

# Commit: docs(controller): update documentation

# Commit: refactor(router): improve test coverage

# Commit: test(server): add coverage for provider registration

# Commit: docs(provider): update documentation

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: feat(hello): add handler resolution

# Commit: refactor(example): restructure memory usage

# Commit: chore(hello): update build script

# Commit: refactor(test): improve code structure

# Commit: fix(docs): resolve path extraction

# Commit: feat(middleware): implement request injection

# Commit: fix(di): resolve type inference

# Commit: fix(di-container): handle routing conflict case

# Commit: refactor(route): restructure type safety

# Commit: refactor(http): restructure concurrency handling

# Commit: feat(ioc): implement response writing

# Commit: fix(container): resolve body parsing

# Commit: chore(module): update readme

# Commit: docs(controller): update documentation

# Commit: perf(router): optimize error handling

# Commit: fix(server): handle nil pointer case

# Commit: feat(provider): add provider registration

# Commit: docs(core): update documentation

# Commit: chore(example): update readme

# Commit: docs(hello): update documentation

# Commit: refactor(test): improve performance

# Commit: feat(hello): add middleware chain

# Commit: feat(test): implement error handling

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: perf(di): optimize module setup

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for route matching

# Commit: refactor(http): improve performance

# Commit: feat(ioc): add middleware chain

# Commit: feat(container): implement provider registration

# Commit: docs(module): update documentation

# Commit: docs(controller): update documentation

# Commit: feat(router): add route matching

# Commit: refactor(server): improve concurrency handling

# Commit: feat(provider): implement middleware chain

# Commit: feat(core): implement request injection

# Commit: fix(example): resolve pattern matching

# Commit: docs(hello): update documentation

# Commit: fix(test): handle type inference case

# Commit: fix(docs): handle path extraction case

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add route matching

# Commit: refactor(middleware): improve concurrency handling

# Commit: feat(di): implement controller routing

# Commit: chore(di-container): update dependencies

# Commit: feat(route): add param extraction

# Commit: feat(http): add handler resolution

# Commit: test(ioc): add coverage for param extraction

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): improve concurrency handling

# Commit: feat(controller): implement route matching

# Commit: refactor(router): improve memory usage

# Commit: chore(server): update test suite

# Commit: test(provider): add coverage for provider registration

# Commit: docs(core): update documentation

# Commit: feat(example): implement response writing

# Commit: feat(hello): add middleware chain

# Commit: feat(test): implement error handling

# Commit: fix(docs): handle body parsing case

# Commit: chore(middleware): update go mod

# Commit: feat(docs): add module setup

# Commit: perf(middleware): optimize singleton scope

# Commit: feat(di): add module setup

# Commit: perf(di-container): optimize route matching

# Commit: refactor(route): improve performance

# Commit: feat(http): add middleware chain

# Commit: feat(ioc): implement response writing

# Commit: chore(container): update license

# Commit: fix(module): handle path extraction case

# Commit: feat(controller): implement controller routing

# Commit: chore(router): update go mod

# Commit: fix(server): handle path extraction case

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize provider registration

# Commit: docs(example): update documentation

# Commit: refactor(hello): improve performance

# Commit: feat(test): add module setup

# Commit: perf(docs): optimize response writing

# Commit: fix(middleware): handle nil pointer case

# Commit: feat(di): add module setup

# Commit: fix(middleware): resolve header setting

# Commit: perf(di): optimize middleware chain

# Commit: feat(di-container): implement provider registration

# Commit: docs(route): update documentation

# Commit: docs(http): update documentation

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve memory usage

# Commit: chore(module): update ci configuration

# Commit: fix(controller): resolve path extraction

# Commit: feat(router): implement provider registration

# Commit: docs(server): update documentation

# Commit: feat(provider): implement handler resolution

# Commit: test(core): add coverage for handler resolution

# Commit: test(example): add coverage for provider registration

# Commit: docs(hello): update documentation

# Commit: fix(test): resolve routing conflict

# Commit: fix(docs): resolve body parsing

# Commit: chore(middleware): update gitignore

# Commit: chore(di): update gitignore

# Commit: chore(di-container): update go mod

# Commit: fix(di): handle type inference case

# Commit: test(di-container): add coverage for controller routing

# Commit: chore(route): update test suite

# Commit: test(http): add coverage for handler resolution

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add handler resolution

# Commit: test(module): add coverage for handler resolution

# Commit: test(controller): add coverage for module setup

# Commit: perf(router): optimize provider registration

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve memory usage

# Commit: fix(example): handle type inference case

# Commit: test(hello): add coverage for singleton scope

# Commit: feat(test): add handler resolution

# Commit: test(docs): add coverage for provider registration

# Commit: docs(middleware): update documentation

# Commit: perf(di): optimize param extraction

# Commit: test(di-container): add coverage for param extraction

# Commit: fix(route): handle path extraction case

# Commit: feat(di-container): add route matching

# Commit: refactor(route): improve test coverage

# Commit: fix(http): handle header setting case

# Commit: perf(ioc): optimize controller routing

# Commit: chore(container): update test suite

# Commit: test(module): add coverage for controller routing

# Commit: chore(controller): update dependencies

# Commit: feat(router): add response writing

# Commit: feat(server): add context propagation

# Commit: refactor(provider): restructure memory usage

# Commit: chore(core): update readme

# Commit: docs(example): update documentation

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update license

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add provider registration

# Commit: docs(di): update documentation

# Commit: feat(di-container): add response writing

# Commit: feat(route): implement request injection

# Commit: fix(http): resolve header setting

# Commit: chore(route): update gitignore

# Commit: chore(http): update build script

# Commit: fix(ioc): handle body parsing case

# Commit: chore(container): update dependencies

# Commit: feat(module): add singleton scope

# Commit: feat(controller): add singleton scope

# Commit: feat(router): add response writing

# Commit: test(server): add coverage for response writing

# Commit: feat(provider): implement context propagation

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve documentation

# Commit: docs(hello): update documentation

# Commit: refactor(test): improve performance

# Commit: feat(docs): add error handling

# Commit: fix(middleware): handle routing conflict case

# Commit: fix(di): handle path extraction case

# Commit: feat(di-container): implement singleton scope

# Commit: feat(route): add singleton scope

# Commit: feat(http): add singleton scope

# Commit: feat(ioc): add request injection

# Commit: feat(http): add route matching

# Commit: refactor(ioc): improve type safety

# Commit: perf(container): optimize controller routing

# Commit: chore(module): update readme

# Commit: fix(controller): handle body parsing case

# Commit: chore(router): update build script

# Commit: refactor(server): improve concurrency handling

# Commit: feat(provider): implement controller routing

# Commit: chore(core): update gitignore

# Commit: chore(example): update readme

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize response writing

# Commit: fix(docs): resolve type inference

# Commit: test(middleware): add coverage for param extraction

# Commit: refactor(di): restructure concurrency handling

# Commit: feat(di-container): implement handler resolution

# Commit: test(route): add coverage for middleware chain

# Commit: feat(http): implement module setup

# Commit: perf(ioc): optimize request injection

# Commit: fix(container): resolve header setting

# Commit: fix(ioc): resolve type inference

# Commit: test(container): add coverage for provider registration

# Commit: docs(module): update documentation

# Commit: perf(controller): optimize param extraction

# Commit: feat(router): add handler resolution

# Commit: test(server): add coverage for controller routing

# Commit: chore(provider): update dependencies

# Commit: feat(core): add handler resolution

# Commit: test(example): add coverage for param extraction

# Commit: feat(hello): implement param extraction

# Commit: test(test): add coverage for middleware chain

# Commit: feat(docs): implement singleton scope

# Commit: feat(middleware): add middleware chain

# Commit: feat(di): implement handler resolution

# Commit: test(di-container): add coverage for route matching

# Commit: refactor(route): improve performance

# Commit: feat(http): add response writing

# Commit: perf(ioc): optimize response writing

# Commit: docs(container): update documentation

# Commit: refactor(module): improve memory usage

# Commit: feat(container): add request injection

# Commit: fix(module): resolve nil pointer

# Commit: fix(controller): handle scope resolution case

# Commit: refactor(router): improve performance

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: refactor(core): improve documentation

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve path extraction

# Commit: refactor(test): restructure type safety

# Commit: fix(docs): handle header setting case

# Commit: perf(middleware): optimize response writing

# Commit: perf(di): optimize request injection

# Commit: fix(di-container): resolve body parsing

# Commit: chore(route): update ci configuration

# Commit: fix(http): resolve routing conflict

# Commit: fix(ioc): resolve nil pointer

# Commit: feat(container): add module setup

# Commit: perf(module): optimize route matching

# Commit: refactor(controller): improve performance

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update build script

# Commit: refactor(router): improve type safety

# Commit: perf(server): optimize handler resolution

# Commit: test(provider): add coverage for provider registration

# Commit: docs(core): update documentation

# Commit: fix(example): resolve header setting

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update go mod

# Commit: feat(docs): implement context propagation

# Commit: refactor(middleware): restructure test coverage

# Commit: fix(di): handle header setting case

# Commit: perf(di-container): optimize error handling

# Commit: fix(route): handle path extraction case

# Commit: feat(http): implement middleware chain

# Commit: feat(ioc): implement route matching

# Commit: refactor(container): improve error messages

# Commit: refactor(module): improve code structure

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add route matching

# Commit: docs(controller): update documentation

# Commit: fix(router): handle path extraction case

# Commit: fix(server): handle body parsing case

# Commit: chore(provider): update gitignore

# Commit: chore(core): update test suite

# Commit: fix(example): handle body parsing case

# Commit: chore(hello): update test suite

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve documentation

# Commit: docs(middleware): update documentation

# Commit: docs(di): update documentation

# Commit: test(di-container): add coverage for middleware chain

# Commit: feat(route): implement response writing

# Commit: perf(http): optimize param extraction

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: feat(module): implement context propagation

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add request injection

# Commit: fix(server): resolve pattern matching

# Commit: feat(router): add controller routing

# Commit: chore(server): update build script

# Commit: refactor(provider): improve code structure

# Commit: fix(core): resolve path extraction

# Commit: refactor(example): restructure concurrency handling

# Commit: fix(hello): handle scope resolution case

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for controller routing

# Commit: chore(middleware): update readme

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): handle type inference case

# Commit: test(http): add coverage for provider registration

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve memory usage

# Commit: refactor(module): restructure code structure

# Commit: fix(controller): resolve header setting

# Commit: fix(router): handle body parsing case

# Commit: refactor(server): restructure memory usage

# Commit: chore(provider): update readme

# Commit: fix(server): handle routing conflict case

# Commit: fix(provider): resolve body parsing

# Commit: chore(core): update go mod

# Commit: feat(example): implement singleton scope

# Commit: feat(hello): add provider registration

# Commit: docs(test): update documentation

# Commit: perf(docs): optimize param extraction

# Commit: feat(middleware): implement controller routing

# Commit: chore(di): update build script

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement param extraction

# Commit: chore(http): update go mod

# Commit: feat(ioc): implement module setup

# Commit: perf(container): optimize middleware chain

# Commit: feat(module): implement request injection

# Commit: fix(controller): resolve type inference

# Commit: test(router): add coverage for provider registration

# Commit: docs(server): update documentation

# Commit: docs(provider): update documentation

# Commit: test(core): add coverage for singleton scope

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve path extraction

# Commit: feat(example): implement middleware chain

# Commit: feat(hello): implement context propagation

# Commit: refactor(test): restructure test coverage

# Commit: fix(docs): handle pattern matching case

# Commit: docs(middleware): update documentation

# Commit: docs(di): update documentation

# Commit: test(di-container): add coverage for response writing

# Commit: refactor(route): improve documentation

# Commit: docs(http): update documentation

# Commit: perf(ioc): optimize provider registration

# Commit: docs(container): update documentation

# Commit: perf(module): optimize middleware chain

# Commit: feat(controller): implement module setup

# Commit: perf(router): optimize response writing

# Commit: feat(server): implement middleware chain

# Commit: feat(provider): implement route matching

# Commit: refactor(core): improve documentation

# Commit: docs(example): update documentation

# Commit: test(core): add coverage for provider registration

# Commit: docs(example): update documentation

# Commit: feat(hello): add param extraction

# Commit: feat(test): add response writing

# Commit: fix(docs): resolve scope resolution

# Commit: refactor(middleware): improve error messages

# Commit: refactor(di): improve type safety

# Commit: perf(di-container): optimize handler resolution

# Commit: test(route): add coverage for response writing

# Commit: fix(http): resolve pattern matching

# Commit: docs(ioc): update documentation

# Commit: chore(container): update test suite

# Commit: test(module): add coverage for provider registration

# Commit: docs(controller): update documentation

# Commit: refactor(router): restructure concurrency handling

# Commit: feat(server): implement context propagation

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): improve test coverage

# Commit: test(example): add coverage for request injection

# Commit: fix(hello): resolve path extraction

# Commit: perf(example): optimize provider registration

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize context propagation

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): resolve body parsing

# Commit: chore(di): update readme

# Commit: docs(di-container): update documentation

# Commit: fix(route): handle pattern matching case

# Commit: docs(http): update documentation

# Commit: perf(ioc): optimize error handling

# Commit: fix(container): handle body parsing case

# Commit: chore(module): update dependencies

# Commit: feat(controller): add request injection

# Commit: fix(router): resolve nil pointer

# Commit: feat(server): add route matching

# Commit: refactor(provider): improve code structure

# Commit: fix(core): resolve nil pointer

# Commit: feat(example): add singleton scope

# Commit: feat(hello): add module setup

# Commit: perf(test): optimize provider registration

# Commit: feat(hello): add context propagation

# Commit: refactor(test): restructure concurrency handling

# Commit: feat(docs): implement context propagation

# Commit: refactor(middleware): restructure performance

# Commit: feat(di): add route matching

# Commit: refactor(di-container): improve documentation

# Commit: docs(route): update documentation

# Commit: fix(http): resolve pattern matching

# Commit: docs(ioc): update documentation

# Commit: fix(container): handle routing conflict case

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve documentation

# Commit: fix(router): handle path extraction case

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update build script

# Commit: refactor(core): improve code structure

# Commit: refactor(example): restructure performance

# Commit: feat(hello): add handler resolution

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve memory usage

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for middleware chain

# Commit: feat(middleware): implement request injection

# Commit: fix(di): resolve nil pointer

# Commit: feat(di-container): add handler resolution

# Commit: test(route): add coverage for error handling

# Commit: fix(http): handle pattern matching case

# Commit: docs(ioc): update documentation

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: perf(controller): optimize controller routing

# Commit: chore(router): update license

# Commit: perf(server): optimize singleton scope

# Commit: feat(provider): add singleton scope

# Commit: feat(core): add middleware chain

# Commit: feat(example): implement controller routing

# Commit: chore(hello): update test suite

# Commit: fix(test): handle nil pointer case

# Commit: feat(docs): add middleware chain

# Commit: feat(middleware): implement error handling

# Commit: chore(docs): update build script

# Commit: refactor(middleware): improve test coverage

# Commit: test(di): add coverage for controller routing

# Commit: chore(di-container): update gitignore

# Commit: refactor(route): restructure type safety

# Commit: perf(http): optimize provider registration

# Commit: docs(ioc): update documentation

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): improve documentation

# Commit: fix(controller): handle nil pointer case

# Commit: feat(router): add request injection

# Commit: fix(server): resolve scope resolution

# Commit: refactor(provider): improve code structure

# Commit: fix(core): resolve body parsing

# Commit: chore(example): update go mod

# Commit: feat(hello): implement provider registration

# Commit: docs(test): update documentation

# Commit: chore(docs): update test suite

# Commit: test(middleware): add coverage for controller routing

# Commit: chore(di): update ci configuration

# Commit: test(middleware): add coverage for response writing

# Commit: test(di): add coverage for error handling

# Commit: fix(di-container): handle nil pointer case

# Commit: feat(route): add singleton scope

# Commit: feat(http): add handler resolution

# Commit: test(ioc): add coverage for handler resolution

# Commit: test(container): add coverage for middleware chain

# Commit: feat(module): implement module setup

# Commit: perf(controller): optimize provider registration

# Commit: docs(router): update documentation

# Commit: test(server): add coverage for error handling

# Commit: fix(provider): handle scope resolution case

# Commit: refactor(core): improve type safety

# Commit: refactor(example): restructure type safety

# Commit: perf(hello): optimize response writing

# Commit: refactor(test): improve performance

# Commit: feat(docs): add provider registration

# Commit: docs(middleware): update documentation

# Commit: refactor(di): improve concurrency handling

# Commit: feat(di-container): implement route matching

# Commit: feat(di): implement controller routing

# Commit: chore(di-container): update gitignore

# Commit: chore(route): update test suite

# Commit: test(http): add coverage for request injection

# Commit: fix(ioc): resolve nil pointer

# Commit: fix(container): handle nil pointer case

# Commit: fix(module): handle header setting case

# Commit: fix(controller): handle path extraction case

# Commit: feat(router): implement module setup

# Commit: perf(server): optimize provider registration

# Commit: docs(provider): update documentation

# Commit: fix(core): resolve routing conflict

# Commit: fix(example): handle nil pointer case

# Commit: feat(hello): add param extraction

# Commit: chore(test): update ci configuration

# Commit: fix(docs): resolve nil pointer

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add handler resolution

# Commit: test(route): add coverage for route matching

# Commit: feat(di-container): implement context propagation

# Commit: refactor(route): restructure concurrency handling

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure type safety

# Commit: fix(container): handle path extraction case

# Commit: feat(module): implement route matching

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve performance

# Commit: feat(server): add singleton scope

# Commit: feat(provider): add middleware chain

# Commit: feat(core): implement param extraction

# Commit: refactor(example): restructure type safety

# Commit: perf(hello): optimize param extraction

# Commit: docs(test): update documentation

# Commit: docs(docs): update documentation

# Commit: docs(middleware): update documentation

# Commit: perf(di): optimize param extraction

# Commit: fix(di-container): handle type inference case

# Commit: test(route): add coverage for context propagation

# Commit: refactor(http): restructure concurrency handling

# Commit: fix(route): resolve nil pointer

# Commit: feat(http): add context propagation

# Commit: refactor(ioc): restructure code structure

# Commit: fix(container): resolve scope resolution

# Commit: refactor(module): improve error messages

# Commit: refactor(controller): improve test coverage

# Commit: test(router): add coverage for controller routing

# Commit: chore(server): update readme

# Commit: docs(provider): update documentation

# Commit: docs(core): update documentation

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure code structure

# Commit: fix(test): resolve body parsing

# Commit: chore(docs): update license

# Commit: perf(middleware): optimize response writing

# Commit: fix(di): resolve type inference

# Commit: test(di-container): add coverage for request injection

# Commit: fix(route): resolve nil pointer

# Commit: fix(http): handle pattern matching case

# Commit: refactor(ioc): restructure concurrency handling

# Commit: perf(http): optimize error handling

# Commit: fix(ioc): handle pattern matching case

# Commit: docs(container): update documentation

# Commit: feat(module): implement param extraction

# Commit: test(controller): add coverage for singleton scope

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: feat(provider): implement param extraction

# Commit: fix(core): handle pattern matching case

# Commit: docs(example): update documentation

# Commit: feat(hello): implement handler resolution

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize middleware chain

# Commit: feat(di): implement context propagation

# Commit: refactor(di-container): restructure error messages

# Commit: refactor(route): improve concurrency handling

# Commit: feat(http): implement singleton scope

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle pattern matching case

# Commit: refactor(ioc): improve documentation

# Commit: refactor(container): restructure memory usage

# Commit: chore(module): update go mod

# Commit: fix(controller): handle scope resolution case

# Commit: refactor(router): improve type safety

# Commit: perf(server): optimize route matching

# Commit: refactor(provider): improve performance

# Commit: feat(core): add param extraction

# Commit: refactor(example): improve test coverage

# Commit: test(hello): add coverage for singleton scope

# Commit: feat(test): add context propagation

# Commit: refactor(docs): restructure type safety

# Commit: fix(middleware): handle pattern matching case

# Commit: docs(di): update documentation

# Commit: test(di-container): add coverage for error handling

# Commit: fix(route): handle routing conflict case

# Commit: fix(http): resolve body parsing

# Commit: chore(ioc): update license

# Commit: refactor(container): restructure type safety

# Commit: fix(module): handle routing conflict case

# Commit: fix(container): handle type inference case

# Commit: test(module): add coverage for provider registration

# Commit: docs(controller): update documentation

# Commit: feat(router): add provider registration

# Commit: docs(server): update documentation

# Commit: fix(provider): handle scope resolution case

# Commit: refactor(core): improve concurrency handling

# Commit: feat(example): implement controller routing

# Commit: chore(hello): update readme

# Commit: docs(test): update documentation

# Commit: fix(docs): resolve type inference

# Commit: fix(middleware): handle path extraction case

# Commit: refactor(di): restructure concurrency handling

# Commit: feat(di-container): implement provider registration

# Commit: docs(route): update documentation

# Commit: test(http): add coverage for middleware chain

# Commit: feat(ioc): implement context propagation

# Commit: refactor(container): restructure test coverage

# Commit: refactor(module): restructure test coverage

# Commit: refactor(controller): restructure test coverage

# Commit: perf(module): optimize response writing

# Commit: fix(controller): resolve path extraction

# Commit: feat(router): implement param extraction

# Commit: docs(server): update documentation

# Commit: test(provider): add coverage for module setup

# Commit: perf(core): optimize singleton scope

# Commit: feat(example): add middleware chain

# Commit: feat(hello): implement error handling

# Commit: fix(test): handle path extraction case

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): resolve nil pointer

# Commit: feat(di): add error handling

# Commit: fix(di-container): handle scope resolution case

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update ci configuration

# Commit: fix(ioc): resolve type inference

# Commit: refactor(container): restructure memory usage

# Commit: refactor(module): restructure test coverage

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add response writing

# Commit: docs(controller): update documentation

# Commit: refactor(router): restructure performance

# Commit: fix(server): handle path extraction case

# Commit: feat(provider): implement provider registration

# Commit: docs(core): update documentation

# Commit: chore(example): update ci configuration

# Commit: fix(hello): resolve pattern matching

# Commit: docs(test): update documentation

# Commit: feat(docs): implement response writing

# Commit: test(middleware): add coverage for route matching

# Commit: refactor(di): improve concurrency handling

# Commit: feat(di-container): implement handler resolution

# Commit: test(route): add coverage for route matching

# Commit: refactor(http): improve memory usage

# Commit: chore(ioc): update go mod

# Commit: feat(container): implement context propagation

# Commit: refactor(module): restructure code structure

# Commit: fix(controller): resolve pattern matching

# Commit: refactor(router): restructure concurrency handling

# Commit: refactor(server): restructure memory usage

# Commit: perf(router): optimize context propagation

# Commit: refactor(server): restructure memory usage

# Commit: chore(provider): update gitignore

# Commit: chore(core): update ci configuration

# Commit: fix(example): resolve path extraction

# Commit: feat(hello): implement context propagation

# Commit: refactor(test): restructure type safety

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement singleton scope

# Commit: feat(di): add response writing

# Commit: test(di-container): add coverage for error handling

# Commit: fix(route): handle header setting case

# Commit: perf(http): optimize request injection

# Commit: fix(ioc): resolve header setting

# Commit: refactor(container): restructure test coverage

# Commit: test(module): add coverage for context propagation

# Commit: refactor(controller): restructure performance

# Commit: feat(router): add request injection

# Commit: fix(server): resolve type inference

# Commit: test(provider): add coverage for provider registration

# Commit: feat(server): implement request injection

# Commit: fix(provider): resolve routing conflict

# Commit: fix(core): handle body parsing case

# Commit: refactor(example): restructure performance

# Commit: feat(hello): add route matching

# Commit: refactor(test): improve concurrency handling

# Commit: refactor(docs): restructure documentation

# Commit: docs(middleware): update documentation

# Commit: refactor(di): restructure performance

# Commit: feat(di-container): add provider registration

# Commit: docs(route): update documentation

# Commit: refactor(http): restructure test coverage

# Commit: test(ioc): add coverage for param extraction

# Commit: refactor(container): restructure code structure

# Commit: fix(module): resolve body parsing

# Commit: chore(controller): update dependencies

# Commit: refactor(router): restructure test coverage

# Commit: test(server): add coverage for error handling

# Commit: fix(provider): handle routing conflict case

# Commit: fix(core): resolve type inference

# Commit: docs(provider): update documentation

# Commit: docs(core): update documentation

# Commit: fix(example): handle scope resolution case

# Commit: refactor(hello): improve type safety

# Commit: fix(test): handle scope resolution case

# Commit: refactor(docs): improve error messages

# Commit: refactor(middleware): improve performance

# Commit: feat(di): add response writing

# Commit: chore(di-container): update readme

# Commit: docs(route): update documentation

# Commit: fix(http): handle nil pointer case

# Commit: feat(ioc): add error handling

# Commit: fix(container): handle routing conflict case

# Commit: fix(module): resolve type inference

# Commit: test(controller): add coverage for response writing

# Commit: docs(router): update documentation

# Commit: test(server): add coverage for module setup

# Commit: perf(provider): optimize middleware chain

# Commit: feat(core): implement provider registration

# Commit: docs(example): update documentation

# Commit: test(core): add coverage for controller routing

# Commit: chore(example): update dependencies

# Commit: fix(hello): handle scope resolution case

# Commit: refactor(test): improve memory usage

# Commit: chore(docs): update dependencies

# Commit: feat(middleware): add module setup

# Commit: perf(di): optimize singleton scope

# Commit: feat(di-container): add request injection

# Commit: fix(route): resolve type inference

# Commit: test(http): add coverage for handler resolution

# Commit: test(ioc): add coverage for response writing

# Commit: docs(container): update documentation

# Commit: refactor(module): improve test coverage

# Commit: test(controller): add coverage for provider registration

# Commit: docs(router): update documentation

# Commit: test(server): add coverage for middleware chain

# Commit: feat(provider): implement param extraction

# Commit: chore(core): update readme

# Commit: docs(example): update documentation

# Commit: fix(hello): resolve path extraction

# Commit: perf(example): optimize param extraction

# Commit: refactor(hello): restructure concurrency handling

# Commit: refactor(test): restructure type safety

# Commit: perf(docs): optimize middleware chain

# Commit: feat(middleware): implement response writing

# Commit: feat(di): implement singleton scope

# Commit: feat(di-container): add context propagation

# Commit: refactor(route): restructure concurrency handling

# Commit: feat(http): implement handler resolution

# Commit: test(ioc): add coverage for singleton scope

# Commit: feat(container): add controller routing

# Commit: chore(module): update license

# Commit: fix(controller): handle nil pointer case

# Commit: refactor(router): restructure memory usage

# Commit: chore(server): update ci configuration

# Commit: fix(provider): resolve scope resolution

# Commit: refactor(core): improve code structure

# Commit: fix(example): resolve path extraction

# Commit: feat(hello): implement singleton scope

# Commit: feat(test): add handler resolution

# Commit: feat(hello): add param extraction

# Commit: fix(test): handle scope resolution case

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize provider registration

# Commit: docs(di): update documentation

# Commit: fix(di-container): handle path extraction case

# Commit: feat(route): implement singleton scope

# Commit: feat(http): add controller routing

# Commit: chore(ioc): update dependencies

# Commit: feat(container): add param extraction

# Commit: refactor(module): restructure documentation

# Commit: refactor(controller): restructure concurrency handling

# Commit: feat(router): implement error handling

# Commit: fix(server): handle type inference case

# Commit: test(provider): add coverage for context propagation

# Commit: refactor(core): restructure error messages

# Commit: refactor(example): improve type safety

# Commit: perf(hello): optimize handler resolution

# Commit: test(test): add coverage for context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: chore(test): update go mod

# Commit: feat(docs): implement route matching

# Commit: refactor(middleware): improve memory usage

# Commit: chore(di): update build script

# Commit: refactor(di-container): improve concurrency handling

# Commit: feat(route): implement controller routing

# Commit: chore(http): update ci configuration

# Commit: fix(ioc): resolve routing conflict

# Commit: fix(container): resolve pattern matching

# Commit: docs(module): update documentation

# Commit: chore(controller): update build script

# Commit: refactor(router): improve type safety

# Commit: perf(server): optimize param extraction

# Commit: chore(provider): update license

# Commit: perf(core): optimize controller routing

# Commit: chore(example): update license

# Commit: refactor(hello): restructure error messages

# Commit: refactor(test): restructure memory usage

# Commit: chore(docs): update gitignore

# Commit: chore(middleware): update gitignore

# Commit: feat(docs): add singleton scope

# Commit: feat(middleware): add provider registration

# Commit: docs(di): update documentation

# Commit: test(di-container): add coverage for route matching

# Commit: refactor(route): improve type safety

# Commit: fix(http): handle path extraction case

# Commit: feat(ioc): implement handler resolution

# Commit: test(container): add coverage for provider registration

# Commit: docs(module): update documentation

# Commit: fix(controller): handle pattern matching case

# Commit: docs(router): update documentation

# Commit: test(server): add coverage for response writing

# Commit: refactor(provider): improve documentation

# Commit: docs(core): update documentation

# Commit: feat(example): add controller routing

# Commit: chore(hello): update gitignore

# Commit: chore(test): update build script

# Commit: refactor(docs): improve memory usage

# Commit: refactor(middleware): restructure documentation

# Commit: docs(di): update documentation

# Commit: feat(middleware): add request injection

# Commit: fix(di): resolve path extraction

# Commit: feat(di-container): implement response writing

# Commit: test(route): add coverage for request injection

# Commit: fix(http): resolve body parsing

# Commit: refactor(ioc): restructure performance

# Commit: feat(container): add controller routing

# Commit: chore(module): update test suite

# Commit: test(controller): add coverage for provider registration

# Commit: docs(router): update documentation

# Commit: refactor(server): improve test coverage

# Commit: test(provider): add coverage for middleware chain

# Commit: feat(core): implement response writing

# Commit: docs(example): update documentation

# Commit: refactor(hello): restructure documentation

# Commit: docs(test): update documentation

# Commit: fix(docs): handle scope resolution case

# Commit: fix(middleware): handle scope resolution case

# Commit: refactor(di): improve performance

# Commit: feat(di-container): add provider registration

# Commit: fix(di): resolve path extraction

# Commit: feat(di-container): implement singleton scope

# Commit: feat(route): add error handling

# Commit: fix(http): handle pattern matching case

# Commit: docs(ioc): update documentation

# Commit: perf(container): optimize request injection

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve memory usage

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for route matching

# Commit: refactor(provider): improve performance

# Commit: feat(core): add context propagation

# Commit: refactor(example): restructure code structure

# Commit: fix(hello): resolve body parsing

# Commit: chore(test): update test suite

# Commit: test(docs): add coverage for error handling

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update build script

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement response writing

# Commit: perf(di-container): optimize provider registration

# Commit: docs(route): update documentation

# Commit: fix(http): resolve pattern matching

# Commit: docs(ioc): update documentation

# Commit: chore(container): update build script

# Commit: refactor(module): improve documentation

# Commit: docs(controller): update documentation

# Commit: refactor(router): improve error messages

# Commit: refactor(server): improve documentation

# Commit: refactor(provider): restructure error messages

# Commit: refactor(core): improve test coverage

# Commit: fix(example): handle path extraction case

# Commit: feat(hello): implement route matching

# Commit: refactor(test): improve test coverage

# Commit: test(docs): add coverage for middleware chain

# Commit: feat(middleware): implement response writing

# Commit: refactor(di): restructure test coverage

# Commit: test(di-container): add coverage for request injection

# Commit: fix(route): resolve nil pointer

# Commit: feat(http): add module setup

# Commit: feat(route): implement module setup

# Commit: perf(http): optimize response writing

# Commit: chore(ioc): update ci configuration

# Commit: fix(container): resolve body parsing

# Commit: chore(module): update test suite

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle header setting case

# Commit: perf(provider): optimize request injection

# Commit: fix(core): resolve scope resolution

# Commit: fix(example): handle path extraction case

# Commit: fix(hello): handle routing conflict case

# Commit: fix(test): resolve path extraction

# Commit: refactor(docs): restructure error messages

# Commit: refactor(middleware): improve error messages

# Commit: refactor(di): improve test coverage

# Commit: refactor(di-container): restructure type safety

# Commit: refactor(route): restructure code structure

# Commit: fix(http): resolve scope resolution

# Commit: refactor(ioc): improve documentation

# Commit: perf(http): optimize context propagation

# Commit: refactor(ioc): restructure test coverage

# Commit: test(container): add coverage for singleton scope

# Commit: feat(module): add param extraction

# Commit: feat(controller): add middleware chain

# Commit: feat(router): implement param extraction

# Commit: refactor(server): restructure error messages

# Commit: refactor(provider): restructure test coverage

# Commit: test(core): add coverage for handler resolution

# Commit: test(example): add coverage for handler resolution

# Commit: test(hello): add coverage for module setup

# Commit: perf(test): optimize response writing

# Commit: refactor(docs): improve type safety

# Commit: perf(middleware): optimize module setup

# Commit: perf(di): optimize response writing

# Commit: chore(di-container): update ci configuration

# Commit: fix(route): resolve scope resolution

# Commit: fix(http): handle header setting case

# Commit: perf(ioc): optimize request injection

# Commit: fix(container): resolve routing conflict

# Commit: docs(ioc): update documentation

# Commit: refactor(container): improve type safety

# Commit: perf(module): optimize provider registration

# Commit: docs(controller): update documentation

# Commit: feat(router): add param extraction

# Commit: refactor(server): restructure concurrency handling

# Commit: feat(provider): implement response writing

# Commit: fix(core): handle header setting case

# Commit: perf(example): optimize route matching

# Commit: refactor(hello): improve type safety

# Commit: perf(test): optimize module setup

# Commit: perf(docs): optimize request injection

# Commit: fix(middleware): resolve path extraction

# Commit: feat(di): implement handler resolution

# Commit: test(di-container): add coverage for route matching

# Commit: refactor(route): improve memory usage

# Commit: chore(http): update dependencies

# Commit: feat(ioc): add module setup

# Commit: perf(container): optimize param extraction

# Commit: fix(module): resolve path extraction

# Commit: refactor(container): restructure code structure

# Commit: fix(module): resolve path extraction

# Commit: refactor(controller): restructure concurrency handling

# Commit: feat(router): implement response writing

# Commit: feat(server): add response writing

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize request injection

# Commit: fix(example): resolve type inference

# Commit: test(hello): add coverage for param extraction

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: feat(middleware): add request injection

# Commit: fix(di): resolve pattern matching

# Commit: docs(di-container): update documentation

# Commit: test(route): add coverage for singleton scope

# Commit: feat(http): add singleton scope

# Commit: feat(ioc): add request injection

# Commit: fix(container): resolve routing conflict

# Commit: fix(module): handle header setting case

# Commit: perf(controller): optimize request injection

# Commit: feat(module): add controller routing

# Commit: chore(controller): update test suite

# Commit: test(router): add coverage for request injection

# Commit: fix(server): resolve body parsing

# Commit: refactor(provider): restructure type safety

# Commit: perf(core): optimize response writing

# Commit: test(example): add coverage for response writing

# Commit: perf(hello): optimize controller routing

# Commit: chore(test): update ci configuration

# Commit: refactor(docs): restructure performance

# Commit: feat(middleware): add response writing

# Commit: docs(di): update documentation

# Commit: perf(di-container): optimize response writing

# Commit: fix(route): handle nil pointer case

# Commit: feat(http): add response writing

# Commit: chore(ioc): update license

# Commit: fix(container): handle header setting case

# Commit: perf(module): optimize middleware chain

# Commit: feat(controller): implement param extraction

# Commit: feat(router): add controller routing

# Commit: fix(controller): handle pattern matching case

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve code structure

# Commit: refactor(provider): restructure performance

# Commit: refactor(core): restructure concurrency handling

# Commit: feat(example): implement middleware chain

# Commit: feat(hello): implement param extraction

# Commit: docs(test): update documentation

# Commit: feat(docs): implement request injection

# Commit: fix(middleware): resolve type inference

# Commit: fix(di): handle scope resolution case

# Commit: refactor(di-container): restructure concurrency handling

# Commit: feat(route): implement error handling

# Commit: fix(http): handle path extraction case

# Commit: feat(ioc): implement response writing

# Commit: test(container): add coverage for handler resolution

# Commit: test(module): add coverage for param extraction

# Commit: refactor(controller): improve concurrency handling

# Commit: feat(router): implement route matching

# Commit: refactor(server): improve documentation

# Commit: perf(router): optimize middleware chain

# Commit: feat(server): implement controller routing

# Commit: chore(provider): update readme

# Commit: docs(core): update documentation

# Commit: test(example): add coverage for provider registration

# Commit: docs(hello): update documentation

# Commit: feat(test): add request injection

# Commit: fix(docs): resolve pattern matching

# Commit: fix(middleware): handle body parsing case

# Commit: chore(di): update go mod

# Commit: feat(di-container): implement module setup

# Commit: perf(route): optimize param extraction

# Commit: feat(http): implement middleware chain

# Commit: feat(ioc): implement error handling

# Commit: fix(container): handle type inference case

# Commit: fix(module): handle pattern matching case

# Commit: docs(controller): update documentation

# Commit: fix(router): handle scope resolution case

# Commit: refactor(server): improve test coverage

# Commit: test(provider): add coverage for singleton scope

# Commit: chore(server): update go mod

# Commit: feat(provider): implement handler resolution

# Commit: test(core): add coverage for request injection

# Commit: fix(example): resolve type inference

# Commit: refactor(hello): restructure error messages

# Commit: refactor(test): improve performance

# Commit: feat(docs): add context propagation

# Commit: refactor(middleware): restructure concurrency handling

# Commit: feat(di): implement error handling

# Commit: fix(di-container): handle scope resolution case

# Commit: refactor(route): improve concurrency handling

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure performance

# Commit: feat(container): add provider registration

# Commit: docs(module): update documentation

# Commit: docs(controller): update documentation

# Commit: chore(router): update test suite

# Commit: refactor(server): restructure performance

# Commit: refactor(provider): restructure documentation

# Commit: fix(core): handle type inference case

# Commit: refactor(provider): improve concurrency handling

# Commit: feat(core): implement error handling

# Commit: fix(example): handle scope resolution case

# Commit: refactor(hello): improve memory usage

# Commit: chore(test): update license

# Commit: perf(docs): optimize error handling

# Commit: fix(middleware): handle header setting case

# Commit: refactor(di): restructure memory usage

# Commit: chore(di-container): update gitignore

# Commit: fix(route): handle header setting case

# Commit: perf(http): optimize request injection

# Commit: fix(ioc): resolve nil pointer

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve concurrency handling

# Commit: feat(controller): implement route matching

# Commit: refactor(router): improve type safety

# Commit: perf(server): optimize param extraction

# Commit: perf(provider): optimize error handling

# Commit: fix(core): handle pattern matching case

# Commit: docs(example): update documentation

# Commit: feat(core): implement response writing

# Commit: feat(example): implement module setup

# Commit: perf(hello): optimize singleton scope

# Commit: feat(test): add provider registration

# Commit: docs(docs): update documentation

# Commit: refactor(middleware): improve memory usage

# Commit: chore(di): update ci configuration

# Commit: refactor(di-container): restructure type safety

# Commit: perf(route): optimize provider registration

# Commit: docs(http): update documentation

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update go mod

# Commit: feat(module): implement handler resolution

# Commit: test(controller): add coverage for handler resolution

# Commit: test(router): add coverage for singleton scope

# Commit: feat(server): add provider registration

# Commit: docs(provider): update documentation

# Commit: feat(core): implement request injection

# Commit: fix(example): resolve body parsing

# Commit: fix(hello): handle path extraction case

# Commit: docs(example): update documentation

# Commit: test(hello): add coverage for param extraction

# Commit: docs(test): update documentation

# Commit: fix(docs): handle header setting case

# Commit: perf(middleware): optimize provider registration

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve scope resolution

# Commit: refactor(route): restructure type safety

# Commit: perf(http): optimize handler resolution

# Commit: test(ioc): add coverage for provider registration

# Commit: docs(container): update documentation

# Commit: test(module): add coverage for context propagation

# Commit: refactor(controller): restructure test coverage

# Commit: test(router): add coverage for error handling

# Commit: fix(server): handle path extraction case

# Commit: feat(provider): implement provider registration

# Commit: docs(core): update documentation

# Commit: perf(example): optimize context propagation

# Commit: refactor(hello): restructure type safety

# Commit: refactor(test): restructure concurrency handling

# Commit: docs(hello): update documentation

# Commit: fix(test): handle pattern matching case

# Commit: docs(docs): update documentation

# Commit: perf(middleware): optimize controller routing

# Commit: chore(di): update ci configuration

# Commit: fix(di-container): resolve body parsing

# Commit: refactor(route): restructure documentation

# Commit: docs(http): update documentation

# Commit: feat(ioc): add request injection

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for singleton scope

# Commit: feat(controller): add singleton scope

# Commit: feat(router): add context propagation

# Commit: refactor(server): restructure code structure

# Commit: fix(provider): resolve body parsing

# Commit: refactor(core): restructure type safety

# Commit: perf(example): optimize controller routing

# Commit: chore(hello): update license

# Commit: perf(test): optimize controller routing

# Commit: chore(docs): update dependencies

# Commit: feat(test): implement context propagation

# Commit: refactor(docs): restructure test coverage

# Commit: test(middleware): add coverage for provider registration

# Commit: docs(di): update documentation

# Commit: fix(di-container): resolve routing conflict

# Commit: fix(route): resolve body parsing

# Commit: chore(http): update go mod

# Commit: feat(ioc): implement module setup

# Commit: perf(container): optimize singleton scope

# Commit: feat(module): add singleton scope

# Commit: feat(controller): add route matching

# Commit: refactor(router): improve memory usage

# Commit: chore(server): update test suite

# Commit: test(provider): add coverage for param extraction

# Commit: perf(core): optimize route matching

# Commit: refactor(example): improve type safety

# Commit: refactor(hello): restructure test coverage

# Commit: test(test): add coverage for route matching

# Commit: refactor(docs): improve test coverage

# Commit: test(middleware): add coverage for module setup

# Commit: docs(docs): update documentation

# Commit: chore(middleware): update gitignore

# Commit: chore(di): update readme

# Commit: refactor(di-container): restructure type safety

# Commit: perf(route): optimize middleware chain

# Commit: feat(http): implement context propagation

# Commit: refactor(ioc): restructure concurrency handling

# Commit: feat(container): implement response writing

# Commit: test(module): add coverage for controller routing

# Commit: chore(controller): update license

# Commit: perf(router): optimize response writing

# Commit: fix(server): resolve scope resolution

# Commit: refactor(provider): improve performance

# Commit: refactor(core): restructure concurrency handling

# Commit: feat(example): implement context propagation

# Commit: refactor(hello): restructure memory usage

# Commit: chore(test): update readme

# Commit: refactor(docs): restructure code structure

# Commit: fix(middleware): handle body parsing case

# Commit: refactor(di): restructure code structure

# Commit: chore(middleware): update dependencies

# Commit: feat(di): add request injection

# Commit: fix(di-container): resolve pattern matching

# Commit: docs(route): update documentation

# Commit: perf(http): optimize response writing

# Commit: test(ioc): add coverage for error handling

# Commit: fix(container): handle scope resolution case

# Commit: refactor(module): improve concurrency handling

# Commit: feat(controller): implement middleware chain

# Commit: feat(router): implement response writing

# Commit: feat(server): add route matching

# Commit: refactor(provider): improve memory usage

# Commit: fix(core): handle path extraction case

# Commit: feat(example): implement error handling

# Commit: fix(hello): handle type inference case

# Commit: fix(test): handle body parsing case

# Commit: chore(docs): update gitignore

# Commit: chore(middleware): update readme

# Commit: refactor(di): restructure type safety

# Commit: perf(di-container): optimize request injection

# Commit: test(di): add coverage for handler resolution

# Commit: test(di-container): add coverage for response writing

# Commit: fix(route): handle body parsing case

# Commit: fix(http): handle body parsing case

# Commit: chore(ioc): update build script

# Commit: refactor(container): improve code structure

# Commit: fix(module): resolve scope resolution

# Commit: refactor(controller): improve error messages

# Commit: refactor(router): improve test coverage

# Commit: test(server): add coverage for response writing

# Commit: refactor(provider): restructure concurrency handling

# Commit: fix(core): handle scope resolution case

# Commit: refactor(example): improve memory usage

# Commit: chore(hello): update gitignore

# Commit: chore(test): update test suite

# Commit: test(docs): add coverage for response writing

# Commit: refactor(middleware): restructure performance

# Commit: feat(di): add module setup

# Commit: perf(di-container): optimize middleware chain

# Commit: feat(route): implement context propagation

# Commit: feat(di-container): add provider registration

# Commit: docs(route): update documentation

# Commit: perf(http): optimize controller routing

# Commit: chore(ioc): update gitignore

# Commit: chore(container): update gitignore

# Commit: chore(module): update test suite

# Commit: test(controller): add coverage for provider registration

# Commit: docs(router): update documentation

# Commit: feat(server): add module setup

# Commit: perf(provider): optimize singleton scope

# Commit: feat(core): add module setup

# Commit: perf(example): optimize middleware chain

# Commit: feat(hello): implement module setup

# Commit: perf(test): optimize request injection

# Commit: fix(docs): resolve pattern matching

# Commit: refactor(middleware): restructure type safety

# Commit: fix(di): handle nil pointer case

# Commit: feat(di-container): add request injection

# Commit: fix(route): resolve pattern matching

# Commit: docs(http): update documentation

# Commit: feat(route): add response writing

# Commit: refactor(http): restructure type safety

# Commit: perf(ioc): optimize request injection

# Commit: fix(container): resolve type inference

# Commit: test(module): add coverage for error handling

# Commit: fix(controller): handle body parsing case

# Commit: chore(router): update test suite

# Commit: test(server): add coverage for handler resolution

# Commit: test(provider): add coverage for request injection

# Commit: fix(core): resolve pattern matching

# Commit: fix(example): handle pattern matching case

# Commit: docs(hello): update documentation

# Commit: perf(test): optimize handler resolution
