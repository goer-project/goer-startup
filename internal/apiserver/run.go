package apiserver

import (
	"github.com/gin-gonic/gin"

	"goer-startup/internal/apiserver/config"
	"goer-startup/internal/pkg/known"
	"goer-startup/internal/pkg/middleware"
	"goer-startup/pkg/token"
)

// run 函数是实际的业务代码入口函数.
func run() error {
	// 初始化 store 层
	if err := InitStore(); err != nil {
		return err
	}

	// 初始化 cache
	if err := InitCache(); err != nil {
		return err
	}

	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	token.Init(config.Cfg.JWT.Key, known.XUsernameKey)

	// 设置 Gin 模式
	gin.SetMode(config.Cfg.Server.Mode)

	// 创建 Gin 引擎
	g := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.Cors, middleware.Secure, middleware.RequestID()}

	g.Use(mws...)

	// Swagger
	if config.Cfg.Feature.ApiDoc {
		MapSwagRoutes(g)
	}

	if err := installRouters(g); err != nil {
		return err
	}

	// 创建并运行 HTTP 服务器
	return startInsecureServer(g)
}
