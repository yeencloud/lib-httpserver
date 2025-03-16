package httpserver

import (
	"github.com/gin-gonic/gin"

	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	"github.com/yeencloud/lib-shared/metrics"
)

func (hs *HttpServer) CreateMetricsForRequest(ctx *gin.Context) {
	ctx.Set(metrics.MetricsPointKey, MetricsDomain.Point{
		Tags: MetricsDomain.Tags{},
	})
	ctx.Set(metrics.MetricsValuesKey, MetricsDomain.Values{})
}
