package httpserver

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-shared/errors"
	logShared "github.com/yeencloud/lib-shared/log"
)

func GetLoggerFromGinContext(ctx *gin.Context) (*log.Entry, error) {
	entryFromContext := ctx.Value(logShared.LoggerCtxKey)
	if entryFromContext == nil {
		return nil, &errors.ObjectNotInContextError{Object: logShared.LoggerCtxKey}
	}

	entry, ok := entryFromContext.(*log.Entry)
	if !ok {
		return nil, &errors.WrongObjectTypeInContextError{Object: logShared.LoggerCtxKey, ExpectedType: "*log.Entry"}
	}

	return entry, nil
}

func setLogger(ctx *gin.Context, logger *log.Entry) {
	ctx.Set(logShared.LoggerCtxKey, logger)
}
