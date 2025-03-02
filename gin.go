package httpserver

import (
	"fmt"
	"strings"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-shared/log"

	"github.com/yeencloud/lib-logger"
	"github.com/yeencloud/lib-logger/domain"

	"github.com/yeencloud/lib-httpserver/domain"
)

type HttpServer struct {
	config *domain.HttpServerConfig

	Gin *gin.Engine
}

type ContextKey struct {
	name string
}

func (ck ContextKey) String() string {
	return "context." + ck.name
}

func debugPrintRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	Logger.Log(LogDomain.LogLevelInfo).WithFields(log.Fields{
		domain.LogHttpMethodField:       httpMethod,
		domain.LogHttpPathField:         absolutePath,
		domain.LogHttpHandlerCountField: nuHandlers,
		domain.LogHttpHandlerNameField:  handlerName,
	}).Msg(fmt.Sprintf("%s %s", httpMethod, absolutePath))
}

func NewHttpServer(config *domain.HttpServerConfig) *HttpServer {
	gin.DebugPrintRouteFunc = debugPrintRoutes
	gin.DebugPrintFunc = func(format string, values ...interface{}) {}

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("raw_response", true)
	})

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}
	SetCors(r, strings.Split(config.AllowedOrigins, ","))

	gs := &HttpServer{
		Gin: r,

		config: config,
	}

	gs.handleErrorRoutes()
	r.Use(nice.Recovery(gs.recoverFromPanic))
	r.Use(CreateSharedRequest)
	r.Use(gs.handleRequestID)
	r.Use(gs.handleCorrelationID)

	return gs
}

func (gs *HttpServer) Run() error {
	return gs.Gin.Run(fmt.Sprintf("%s:%d", gs.config.Host, gs.config.Port))
}
