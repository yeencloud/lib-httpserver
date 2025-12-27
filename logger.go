package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yeencloud/lib-httpserver/domain"
	shared "github.com/yeencloud/lib-shared/log"
)

func (hs *HttpServer) createLoggerForRequest(ctx *gin.Context) {
	logEntry := log.NewEntry(log.StandardLogger())

	logEntry = domain.HttpPathKey.WithValue(ctx.Request.RequestURI).AsField(logEntry)
	logEntry = domain.HttpMethodKey.WithValue(ctx.Request.Method).AsField(logEntry)

	setLogger(ctx, logEntry)
}

type HttpRequestMetric struct {
	Method   string `metric:"method"`
	Path     string `metric:"path"`
	Status   int    `metric:"status"`
	Duration int64  `metric:"duration"`
}

func (hs *HttpServer) logRequest(c *gin.Context) {
	path := hs.getPath(c)
	latency := hs.profileNextRequest(c)

	shared.GetLoggerFromContext(c).WithFields(log.Fields{
		domain.LogHttpResponseStatusCodeField.String(): c.Writer.Status(),
		domain.LogHttpMethodField.String():             c.Request.Method,
		domain.LogHttpPathField.String():               path,
		domain.LogHttpResponseTimeField.String():       latency.Milliseconds(),
	}).Log(hs.mapHttpStatusToLoggingLevel(c), fmt.Sprintf("%s %s %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status()))
}

func (hs *HttpServer) mapHttpStatusToLoggingLevel(ctx *gin.Context) log.Level {
	status := ctx.Writer.Status()

	level := log.InfoLevel

	if status >= 400 && status < 500 {
		level = log.WarnLevel
	} else if status >= 500 {
		level = log.ErrorLevel
	}

	return level
}
