package httpserver

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-logger/domain"
)

func (gs *HttpServer) GetPath(ctx *gin.Context) string {
	path := ctx.FullPath()
	if path == "" {
		path = ctx.Request.RequestURI
	}

	return path
}

func (gs *HttpServer) ProfileNextRequest(ctx *gin.Context) time.Duration {
	start := time.Now()
	ctx.Next()
	end := time.Now()
	return end.Sub(start)
}

func (gs *HttpServer) MapHttpStatusToLoggingLevel(ctx *gin.Context) LoggerDomain.Level {
	status := ctx.Writer.Status()
	level := LoggerDomain.LogLevelInfo
	if status >= 400 && status < 500 {
		level = LoggerDomain.LogLevelWarn
	} else if status >= 500 {
		level = LoggerDomain.LogLevelError
	}
	return level
}
