# gofault

A Go backend framework inspired by NestJS architecture, implemented in pure Go idioms.

## 核心接口

```go
// Module - 应用基本单元
type Module interface {
    Name() string
    Controllers() []Controller
    Providers() []Provider
    Middleware() []MiddlewareFunc
}

// Controller - 暴露 HTTP 端点
type Controller interface {
    Prefix() string
    Routes() []Route
}

// Provider - 依赖注入供应者
type Provider interface {
    Provide() any
}

// MiddlewareFunc - 中间件函数
type MiddlewareFunc func(ctx *Ctx, next Handler) error

// Handler - 路由处理函数
type Handler func(ctx *Ctx) error
```

## 架构

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

---

## v0.1.0 — 核心架构 (2019)

首个版本，实现框架核心架构。

### 特性
- IoC 容器（singleton 作用域）
- 基于 Module 的应用组织
- HTTP 路由与路径参数提取
- Controller + Provider 分层
- 最小化中间件链

### 核心接口
```go
type Module interface {
    Name() string
    Controllers() []Controller
    Providers() []Provider
}

type Controller interface {
    Prefix() string
    Routes() []Route
}

type Provider interface {
    Provide() any
}

type Route struct {
    Method string
    Path   string
    Handler Handler
}
```

### 使用示例
```go
type GreetingModule struct{}

func (m *GreetingModule) Name() string { return "greeting" }
func (m *GreetingModule) Controllers() []Controller { return []Controller{&GreetingController{}} }
func (m *GreetingModule) Providers() []Provider { return nil }

type GreetingController struct{}

func (c *GreetingController) Prefix() string { return "/hello" }
func (c *GreetingController) Routes() []Route {
    return []Route{
        {Method: "GET", Path: "/greet/:name", Handler: c.Greet},
        {Method: "GET", Path: "/", Handler: c.Hello},
    }
}
```

---

## v0.5.0 — 中间件链与生命周期

### 新增特性
- 中间件链 (`MiddlewareFunc`)
- 异常过滤器 (`ExceptionFilter`)
- 生命周期钩子 (OnStart, OnShutdown)
- Handler 调度

### 架构
```
Request → MiddlewareChain → Controller → Provider → Response
                    ↓
            ExceptionFilter
```

### 核心接口
```go
type MiddlewareFunc func(ctx *Ctx, next Handler) error

type ExceptionFilter interface {
    Catch(ctx *Ctx, err error)
}

type LifecycleHook interface {
    OnStart()
    OnShutdown()
}
```

---

## v1.0.0 — 完整框架

### 新增特性
- 完整中间件体系
- IoC 容器增强（依赖注入）
- Router 增强（路由冲突检测）
- 模块化应用构建

---

## v1.5.0 — Validator 与 HealthCheck

### 新增中间件

#### ValidatorMiddleware
请求体验证中间件，基于 `go-playground/validator`。

```go
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=0,lte=150"`
}

cfg := validator.DefaultValidatorConfig()
cfg.ValidateOnMount = true

app.Use(validator.ValidatorMiddleware(cfg))
```

#### HealthCheckMiddleware
健康检查端点。

```go
cfg := healthcheck.DefaultConfig()
cfg.Path = "/health"

app.Use(healthcheck.HealthCheckMiddleware(cfg))
// GET /health → 200 OK
```

---

## v1.6.0 — Recovery 与 Compression

### 新增中间件

#### RecoveryMiddleware
panic 恢复，防止服务崩溃。

```go
cfg := recovery.DefaultConfig()
cfg.LogStack = true

app.Use(recovery.RecoveryMiddleware(cfg))
```

#### CompressionMiddleware
HTTP 响应压缩（gzip/deflate）。

```go
cfg := compression.DefaultConfig()
cfg.Level = gzip.BestSpeed

app.Use(compression.CompressionMiddleware(cfg))
```

---

## v1.7.0 — Timeout

### 新增中间件

#### TimeoutMiddleware
请求超时控制。

```go
cfg := timeout.DefaultConfig()
cfg.Duration = 30 * time.Second

