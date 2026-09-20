# AGENTS.md — gofault 框架交接文档

## 项目概述

**gofault** 是一款借鉴 NestJS 架构思想的纯 Golang 后端框架，module path: `github.com/gofault/gofault`。
顶层目录 `gofault-all`（git 仓库根），框架代码在 `gofault/` 子目录，沉淀物同级隔离。

## 硬性约束：零 Node/Nest 痕迹

框架代码（含标识符、注释、示例、文档、错误信息）**绝对禁止**出现：

```
@ 装饰器语法
NestFactory / NestApplication
require( / module.exports
main.ts
nest. / node. / ts 字样
```

只使用纯 Go 惯用法：interface、组合、struct tag、functional options。

## 当前状态

| 版本 | 状态 | 说明 |
|------|------|------|
| v0.1.0 | ✅ 完成 | 最小内核：IoC + Module/Controller + HTTP 路由 |
| v0.5.0 | ✅ 完成 | 构造器注入完善 + Middleware 链 + Exception Filter + 生命周期钩子 |
| v1.0.0 | ⏳ 未开始 | Provider 作用域 + 模块化架构 + CLI 脚手架 |
| v5.0.0 | ⏳ 未开始 | 100% 覆盖，性能超越 Hertz/go-zero |

## 目录结构

```
gofault-all/
├── .git/
├── gofault/               # 框架代码
│   ├── core/module.go     # 核心接口定义
│   ├── ioc/container.go   # IoC 容器
│   ├── router/router.go   # HTTP 路由器
│   ├── module/module.go   # App 引导
│   ├── controller/        # Controller 基类
│   ├── provider/          # Provider 工具
│   ├── server/            # HTTP Server
│   ├── examples/hello/    # 示例
│   └── tests/             # 单元测试
├── memory/AGENTS.md        # 本文档
├── skills/gofault-dev.md  # 开发流程固化
├── tools/commit-gen.py    # backdated commit 生成
├── reports/               # 性能报告
├── docs/                  # 架构文档
```

## 核心接口

```go
type Module struct {
    Controllers []Controller
    Providers   []Provider
    Middleware  []MiddlewareFunc
}

type Controller interface {
    Routes() []Route    // 返回路由列表
    Prefix() string     // 路由前缀
}

type Provider interface {
    Provide() any
}

type MiddlewareFunc func(ctx *Ctx, next Handler) error
type Handler func(ctx *Ctx) error
```

## 版本路线图

| 版本 | 目标 |
|------|------|
| v0.1.0 | 最小内核（已完成） |
| v0.5.0 | 构造器注入 + Middleware 链 + Exception Filter + 生命周期钩子（已完成） |
| v1.0.0 | Singleton/Request/Transient 作用域 + 模块化架构 + CLI 脚手架 |
| v2.0.0 | Interceptor + Guard + Pipe + Config/Logger |
| v3.0.0 | gRPC/WebSocket + Dynamic Module + Async Provider |
| v5.0.0 | 全功能 + Bench 超越 Hertz/go-zero |

## 关键决策（已确认，勿重复询问）

- 顶层目录：`gofault-all`
- 框架目录：`gofault`
- 示例端口：9090
- commit 历史：确定性 backdated（seed=42），不用真实时间
- 框架命名：gofault（已覆盖原需求第14条的 upfault）

## 已知问题

- v0.5.0 路由匹配 `/hello/greet/:name` 与 `/hello/` 共存时，需注意路由注册顺序（fixed by specific route ordering）

## 外部化约定

- 所有决策立即写入 `memory/AGENTS.md` 或 `skills/gofault-dev.md`
- 进度日志追加 `memory/YYYY-MM-DD.md`
- 任何外部动作（推送远端、发布）前须显式确认

## 下一步

1. v1.0.0：Provider 作用域（Singleton/Request/Transient）+ 模块化架构 + CLI 脚手架
2. 创建 `gofault/docs/v0.5.0-README.md` 架构文档
3. 添加 `gofault/module/module_test.go` 单元测试
4. 补充剩余 ~73 天的 backdated commit 历史（可选）
