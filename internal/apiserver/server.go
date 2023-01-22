package apiserver

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"goer-startup/internal/pkg/log"
)

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpSrv := &http.Server{Addr: viper.GetString("server.addr"), Handler: g}

	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("server.addr"))
	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpSrv
}

// startSecureServer 创建并运行 HTTPS 服务器.
func startSecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTPS Server 实例
	httpsSrv := &http.Server{Addr: viper.GetString("tls.addr"), Handler: g}

	// 运行 HTTPS 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTPS 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpsSrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}

	return httpsSrv
}
