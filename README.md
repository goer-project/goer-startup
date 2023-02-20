# Go project startup

开箱即用的 Go 项目开发脚手架

# Usage

## 复制配置文件

```bash
mkdir ~/.goer
cp configs/*.yaml ~/.goer
```

## 生成 CA 证书

```bash
make ca
```

## 构建

> 构建产物路径 ./_output

代码检测及构建

```bash
make
```

仅构建

```bash
make build
```

运行

```bash
# mac
./_output/platforms/darwin/arm64/goer-apiserver

# linux
./_output/platforms/linux/amd64/goer-apiserver
```

## 热重载

> 开发阶段可使用 air 插件热重载


安装 air 插件

```bash
go install github.com/cosmtrek/air@latest
```

复制 air 配置文件

```bash
cp .air.example.toml .air.toml
```

热重载

```bash
air
```

## 健康检查

```bash
curl localhost:8080/healthz
```

输出一下内容代表成功运行

```json
{
  "status": "ok"
}
```

# 开发规范

## 入口文件

```
cmd
├── goer-apiserver
├── goer-watcher
└── goerctl
```

- goer-apiserver: Http/Https 服务
- goer-watcher: 定时任务，分布式锁，同一任务确保同一时间只有一台机器执行
- goerctl: 命令行工具

## 目录组织

```
├── internal
│   ├── apiserver
|   ├── goerctl
|   ├── pkg
|   └── watcher
├── pkg
```

internal: 开发阶段业务逻辑均在 ```internal``` 文件夹下，各模块内容与 ```cmd``` 入口文件对应

- apiserver: 核心业务逻辑位置，其它模块可直接引用该模块业务逻辑代码
- goerctl: 命令行工具业务逻辑
- watcher: 计划任务业务逻辑
- pkg: internal 下公共代码

pkg: 项目公共代码，不涉及项目内部逻辑，可对外部项目公开

## 核心业务逻辑

核心业务逻辑均在 ```internal/apiserver``` 下

```
internal
├── apiserver
│   ├── biz
│   ├── controller
│   ├── store
└── pkg
    └── model
```

### 4层架构

- controller: 处理 Http 请求、表单验证，依赖 Biz 层、Model 层
- biz: 核心业务逻辑，依赖 store 层、Model 层
- store: 数据存储层，与 DB 交互，依赖数据库、Model 层
- model: 数据库表字段的 Go 结构体映射，禁止与 DB 交互，放在 pkg 下是因为方便项目后期公开 sdk 等

开发顺序：
先开发依赖少的组件：
Model 层 -> Store 层 -> Biz 层 -> Controller 层

### 错误码
```
internal
└── pkg
    └── errno
```
示例：
```go
var (
    // OK 代表请求成功.
    OK = &Errno{HTTP: 200, Code: "", Message: ""}
    
    // InternalServerError 表示所有未知的服务器端错误.
    InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal server error."}
)

```

### 路由与中间件
```
internal
├── apiserver
│   ├── router.go
└── pkg
    └── middleware
```
默认中间件可参考 ```internal/apiserver/run.go```，自定义中间件在 router 中使用：

```go
// v1 group
v1 := g.Group("/v1")
v1.Use(middleware.Authn(), middleware.Authz(authz))
```