app.Use(timeout.TimeoutMiddleware(cfg))
// 超时返回 504 Gateway Timeout
```

---

## v1.8.0 — IPFilter

### 新增中间件

#### IPFilterMiddleware
IP 黑名单/白名单过滤。

```go
cfg := ipfilter.DefaultConfig()
cfg.BlockedIPs = []string{"192.168.1.100", "10.0.0.0/8"}

app.Use(ipfilter.IPFilterMiddleware(cfg))
// 阻止的 IP 返回 403 Forbidden
```

---

## v1.9.0 — RequestLogger

### 新增中间件

#### RequestLoggerMiddleware
结构化请求日志。

```go
cfg := logger.DefaultConfig()
cfg.LogBody = true
cfg.LogHeaders = []string{"X-Request-ID"}

app.Use(logger.RequestLoggerMiddleware(cfg))
// 输出: GET /api/users 200 12ms
```

---

## v2.0.0 — OpenAPI

### 新增中间件

#### OpenAPIMiddleware
自动生成 OpenAPI/Swagger 文档。

```go
cfg := openapi.DefaultConfig()
cfg.Title = "My API"
cfg.Version = "1.0.0"
cfg.Path = "/docs"

app.Use(openapi.OpenAPIMiddleware(cfg))
// GET /docs → Swagger UI
```

---

## v2.1.0 — WebSocket

### 新增中间件

#### WebSocketMiddleware / WebSocketHijackMiddleware

支持 WebSocket 全生命周期回调或轻量级 hijack 模式。

**全生命周期模式：**
```go
cfg := websocket.DefaultConfig()
cfg.Path = "/ws"

app.Use(websocket.WebSocketMiddleware(cfg, &websocket.WebSocketHandler{
    OnConnect: func(ctx *core.Ctx, conn *websocket.Conn) error {
        fmt.Println("client connected")
        return nil
    },
    OnMessage: func(ctx *core.Ctx, conn *websocket.Conn, msg []byte) {
        conn.WriteMessage(msg)
    },
    OnDisconnect: func(ctx *core.Ctx, conn *websocket.Conn) {
        fmt.Println("client disconnected")
    },
}))
```

**轻量级 hijack 模式：**
```go
app.Use(websocket.WebSocketHijackMiddleware("/ws"))

app.Get("/ws", func(ctx *core.Ctx) error {
    conn := ctx.Locals["ws_conn"].(*websocket.Conn)
    // 手动处理 WebSocket 消息
    return nil
})
```

### 核心更新
- `Ctx.Locals map[string]any` — 中间件间数据传递

---

## v2.2.0 — Metrics/Prometheus

### 新增中间件

#### MetricsMiddleware
Prometheus 指标收集，使用独立 Registry 避免全局冲突。

```go
cfg := metrics.DefaultMetricsConfig()
cfg.Namespace = "myapp"
cfg.Subsystem = "http"
cfg.SkipHealthCheck = true

mw, registry := metrics.MetricsMiddleware(cfg)
app.Use(mw)

// 使用 registry 暴露 /metrics 端点
http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
```

**指标：**
- `myapp_http_requests_total{method, path, status}` — 请求计数器
- `myapp_http_request_duration_seconds{method, path, status}` — 请求延迟直方图
- `myapp_http_requests_in_flight{method, path, status}` — 当前处理中请求数

---

## v2.3.0 — Cache

### 新增中间件

#### InMemoryCache + CacheMiddleware
内存缓存，支持 TTL 和 LRU 淘汰。

```go
cache := middleware.NewInMemoryCache(middleware.CacheConfig{
    TTL:     5 * time.Minute,
    MaxSize: 1000,
})

app.Use(middleware.CacheMiddleware(cache, middleware.DefaultCacheConfig()))

