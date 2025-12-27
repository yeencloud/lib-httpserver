package httpserver

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain/config"
	"github.com/yeencloud/lib-shared/env"
)

type HttpServer struct {
	config *HttpConfig.HttpServerConfig

	Gin *gin.Engine
	Env env.Environment
}

func NewHttpServer(env env.Environment, config *HttpConfig.HttpServerConfig) (*HttpServer, error) {
	gin.DebugPrintRouteFunc = debugPrintRoutes
	gin.DebugPrintFunc = func(format string, values ...interface{}) {}

	r := gin.New()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	setCors(r, strings.Split(config.AllowedOrigins, ","))

	gs := &HttpServer{
		Gin: r,

		config: config,
		Env:    env,
	}

	gs.setupSession()
	gs.handleMiddleware()
	gs.handleErrorRoutes()
	r.Use(gs.logRequest)

	return gs, nil
}

func (hs *HttpServer) Run() error {
	listenAddr := fmt.Sprintf("%s:%d", hs.config.Host, hs.config.Port)
	if hs.config.TLS {
		return hs.Gin.RunTLS(listenAddr, hs.config.TLSCert, hs.config.TLSKey)
	}
	return hs.Gin.Run(listenAddr)
}
