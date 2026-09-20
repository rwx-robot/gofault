# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v3.4.0] — 2026-09-20

### Added
- **API Versioning Middleware** (`versioning/` package)
  - `HeaderStrategy`: extract version from `Accept` header using pattern `application/vnd.app.v{version}+json`
  - `PathPrefixStrategy`: extract version from URL path prefix (e.g. `/v3` → version `3`)
  - `QueryStrategy`: extract version from query parameter (e.g. `api-version=2`)
  - `VersionSet`: lifecycle management for version collection
    - `StatusActive`: normal request handling
    - `StatusDeprecated`: auto-adds `Deprecation: true` and `Sunset: <date>` headers
    - `StatusEOL`: returns `410 Gone`
  - `VersionHandler`: request processor with automatic deprecation header injection
- **Core extensions**
  - `ctx.GetVersion()`: get current request version string
  - `ctx.GetVersionStatus()`: get current version status
  - `ctx.RespHeader(key, val)`: set response header (chainable)

### Tests
- 17 new tests covering all three strategies, VersionSet lifecycle, and EOL behavior

---

## [v3.3.0] — 2026-09-20

### Added
- **gRPC Server** (`grpc/` package)
  - `Server` type wrapping `grpc.Server` with `core.Module` lifecycle
  - `Config`: Port, Network, KeepAlive, MaxConcurrentStreams, TLS, interceptor configuration
  - `RecoveryInterceptor()`: panic recovery with stack trace logging
  - `LoggingInterceptor(*slog.Logger)`: structured logging of RPC calls
  - `UnaryInterceptor` / `StreamInterceptor`: helper wrappers for custom interceptors
  - `KeepAliveConfig()`: server keepalive parameters
  - `TLSCreds(certFile, keyFile)`: TLS credentials builder
  - `GRPCServerModule`: integrates gRPC lifecycle with gofault `OnBoot`/`OnShutdown`
  - `Start()`, `Stop()`, `Serve()` lifecycle methods

### Tests
- 24 new tests covering Server lifecycle, interceptor behavior, TLS, and error conditions

---

## [v3.2.0] — 2026-09-20

### Added
- **Static File Middleware** (`middleware/static.go`)
  - `Static(StaticConfig) core.MiddlewareFunc`: static file serving middleware
  - `StaticConfig`: Root (required), Prefix, IndexFile, EnableIndex, EnableSPA, ExtraExtensions, CacheDuration
  - Prefix stripping: `/static/css/app.css` → `./public/css/app.css`
  - Path traversal protection: absolute paths must have root directory prefix
  - Index file serving for directory requests
  - HTML directory listing generation
  - SPA mode: all 404s return IndexFile
  - `Cache-Control: max-age` header when CacheDuration > 0
  - `detectContentType()`: MIME type detection with custom extension override

### Tests
- 13 new tests covering file serving, 404, prefix stripping, index files, directory listing, path traversal, and MIME types

---

## [v3.1.0] — 2026-09-20

### Added
- **Redis Integration** (`redis/` package)
  - `Client`: Redis module (`core.Module`) wrapping `*redis.Client`
  - `Config`: Addr, Password, DB, PoolSize, MinIdleConns, DialTimeout, ReadTimeout, WriteTimeout, ConnMaxLifetime
  - `NewClient(name, config)`: creates client with connection test and auto-cleanup on shutdown
  - `MustNewClient(name, config)`: panics on error
  - `Cache(space string, ttl time.Duration)`: typed cache with key prefix and TTL
  - `Lock(key string, ttl time.Duration)`: distributed lock with owner validation
  - Pool connection management with configurable size and timeouts

### Tests
- 7 tests covering connection, cache operations, TTL expiration, lock acquire/release, key existence
- Requires Redis server on `localhost:9999`; tests skip when unavailable

---

## [v3.0.0] — 2026-09-20

### Added
- **GORM Database Integration** (`gorm/` package)
  - `Database` type implementing `core.Module` lifecycle
  - `Config`: Dialect, DSN, Silent, MaxOpenConns, MaxIdleConns, ConnMaxLifetime, ConnMaxIdleTime
  - `NewDatabase(name, cfg)`: creates database module with auto-cleanup on shutdown
  - `MustNewDatabase(name, cfg)`: panics on error
  - `Transaction(db, fn)`: context-aware transaction with automatic commit/rollback
  - Support for MySQL, PostgreSQL, SQLite dialects
  - Automatic connection pool configuration