// 响应头:
// X-Cache: HIT  (命中)
// X-Cache: MISS (未命中)
```

**特性：**
- GET 请求自动缓存
- POST/PUT/DELETE 默认跳过
- 4xx/5xx 响应不缓存
- 支持自定义 SkipFunc

---

## v2.4.0 — CORS

### 新增中间件

#### CORSMiddleware
跨域资源共享支持。

```go
cfg := cors.DefaultConfig()
cfg.AllowOrigins = []string{"http://localhost:3000"}
cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
cfg.AllowHeaders = []string{"Authorization", "Content-Type"}

app.Use(cors.CORSMiddleware(cfg))
```

---

## v2.5.0 — JWT

### 新增中间件

#### JWTMiddleware
JWT 认证与声明提取。

```go
cfg := jwt.DefaultConfig()
cfg.Secret = "your-secret-key"
cfg.Path = "/api"
cfg.ExtractFrom = "Authorization Bearer"

app.Use(jwt.JWTMiddleware(cfg, func(ctx *core.Ctx, claims jwt.MapClaims) error {
    ctx.Locals["user_id"] = claims["sub"]
    return nil
}))
```

---

## v2.6.0 — RateLimit

### 新增中间件

#### RateLimitMiddleware
基于令牌桶算法的限流。

```go
cfg := ratelimit.DefaultConfig()
cfg.Max = 100
cfg.Window = time.Minute

app.Use(ratelimit.RateLimitMiddleware(cfg))
// 超限返回 429 Too Many Requests
```

---

## v2.7.0 — RequestID

### 新增中间件

#### RequestIDMiddleware
为每个请求生成唯一 ID，便于链路追踪。

```go
cfg := requestid.DefaultConfig()
cfg.HeaderName = "X-Request-ID"

app.Use(requestid.RequestIDMiddleware(cfg))
// Response Header: X-Request-ID: abc-123-def
```

---

## v2.8.0 — Session

### 新增中间件

#### SessionMiddleware
会话管理，支持内存存储。

```go
cfg := session.DefaultConfig()
cfg.SessionIDLength = 32

app.Use(session.SessionMiddleware(cfg))

app.Get("/profile", func(ctx *core.Ctx) error {
    sess := session.FromCtx(ctx)
    name := sess.GetString("name")
    return ctx.JSON(200, map[string]any{"name": name})
})
```

---

## v2.9.0 — File Upload

### 新增中间件

#### UploadMiddleware
文件上传处理，支持本地存储后端。

```go
storage := middleware.NewLocalStorage("./uploads")

cfg := middleware.DefaultUploadConfig()
cfg.Storage = storage
cfg.MaxSize = 5 * 1024 * 1024 // 5MB
cfg.AllowedExtensions = []string{".jpg", ".png", ".pdf"}

app.Use(middleware.UploadMiddleware(cfg))

app.Post("/upload", func(ctx *core.Ctx) error {
    files := middleware.GetUploadFiles(ctx)
    for _, f := range files {
        fmt.Println(f.FileName, f.Size, f.StoredPath)
    }
    return ctx.JSON(200, map[string]any{"files": files})
})
```

### 核心类型

| 类型 | 说明 |
|------|------|
| `StorageBackend` | 存储后端接口（可扩展 S3/Redis） |
| `LocalStorage` | 本地文件系统存储实现 |
| `FileInfo` | 上传文件元信息（文件名、大小、Content-Type、存储路径） |

### 规划特性
- [x] Multipart 文件上传处理
- [x] 文件大小限制
- [x] 允许/禁止文件类型白名单（按扩展名和 MIME type）
- [x] 存储后端抽象（`StorageBackend` 接口）
- [ ] S3 存储后端
- [ ] Redis 存储后端

---

## v3.0.0 — 数据库集成

### 新增包

#### `gorm/` — GORM 数据库集成

```go
cfg := gorm.DefaultConfig()
cfg.Dialect = gorm.DialectSQLite
cfg.DSN = ":memory:"
cfg.Silent = true

db, err := gorm.NewDatabase("main", cfg)

// 自动迁移
db.DB.AutoMigrate(&User{})

