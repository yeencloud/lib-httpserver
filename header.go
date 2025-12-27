package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/yeencloud/lib-shared/apperr"

	"github.com/yeencloud/lib-httpserver/domain"
	sharedLog "github.com/yeencloud/lib-shared/log"
	"github.com/yeencloud/lib-shared/namespace"
)

func (hs *HttpServer) handleHeader(headerKey string, namespaceKey namespace.Namespace) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerValue := ctx.Request.Header.Get(headerKey)

		// Add to logger
		logger := sharedLog.GetLoggerFromContext(ctx)
		logger = namespaceKey.WithValue(headerValue).AsField(logger)

		setLogger(ctx, logger)

		if headerValue != "" {
			ctx.Writer.Header().Set(headerKey, headerValue)
		}
	}
}

// MARK: Request ID

func GetRequestID(ctx *gin.Context) string {
	return ctx.Request.Header.Get(domain.HeaderXRequestId)
}

func (hs *HttpServer) RequireRequestID(ctx *gin.Context) {
	requestID := ctx.Request.Header.Get(domain.HeaderXRequestId)
	if requestID == "" {
		hs.ReplyWithError(ctx, &apperr.InvalidArgumentError{})
		ctx.Abort()
	}
}

// MARK: Correlation ID

func GetCorrelationID(ctx *gin.Context) string {
	return ctx.Request.Header.Get(domain.HeaderXCorrelationId)
}

func (hs *HttpServer) RequireCorrelationID(ctx *gin.Context) {
	correlationID := ctx.Request.Header.Get(domain.HeaderXCorrelationId)

	if correlationID == "" {
		hs.ReplyWithError(ctx, &apperr.InvalidArgumentError{})
		ctx.Abort()
	}
}
