package domain

import (
	"github.com/yeencloud/lib-shared/namespace"
)

var HttpContextKey = namespace.Namespace{Identifier: "http"}
var HttpRequestContextKey = namespace.Namespace{Parent: &HttpContextKey, Identifier: "request"}
var HttpResponseContextKey = namespace.Namespace{Parent: &HttpContextKey, Identifier: "response"}

var (
	HttpPathKey          = namespace.Namespace{Parent: &HttpRequestContextKey, Identifier: "path"}
	HttpFullPathKey      = namespace.Namespace{Parent: &HttpRequestContextKey, Identifier: "fullPath"}
	HttpMethodKey        = namespace.Namespace{Parent: &HttpRequestContextKey, Identifier: "method"}
	HttpCorrelationIdKey = namespace.Namespace{Parent: &HttpRequestContextKey, Identifier: "correlationId", IsMetricTag: true}
	HttpRequestIdKey     = namespace.Namespace{Parent: &HttpRequestContextKey, Identifier: "requestId", IsMetricTag: true}
)
