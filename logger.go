package httpserver

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-httpserver/domain"
	shared "github.com/yeencloud/lib-shared/log"
)

func (hs *HttpServer) CreateLoggerForRequest(ctx *gin.Context) {
	logEntry := log.NewEntry(log.StandardLogger())

	logEntry = domain.HttpPathKey.WithValue(ctx.Request.RequestURI).AsField(logEntry)
	logEntry = domain.HttpMethodKey.WithValue(ctx.Request.Method).AsField(logEntry)

	ctx.Set(shared.ContextLoggerKey, logEntry)
}

func GetLoggerFromContext(ctx *gin.Context) *log.Entry {
	logger, exists := ctx.Get(shared.ContextLoggerKey)
	if !exists {
		return log.NewEntry(log.StandardLogger())
	}
	entry, ok := logger.(*log.Entry)
	if !ok {
		return log.NewEntry(log.StandardLogger())
	}
	return entry
}
