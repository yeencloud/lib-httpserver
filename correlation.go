package httpserver

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-httpserver/domain"
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	sharedLog "github.com/yeencloud/lib-shared/log"
	sharedMetrics "github.com/yeencloud/lib-shared/metrics"
)

var HeaderXRequestId = "X-Request-ID"
var HeaderXCorrelationId = "X-Correlation-ID"

func (hs *HttpServer) handleRequestID(ctx *gin.Context) { //nolint:dupl
	requestID := ctx.Request.Header.Get(HeaderXRequestId)

	// Add to logger
	logger, ok := ctx.MustGet(sharedLog.ContextLoggerKey).(*log.Entry)
	if !ok {
		logger = log.NewEntry(log.StandardLogger())
	}
	logger = domain.HttpRequestIdKey.WithValue(requestID).AsField(logger)
	ctx.Set(sharedLog.ContextLoggerKey, logger)

	// Add to metrics
	metricsPoint, ok := ctx.MustGet(sharedMetrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		metricsPoint = MetricsDomain.Point{
			Tags: MetricsDomain.Tags{},
		}
	}
	metricsPoint.SetTag(domain.HttpRequestIdKey.WithValue(requestID))
	ctx.Set(sharedMetrics.MetricsPointKey, metricsPoint)

	if requestID != "" {
		ctx.Writer.Header().Set(HeaderXRequestId, requestID)
	}

	ctx.Next()
}

func (hs *HttpServer) handleCorrelationID(ctx *gin.Context) { //nolint:dupl
	correlationID := ctx.Request.Header.Get(HeaderXCorrelationId)

	// Add to logger
	logger, ok := ctx.MustGet(sharedLog.ContextLoggerKey).(*log.Entry)
	if !ok {
		logger = log.NewEntry(log.StandardLogger())
	}
	logger = domain.HttpCorrelationIdKey.WithValue(correlationID).AsField(logger)
	ctx.Set(sharedLog.ContextLoggerKey, logger)

	// Add to metrics
	metricsPoint, ok := ctx.MustGet(sharedMetrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		metricsPoint = MetricsDomain.Point{
			Tags: MetricsDomain.Tags{},
		}
	}
	metricsPoint.SetTag(domain.HttpCorrelationIdKey.WithValue(correlationID))
	ctx.Set(sharedMetrics.MetricsPointKey, metricsPoint)

	if correlationID != "" {
		ctx.Writer.Header().Set(HeaderXCorrelationId, correlationID)
	}

	ctx.Next()
}

func (hs *HttpServer) GetCorrelationID(ctx *gin.Context) string {
	return ctx.GetString(HeaderXCorrelationId)
}

func (hs *HttpServer) RequireRequestID(ctx *gin.Context) {
	requestID := ctx.Request.Header.Get(HeaderXRequestId)
	if requestID == "" {
		hs.ReplyWithError(ctx, &domain.BadRequestError{
			Msg: "Request ID is required",
		})
		ctx.Abort()
	}
	ctx.Next()
}

func (hs *HttpServer) RequireCorrelationID(ctx *gin.Context) {
	correlationID := ctx.Request.Header.Get(HeaderXCorrelationId)

	ctx.Next()

	if correlationID == "" {
		hs.ReplyWithError(ctx, &domain.BadRequestError{
			Msg: "Correlation ID is required",
		})
		ctx.Abort()
	}
	ctx.Next()
}
