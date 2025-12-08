package httpserver

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func (gs *HttpServer) MapHttpStatusToLoggingLevel(ctx *gin.Context) logrus.Level {
	status := ctx.Writer.Status()

	level := logrus.InfoLevel

	if status >= 400 && status < 500 {
		level = logrus.WarnLevel
	} else if status >= 500 {
		level = logrus.ErrorLevel
	}

	return level
}