// 事务
gorm.Transaction(db.DB, func(tx *gorm.DB) error {
    return tx.Create(&User{Name: "alice"}).Error
})
```

### 核心类型

| 类型 | 说明 |
|------|------|
| `Database` | 数据库模块（`core.Module`），封装 `*gorm.DB` |
| `Config` | 数据库配置（连接池、超时、日志级别） |
| `Transaction()` | 事务辅助函数，自动 commit/rollback |

### 支持驱动

| Dialect | DSN 示例 |
|---------|---------|
| `mysql` | `user:password@tcp(localhost:3306)/dbname` |
| `postgres` | `host=localhost user=gorm password=gorm dbname=gorm port=5432` |
| `sqlite` | `./data.db` 或 `:memory:` |

### 规划特性
- [x] GORM 适配层
- [x] 连接池配置（MaxOpenConns, MaxIdleConns, ConnMaxLifetime）
- [x] 事务支持（自动 rollback on panic/error）
- [ ] 仓储模式 (Repository Pattern)

---

## v3.1.0 — Redis 集成

### 新增包

#### `redis/` — Redis 集成

```go
cfg := redis.DefaultConfig()
cfg.Addr = "localhost:6379"
cfg.PoolSize = 100

client := redis.MustNewClient("redis", cfg)

// Cache 操作
cache := client.Cache("myapp:", 10*time.Minute)
cache.Set(ctx, "key", "value", 5*time.Minute)
val, _ := cache.Get(ctx, "key")
cache.Del(ctx, "key")

// 分布式锁
lock := client.Lock("distributed-lock", 30*time.Second)
ok, _ := lock.TryAcquire(ctx)
defer lock.Release(ctx)
```

### 核心类型

| 类型 | 说明 |
|------|------|
| `Client` | Redis 模块（`core.Module`），封装 `*redis.Client` |
| `Config` | 连接池、超时、密码配置 |
| `Cache` | 带 TTL 和前缀的缓存抽象 |
| `Lock` | 分布式锁（SET NX + owner 校验） |

### 配置字段

| 字段 | 默认值 | 说明 |
|------|--------|------|
| `Addr` | `localhost:6379` | Redis 地址 |
| `Password` | `""` | AUTH 密码 |
| `DB` | `0` | 数据库编号 |
| `PoolSize` | `100` | 最大连接数 |
| `MinIdleConns` | `10` | 最小空闲连接 |
| `DialTimeout` | `5` | 连接超时（秒） |
| `ReadTimeout` | `3` | 读超时（秒） |
| `WriteTimeout` | `3` | 写超时（秒） |

---

## v3.2.0 — 静态文件服务

### 新增中间件

#### `middleware/static.go` — Static

```go
staticCfg := middleware.StaticConfig{
    Root:      "./public",
    Prefix:    "/static",
    IndexFile: "index.html",
    EnableIndex:  true,
    EnableSPA:    false,
    CacheDuration: 1 * time.Hour,
}

app.UseMiddleware(middleware.Static(staticCfg))
```

### 配置字段

| 字段 | 默认值 | 说明 |
|------|--------|------|
| `Root` | 必填 | 静态文件根目录（绝对路径） |
| `Prefix` | `"/"` | URL 路径前缀 |
| `IndexFile` | `"index.html"` | 目录索引文件 |
| `EnableIndex` | `false` | 启用目录列表 HTML |
| `EnableSPA` | `false` | 所有 404 指向 IndexFile |
| `ExtraExtensions` | `nil` | 自定义扩展名 → MIME 映射 |
| `CacheDuration` | `0` | 缓存 max-age（0=不设置） |

### 功能

- **前缀剥离**：请求 `/static/css/app.css` → 文件系统 `./public/css/app.css`
- **路径遍历保护**：绝对路径必须以根目录绝对路径为前缀
- **索引文件**：`/static/dir/` → 自动查找 `./public/dir/index.html`
- **目录列表**：启用后对目录请求生成 HTML 索引
- **SPA 支持**：所有未匹配请求返回 `IndexFile`
- **缓存头**：`CacheDuration > 0` 时设置 `Cache-Control: max-age`

---

## v3.3.0 — gRPC 支持

### 新增包

#### `grpc/` — gRPC 服务端集成

```go
cfg := grpc.DefaultConfig()
cfg.Port = 9000
cfg.Network = "tcp"
cfg.MaxConcurrentStreams = 100
cfg.UnaryInterceptors = []grpc.UnaryServerInterceptor{
    grpc.RecoveryInterceptor(),
    grpc.LoggingInterceptor(log.Default()),
}
cfg.Insecure = false
cfg.TLSCert = "cert.pem"
cfg.TLSKey = "key.pem"

