package httpserver

import (
	"context"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-shared/apperr"
	logShared "github.com/yeencloud/lib-shared/log"
)

func mapRequestContextToContext(requestHeader string, contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		ctx := req.Context()

		requestHeaderValue := c.GetHeader(requestHeader)

		if requestHeaderValue != "" {
			ctx = context.WithValue(ctx, contextKey, requestHeaderValue)
		}

		// Attach the new context back to the request
		c.Request = req.WithContext(ctx)

		c.Next()
	}
}

func GetLoggerFromGinContext(ctx *gin.Context) (*log.Entry, error) {
	entryFromContext, exists := ctx.Get(logShared.LoggerCtxKey)
	if !exists {
		return nil, &apperr.ObjectNotInContextError{Object: logShared.LoggerCtxKey}
	}

	entry, ok := entryFromContext.(*log.Entry)
	if !ok {
		return nil, &apperr.WrongObjectTypeInContextError{Object: logShared.LoggerCtxKey, ExpectedType: "*log.Entry"}
	}

	return entry, nil
}

func setLogger(ctx *gin.Context, logger *log.Entry) {
	ctx.Set(logShared.LoggerCtxKey, logger)
}
