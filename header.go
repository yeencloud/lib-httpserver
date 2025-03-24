package httpserver

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain"
	"github.com/yeencloud/lib-httpserver/domain/error"
	sharedLog "github.com/yeencloud/lib-shared/log"
	sharedMetrics "github.com/yeencloud/lib-shared/metrics"
	"github.com/yeencloud/lib-shared/namespace"
)

func (hs *HttpServer) handleHeader(headerKey string, namespaceKey namespace.Namespace) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerValue := ctx.Request.Header.Get(headerKey)

		// Add to logger
		logger := sharedLog.GetLoggerFromContext(ctx)
		logger = namespaceKey.WithValue(headerValue).AsField(logger)

		setLogger(ctx, logger)

		// Add to metrics
		metricsPoint := GetMetricsFromContext(ctx)
		metricsPoint.SetTag(namespaceKey.WithValue(headerValue))
		ctx.Set(sharedMetrics.MetricsPointKey, metricsPoint)

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
		hs.ReplyWithError(ctx, &HttpError.RequestIDRequiredError{})
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
		hs.ReplyWithError(ctx, &HttpError.CorrelationIDRequiredError{})
		ctx.Abort()
	}
}