server := grpc.NewServer(cfg)
server.RegisterService(func(srv *grpc.Server) {
    pb.RegisterEchoServiceServer(srv, &echoHandler{})
})
server.Start()
defer server.Stop()
```

### 核心类型

| 类型 | 说明 |
|------|------|
| `Server` | gRPC 服务器模块（`core.Module`） |
| `Config` | 端口、网络、KeepAlive、TLS、拦截器配置 |

### 配置字段

| 字段 | 默认值 | 说明 |
|------|--------|------|
| `Port` | `9000` | 监听端口 |
| `Network` | `"tcp"` | 网络类型 |
| `KeepAlive` | `nil` | KeepAlive 参数 |
| `MaxConcurrentStreams` | `100` | 最大并发流数量 |
| `Insecure` | `false` | 跳过 TLS（仅开发环境） |
| `TLSCert` | `""` | TLS 证书文件路径 |
| `TLSKey` | `""` | TLS 私钥文件路径 |
| `UnaryInterceptors` | `[]` | Unary 拦截器链 |
| `StreamInterceptors` | `[]` | Stream 拦截器链 |

### 内置拦截器

| 函数 | 说明 |
|------|------|
| `RecoveryInterceptor()` | 恐慌恢复，记录栈追踪 |
| `LoggingInterceptor(*slog.Logger)` | 结构化日志记录 RPC 调用 |
| `UnaryInterceptor(func(ctx, req, info, handler) error)` | Unary RPC 拦截器封装 |
| `StreamInterceptor(func(srv, ss, info, handler) error)` | Stream RPC 拦截器封装 |

### 生命周期

- `Start()` → `OnBoot`：启动监听
- `Stop()` / `graceful Stop` → `OnShutdown`：优雅停止

---

## v3.4.0 — API Versioning

### 新增包

#### `versioning/` — API 版本控制

### 三种版本提取策略

**Header 策略**（通过 `Accept` 头）：
```go
vm := versioning.NewVersioningMiddleware(versioning.HeaderStrategy(
    "Accept",
    "application/vnd.app.v{version}+json",
))
// Accept: application/vnd.app.v2+json → Locals["version"] = "2"
```

**PathPrefix 策略**（从 URL 路径前缀提取）：
```go
vm := versioning.NewVersioningMiddleware(versioning.PathPrefixStrategy("/v3"))
// 请求 /v3/users → 提取版本 "3"，路径剥离前缀后为 /users
```

**Query 策略**（通过 URL 查询参数）：
```go
vm := versioning.NewVersioningMiddleware(versioning.QueryStrategy("api-version"))
// GET /users?api-version=2 → Locals["version"] = "2"
```

### 版本集生命周期

```go
vs := versioning.NewVersionSet(map[string]versioning.VersionConfig{
    "1": {Status: versioning.StatusActive},      // 当前活跃
    "2": {Status: versioning.StatusDeprecated},  // 已废弃（自动添加 Deprecation 头）
    "3": {Status: versioning.StatusEOL},         // 已停止服务（返回 410 Gone）
})
```

### VersionHandler 自动行为

- `StatusActive`：正常处理
- `StatusDeprecated`：自动添加响应头 `Deprecation: true` + `Sunset: <date>`
- `StatusEOL`：中断请求，返回 `410 Gone`

### 核心类型

| 类型 | 说明 |
|------|------|
| `VersionSet` | 版本集合，管理所有版本的配置和生命周期 |
| `VersionConfig` | 单个版本配置（状态、Sunset 日期、custom headers） |
| `VersionStatus` | 枚举：`StatusActive` / `StatusDeprecated` / `StatusEOL` |
| `Middleware` | 版本控制中间件（`core.MiddlewareFunc`） |
| `VersionHandler` | 版本化请求处理器，自动注入版本信息 |

### Core 扩展

| 方法 | 说明 |
|------|------|
| `ctx.GetVersion()` | 获取当前请求版本号字符串 |
| `ctx.GetVersionStatus()` | 获取当前版本状态（Active/Deprecated/EOL） |
| `ctx.RespHeader(key, val)` | 设置响应头（链式调用） |

---

## 安装

```bash
go get github.com/gofault/gofault
```

## 快速开始

```go
package main

