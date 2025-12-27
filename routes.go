package httpserver

import (
	"encoding/json"
	"fmt"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yeencloud/lib-httpserver/domain"
	"github.com/yeencloud/lib-shared/apperr"
	SharedDomain "github.com/yeencloud/lib-shared/domain"
)

func debugPrintRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	entry := log.NewEntry(log.StandardLogger())

	entry = domain.LogHttpMethodField.WithValue(httpMethod).AsField(entry)
	entry = domain.LogHttpPathField.WithValue(absolutePath).AsField(entry)
	entry = domain.LogHttpHandlerCountField.WithValue(nuHandlers).AsField(entry)
	entry = domain.LogHttpHandlerNameField.WithValue(handlerName).AsField(entry)

	entry.Info(fmt.Sprintf("%s %s", httpMethod, absolutePath))
}

// MARK: Internal Server Error
type InternalServerError struct {
	AdditionalData interface{}
}

func (e *InternalServerError) Error() string {
	marshal, err := json.Marshal(e.AdditionalData)
	placeholderError := "internal server error"

	if err != nil {
		return placeholderError
	}

	return fmt.Sprintf("%s: %s", placeholderError, string(marshal))
}

func (e *InternalServerError) RestCode() int {
	return 500
}

func (e *InternalServerError) Unwrap() error {
	return apperr.InternalError{}
}

func (hs *HttpServer) recoverFromPanic(c *gin.Context, err interface{}) {
	hs.ReplyWithError(c, &InternalServerError{
		AdditionalData: err,
	})
}

func (hs *HttpServer) handleErrorRoutes() {
	hs.Gin.NoRoute(func(ctx *gin.Context) {
		hs.ReplyWithError(ctx, domain.NewPageNotFoundError(ctx.Request.Method, ctx.Request.RequestURI))
	})
}

func (hs *HttpServer) handleMiddleware() {
	r := hs.Gin

	hs.handleErrorRoutes()
	r.Use(hs.securityHeaders())
	r.Use(nice.Recovery(hs.recoverFromPanic))
	r.Use(hs.createLoggerForRequest)
	// Set context
	r.Use(mapRequestContextToContext(domain.HeaderXRequestId, SharedDomain.ContextRequestIdKey))
	r.Use(mapRequestContextToContext(domain.HeaderXCorrelationId, SharedDomain.ContextCorrelationIdKey))
	// Write headers back
	r.Use(hs.handleHeader(domain.HeaderXRequestId, domain.HttpRequestIdKey))
	r.Use(hs.handleHeader(domain.HeaderXCorrelationId, domain.HttpCorrelationIdKey))
}
