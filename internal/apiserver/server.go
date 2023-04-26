package apiserver

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func startInsecureServer(g *gin.Engine) error {
	s := endless.NewServer(viper.GetString("server.addr"), g)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	err := s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