import (
    "github.com/gofault/gofault/server"
    "github.com/gofault/gofault/middleware"
)

func main() {
    app := server.New()

    // 全局中间件
    app.Use(middleware.RecoveryMiddleware(middleware.DefaultRecoveryConfig()))
    app.Use(middleware.LoggerMiddleware(middleware.DefaultLoggerConfig()))

    // 注册模块
    app.RegisterModule(&HelloModule{})

    app.Run()
}

type HelloModule struct{}

func (m *HelloModule) Name() string                     { return "hello" }
func (m *HelloModule) Controllers() []server.Controller { return []server.Controller{&HelloController{}} }
func (m *HelloModule) Providers() []server.Provider    { return nil }

type HelloController struct{}

func (c *HelloController) Prefix() string {
    return "/hello"
}

func (c *HelloController) Routes() []server.Route {
    return []server.Route{
        {Method: "GET", Path: "/", Handler: c.Hello},
    }
}

func (c *HelloController) Hello(ctx *server.Ctx) error {
    return ctx.JSON(200, map[string]string{"message": "Hello, World!"})
}
```

## 项目结构

```
gofault/
├── core/          # 核心接口与上下文
├── server/        # HTTP 服务器
├── router/        # 路由匹配
├── controller/    # 控制器
├── middleware/    # 中间件 (18个)
├── ioc/           # 依赖注入容器
├── module/        # 模块系统
├── provider/      # 供应者接口
├── logger/        # 日志
├── exception/     # 异常处理
├── config/        # 配置管理
└── examples/      # 示例
```

## 测试

```bash
go test ./...
```

## Roadmap

### v3.5.0 — 近期计划

| 特性 | 说明 | 优先级 |
|------|------|--------|
| **Repository 模式** | GORM 上的仓储抽象层，分离 DB 访问逻辑 | 高 |
| **Redis 限流后端** | RateLimit 中间件持久化到 Redis，支持分布式 | 高 |
| **OAuth2 / SSO 中间件** | JWT 之外的 OAuth2、JWK 校验、SSO 集成 | 中 |
| **OpenTelemetry 追踪** | 链路追踪、span 注入，对接 Jaeger/Zipkin | 中 |

### v4.0.0 — 远期演进

| 特性 | 说明 |
|------|------|
| **GraphQL 支持** | `graphql/` 包，schema-first 或 code-first |
| **Event / CQRS** | 事件驱动架构支持 |
| **消息队列集成** | RabbitMQ / Kafka 后端集成 |
| **自动 OpenAPI 生成** | 从 Controller routes 自动生成 OpenAPI 3.0 spec |
| **热重载 / Devtools** | 开发阶段模板/路由热更新 |

### 待改进项

| 问题 | 说明 |
|------|------|
| Redis 测试 | 用 miniredis 替代真实 Redis，CI 可跑 |
| Examples 目录 | 当前只有 Quick Start 代码片段 |
| CI 覆盖率报告 | 建议加上 `go test -cover` + codecov |

---

## License

MIT
