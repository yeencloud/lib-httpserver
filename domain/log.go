package domain

import (
	"github.com/yeencloud/lib-shared/log"
)

var (
	LogHTTPScope = log.Path{Identifier: "http"}

	LogRouteScope    = log.Path{Parent: &LogHTTPScope, Identifier: "route"}
	LogResponseScope = log.Path{Parent: &LogHTTPScope, Identifier: "response"}
)

var (
	LogHttpMethodField       = log.Path{Parent: &LogRouteScope, Identifier: "method"}
	LogHttpPathField         = log.Path{Parent: &LogRouteScope, Identifier: "path"}
	LogHttpHandlerCountField = log.Path{Parent: &LogRouteScope, Identifier: "handler_count"}
	LogHttpHandlerNameField  = log.Path{Parent: &LogRouteScope, Identifier: "handler"}

	LogHttpResponseStatusCodeField = log.Path{Parent: &LogResponseScope, Identifier: "status"}
	LogHttpResponseTimeField       = log.Path{Parent: &LogResponseScope, Identifier: "ms"}
)
