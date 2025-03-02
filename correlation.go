package httpserver

import (
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain"
)

var HeaderXRequestId = "X-Request-ID"
var HeaderXCorrelationId = "X-Correlation-ID"

func (hs *HttpServer) handleRequestID(ctx *gin.Context) {
	requestID := ctx.Request.Header.Get(HeaderXRequestId)
	if requestID != "" {
		ctx.Set(HeaderXRequestId, requestID)
		ctx.Writer.Header().Set(HeaderXRequestId, requestID)
		SetRequestContextValue(ctx, HeaderXRequestId, requestID)
	}
	ctx.Next()
}

func (hs *HttpServer) handleCorrelationID(ctx *gin.Context) {
	correlationID := ctx.Request.Header.Get(HeaderXCorrelationId)
	if correlationID != "" {
		ctx.Set(HeaderXCorrelationId, correlationID)
		ctx.Writer.Header().Set(HeaderXCorrelationId, correlationID)
		SetRequestContextValue(ctx, domain.HttpCorrelationIdKey, correlationID)
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
		}, nil)
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
		}, nil)
		ctx.Abort()
	}
	ctx.Next()
}
