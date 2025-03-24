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

func GetMetricsFromContext(ctx *gin.Context) MetricsDomain.Point {
	metricsPoint, ok := ctx.MustGet(metrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		metricsPoint = MetricsDomain.Point{
			Tags: MetricsDomain.Tags{},
		}
	}

	return metricsPoint
}
