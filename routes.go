package httpserver

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain"
)

func pageNotFoundError(method string, path string) error {
	page := domain.PageNotFoundError{}
	page.Method = method
	page.Path = path
	return &page
}

func (gs *HttpServer) recoverFromPanic(c *gin.Context, err interface{}) {
	gs.ReplyWithError(c, &domain.InternalServerError{
		AdditionalData: err,
	})
}

func (gs *HttpServer) handleErrorRoutes() {
	gs.Gin.NoRoute(func(ctx *gin.Context) {
		gs.ReplyWithError(ctx, pageNotFoundError(ctx.Request.Method, ctx.Request.URL.Path))
	})
}

func (gs *HttpServer) handleMiddleware() {
	r := gs.Gin

	gs.handleErrorRoutes()
	r.Use(nice.Recovery(gs.recoverFromPanic))
	r.Use(gs.CreateLoggerForRequest)
	r.Use(gs.CreateMetricsForRequest)
	r.Use(gs.handleRequestID)
	r.Use(gs.handleCorrelationID)
}
