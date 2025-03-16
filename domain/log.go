package domain

import (
	"github.com/yeencloud/lib-shared/namespace"
)

var (
	LogHTTPScope = namespace.Namespace{Identifier: "http"}

	LogRouteScope    = namespace.Namespace{Parent: &LogHTTPScope, Identifier: "route"}
	LogResponseScope = namespace.Namespace{Parent: &LogHTTPScope, Identifier: "response"}
)

var (
	LogHttpMethodField       = namespace.Namespace{Parent: &LogRouteScope, Identifier: "method"}
	LogHttpPathField         = namespace.Namespace{Parent: &LogRouteScope, Identifier: "path"}
	LogHttpHandlerCountField = namespace.Namespace{Parent: &LogRouteScope, Identifier: "handler_count"}
	LogHttpHandlerNameField  = namespace.Namespace{Parent: &LogRouteScope, Identifier: "handler"}

	LogHttpResponseStatusCodeField = namespace.Namespace{Parent: &LogResponseScope, Identifier: "status"}
	LogHttpResponseTimeField       = namespace.Namespace{Parent: &LogResponseScope, Identifier: "ms"}
)
