package domain

import (
	"github.com/yeencloud/lib-shared/log"
)

var HttpContextKey = log.Path{Identifier: "http"}
var HttpRequestContextKey = log.Path{Parent: &HttpContextKey, Identifier: "request"}
var HttpResponseContextKey = log.Path{Parent: &HttpContextKey, Identifier: "response"}

var (
	HttpPathKey          = log.Path{Parent: &HttpRequestContextKey, Identifier: "path"}
	HttpFullPathKey      = log.Path{Parent: &HttpRequestContextKey, Identifier: "full_path"}
	HttpMethodKey        = log.Path{Parent: &HttpRequestContextKey, Identifier: "method"}
	HttpCorrelationIdKey = log.Path{Parent: &HttpRequestContextKey, Identifier: "correlation_id", IsMetricTag: true}
)
