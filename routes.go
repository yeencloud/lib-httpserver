package httpserver

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/httpserver/domain"
)

func pageNotFoundError(method string, path string) error {
	page := domain.PageNotFoundError{}
	page.Method = method
	page.Path = path
	return &page
}

func (gs *HttpServer) recoverFromPanic(c *gin.Context, err interface{}) {
	gs.ReplyWithError(c, &domain.InternalServerError{}, err)
}

func (gs *HttpServer) handleErrorRoutes() {
	gs.Gin.NoRoute(func(ctx *gin.Context) {
		gs.ReplyWithError(ctx, pageNotFoundError(ctx.Request.Method, ctx.Request.URL.Path))
	})
}
