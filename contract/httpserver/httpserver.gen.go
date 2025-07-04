// Package httpserver provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package httpserver

import (
	"github.com/gin-gonic/gin"
)

// Response defines model for Response.
type Response struct {
	Body          *map[string]interface{} `json:"body,omitempty"`
	CorrelationId *string                 `json:"correlationId,omitempty"`
	Error         *ResponseError          `json:"error,omitempty"`
	RequestId     *string                 `json:"requestId,omitempty"`
	Status        int                     `json:"status"`
}

// ResponseError defines model for ResponseError.
type ResponseError struct {
	Code     *string `json:"code"`
	HowToFix *string `json:"howToFix"`
	Message  string  `json:"message"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {

}
