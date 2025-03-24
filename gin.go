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

func NewHttpServer(env env.Environment, config *HttpConfig.HttpServerConfig) *HttpServer {
	gin.DebugPrintRouteFunc = debugPrintRoutes
	gin.DebugPrintFunc = func(format string, values ...interface{}) {}

	r := gin.New()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}
	SetCors(r, strings.Split(config.AllowedOrigins, ","))

	gs := &HttpServer{
		Gin: r,

		config: config,

		Env: env,
	}

	gs.handleMiddleware()
	gs.handleErrorRoutes()
	r.Use(gs.LogRequest)

	return gs
}

func (gs *HttpServer) Run() error {
	return gs.Gin.Run(fmt.Sprintf("%s:%d", gs.config.Host, gs.config.Port))
}