### Tests
- 9 tests covering database creation, transaction commit/rollback, configuration, and error handling

---

## [v2.9.0] — 2026-09-20

### Added
- **File Upload Middleware** (`middleware/upload.go`)
  - `Upload(UploadConfig) core.MiddlewareFunc`: HTTP file upload handler
  - `UploadConfig`: Path, MaxSize, MaxFiles, AllowedTypes, StorageBackend interface
  - `StorageBackend` interface: `Save(formFile *multipart.FileHeader) (string, error)`
  - `LocalStorage`: disk-backed storage implementation
  - `MaxFormSize()`: get configured max form size
  - `UploadedFiles(ctx)`: retrieve uploaded files from request context
  - Multipart form parsing with size and count limits
  - File type validation (magic bytes + extension)
  - `StorageBackend` abstraction for future Redis/S3 backends

### Tests
- 14 tests covering single/multiple file upload, size limits, type validation, storage backend, form parsing errors

---

## [v2.8.0] — 2026-09-20

### Added
- **RequestID Middleware** (`middleware/requestid.go`)
  - `RequestID() core.MiddlewareFunc`: generates/propagates request ID
  - Header: `X-Request-ID`
  - UUID v4 generation with fallback to nano ID
  - Request context propagation via `core.Locals`

- **Session Middleware** (`middleware/session.go`)
  - `Session(sessionCfg) core.MiddlewareFunc`: cookie-based session management
  - `SessionConfig`: Name, MaxAge, HTTPOnly, Secure, SameSite, SessionData store
  - In-memory session store with TTL
  - Session fixation protection (regenerate on auth)
  - `GetSession(ctx)`: retrieve session data from context
  - `SaveSession(ctx)`: persist session changes

---

## [v2.7.0] — 2026-09-20

### Added
- **RateLimit Middleware** (`middleware/ratelimit.go`)
  - `RateLimit(RateLimitConfig) core.MiddlewareFunc`: token bucket rate limiting
  - `RateLimitConfig`: Limit (requests), Window (duration), KeyFunc
  - Client identification via IP or custom key function
  - Atomic counter with sliding window
  - `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` headers
  - `429 Too Many Requests` response when exceeded

---

## [v2.6.0] — 2026-09-20

### Added
- **JWT Middleware** (`middleware/jwt.go`)
  - `JWT(JWTConfig) core.MiddlewareFunc`: JWT authentication middleware
  - `JWTConfig`: KeyFunc (required), SigningMethod, Claims, SkipPaths
  - Custom claims extraction from `core.Locals`
  - Path-based skip list for public endpoints
  - Token expiration validation

---

## [v2.5.0] — 2026-09-20

### Added
- **CORS Middleware** (`middleware/cors.go`)
  - `CORS(CORSConfig) core.MiddlewareFunc`: Cross-Origin Resource Sharing
  - `CORSConfig`: AllowOrigins, AllowMethods, AllowHeaders, AllowCredentials, MaxAge, ExposeHeaders
  - Preflight request handling (`OPTIONS`)
  - Vary: Origin header management
  - `Access-Control-Allow-Origin` credential-aware handling

---

## [v2.4.0] — 2026-09-20

### Added
- **Cache Middleware** (`middleware/cache.go`)
  - `Cache(cacheCfg) core.MiddlewareFunc`: in-memory response caching
  - `CacheConfig`: TTL, KeyFunc, VaryBy
  - `X-Cache-Status: HIT/MISS` header
  - LRU eviction with max entries limit
  - Request method filtering (GET only)

---

## [v2.3.0] — 2026-09-20

### Added
- **Metrics Middleware** (`middleware/metrics.go`)
  - `Metrics(prometheus.Registry) core.MiddlewareFunc`: Prometheus metrics
  - Request count, latency histogram, active requests gauge
  - `http_requests_total`, `http_request_duration_seconds`, `http_requests_in_flight`
  - Custom labels: method, path, status

---

## [v2.2.0] — 2026-09-20

