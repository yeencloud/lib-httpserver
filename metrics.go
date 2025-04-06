package httpserver

import (
	"github.com/gin-gonic/gin"
	metrics "github.com/yeencloud/lib-metrics"

	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	SharedMetrics "github.com/yeencloud/lib-shared/metrics"
)

func (hs *HttpServer) CreateMetricsForRequest(ctx *gin.Context) {
	ctx.Set(SharedMetrics.MetricsPointKey, metrics.NewPoint())

}

func GetMetricsFromContext(ctx *gin.Context) MetricsDomain.Point {
	metricsPoint, ok := ctx.MustGet(SharedMetrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		metricsPoint = MetricsDomain.Point{
			Tags: MetricsDomain.Tags{},
		}
	}

	return metricsPoint
}
