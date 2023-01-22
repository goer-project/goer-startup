package apiserver

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"goer-startup/internal/pkg/core"
	"goer-startup/internal/pkg/errno"
	"goer-startup/internal/pkg/log"
)

func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// 注册 pprof 路由
	pprof.Register(g)

	return nil
}
