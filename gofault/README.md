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

## v3.0.0 — 数据库集成 (规划中)

### 规划特性
- GORM 适配层
- 自动迁移
- 事务支持
- 仓储模式 (Repository Pattern)

---

## v3.1.0 — Redis 集成 (规划中)

### 规划特性
- Redis 连接池
- 分布式 Cache 后端
- Session 存储后端
- 分布式锁

---

## v3.2.0 — 静态文件服务 (规划中)

### 规划特性
- 静态文件中间件
- 目录索引
- 缓存控制头
- SPA 支持

---

## v3.3.0 — gRPC 支持 (规划中)

### 规划特性
- gRPC Server 生成
- Protobuf 适配
- HTTP/gRPC 双协议支持

---

## v3.4.0 — API Versioning (规划中)

### 规划特性
- URL 路径版本 (`/api/v1/`, `/api/v2/`)
- Header 版本 (`Accept: application/vnd.app.v1+json`)
- 版本化路由分组

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

## License

MIT
