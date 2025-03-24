package httpserver

import (
	"fmt"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-httpserver/domain"
	"github.com/yeencloud/lib-httpserver/domain/error"
)

func debugPrintRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	entry := log.NewEntry(log.StandardLogger())

	entry = domain.LogHttpMethodField.WithValue(httpMethod).AsField(entry)
	entry = domain.LogHttpPathField.WithValue(absolutePath).AsField(entry)
	entry = domain.LogHttpHandlerCountField.WithValue(nuHandlers).AsField(entry)
	entry = domain.LogHttpHandlerNameField.WithValue(handlerName).AsField(entry)

	entry.Info(fmt.Sprintf("%s %s", httpMethod, absolutePath))
}

func (gs *HttpServer) recoverFromPanic(c *gin.Context, err interface{}) {
	gs.ReplyWithError(c, &HttpError.InternalServerError{
		AdditionalData: err,
	})
}

func (gs *HttpServer) handleErrorRoutes() {
	gs.Gin.NoRoute(func(ctx *gin.Context) {
		gs.ReplyWithError(ctx, &HttpError.PageNotFoundError{
			Method: ctx.Request.Method,
			Path:   ctx.Request.URL.Path,
		})
	})
}

func (gs *HttpServer) handleMiddleware() {
	r := gs.Gin

	gs.handleErrorRoutes()
	r.Use(nice.Recovery(gs.recoverFromPanic))
	r.Use(gs.CreateLoggerForRequest)
	r.Use(gs.CreateMetricsForRequest)
	r.Use(gs.handleHeader(domain.HeaderXRequestId, domain.HttpRequestIdKey))
	r.Use(gs.handleHeader(domain.HeaderXCorrelationId, domain.HttpCorrelationIdKey))
}