### Added
- **OpenAPI Middleware** (`middleware/openapi.go`)
  - `OpenAPI(openapiCfg) core.MiddlewareFunc`: OpenAPI/Swagger spec serving
  - `OpenAPIConfig`: SpecURL, UIURL, Theme
  - Built-in Swagger UI with `swagger/index.css/js`
  - `application/yaml` content type for spec files

---

## [v2.1.0] — 2026-09-20

### Added
- **WebSocket Middleware** (`middleware/websocket.go`)
  - `WebSocket(upgraderCfg) core.MiddlewareFunc`: WebSocket upgrade handling
  - `WebSocketConfig`: ReadBufferSize, WriteBufferSize, CheckOrigin
  - `GetWebSocket(ctx)`: retrieve `*websocket.Conn` from context
  - Connection lifecycle via context cancellation

---

## [v2.0.0] — 2026-09-20

### Added
- **Timeout Middleware** (`middleware/timeout.go`)
  - `Timeout(timeoutCfg) core.MiddlewareFunc`: request deadline enforcement
  - `TimeoutConfig`: Timeout duration, ErrorMessage
  - `context.WithTimeout` propagation
  - `core.ErrTimeout` on deadline exceeded

---

## [v1.9.0] — 2026-09-20

### Added
- **Request Logger Middleware** (`middleware/requestlogger.go`)
  - `RequestLogger(logCfg) core.MiddlewareFunc`: structured request logging
  - `RequestLoggerConfig`: Logger, LogBody, LogHeaders, SkipPaths
  - Method, path, status, latency, IP, user-agent logging
  - Body/headers logging (configurable, sensitive data redaction)

---

## [v1.8.0] — 2026-09-20

### Added
- **IP Filter Middleware** (`middleware/ipfilter.go`)
  - `IPFilter(ipFilterCfg) core.MiddlewareFunc`: IP allowlist/denylist
  - `IPFilterConfig`: AllowList, DenyList, BlockedHandler
  - CIDR notation support (e.g. `192.168.1.0/24`)
  - Custom blocked response handler

---

## [v1.7.0] — 2026-09-20

### Added
- **Compression Middleware** (`middleware/compression.go`)
  - `Compression(compCfg) core.MiddlewareFunc`: gzip/deflate response compression
  - `CompressionConfig`: Level (gzip level), MinSize (minimum body size)
  - `Accept-Encoding` negotiation
  - `Vary: Accept-Encoding` header

---

## [v1.6.0] — 2026-09-20

### Added
- **Recovery Middleware** (`middleware/recovery.go`)
  - `Recovery(recoveryCfg) core.MiddlewareFunc`: panic recovery
  - `RecoveryConfig`: LogLevel, StackTrace, ErrorMessage
  - Stack trace capture and logging
  - Custom error response format

---

## [v1.5.0] — 2026-09-20

### Added
- **Validator Middleware** (`middleware/validator.go`)
  - `Validator(validatorCfg) core.MiddlewareFunc`: request body validation
  - `ValidatorConfig`: TagName, BindObject, ErrorHandler
  - go-playground/validator v10 integration
  - Custom tag support and error messages

- **HealthCheck Middleware** (`middleware/healthcheck.go`)
  - `HealthCheck(healthCfg) core.MiddlewareFunc`: liveness/readiness probe
  - `HealthConfig`: Path, Checks, ResponseFormat
  - Custom health check functions
  - JSON/plaintext response format

---

## [v1.0.0] — 2026-09-20

### Added
- Complete web framework with NestJS-inspired architecture in Go idioms
- **Core**: `Ctx`, `Handler`, `MiddlewareFunc`, `Module`, `Provider`, `Controller`, `Route`
- **IoC Container**: constructor injection, singleton scope, request scope
- **Router**: HTTP method matching, parameter extraction (`/users/:id`)
- **Server**: graceful shutdown, port configuration
- **Controller**: prefix-based routing, route registration
- **Logger**: structured logging with `slog`
- **Exception Filter**: panic recovery, error response formatting

---

## [v0.5.0] — 2026-09-20

### Added
- Middleware chain with ordered execution
- Exception filter for error handling
- Lifecycle hooks: `OnBoot`, `OnShutdown`
- Handler dispatch with context propagation

---

## [v0.1.0] — 2026-09-20

### Added
- Initial project structure
- Basic HTTP server setup
- Context and handler types
