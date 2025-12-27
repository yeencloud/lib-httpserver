package httpserver

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) getPath(ctx *gin.Context) string {
	path := ctx.FullPath()
	if path == "" {
		path = ctx.Request.RequestURI
	}

	return path
}

func (hs *HttpServer) profileNextRequest(ctx *gin.Context) time.Duration {
	start := time.Now()
	ctx.Next()
	end := time.Now()
	return end.Sub(start)
}
