package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-httpserver/domain"
	metrics "github.com/yeencloud/lib-metrics"
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	shared "github.com/yeencloud/lib-shared/log"
)

func (hs *HttpServer) CreateLoggerForRequest(ctx *gin.Context) {
	logEntry := log.NewEntry(log.StandardLogger())

	logEntry = domain.HttpPathKey.WithValue(ctx.Request.RequestURI).AsField(logEntry)
	logEntry = domain.HttpMethodKey.WithValue(ctx.Request.Method).AsField(logEntry)

	setLogger(ctx, logEntry)
}

// TODO: should find a way to log a point with the same key/values as the log entry without having to repeat the key/values setting (maybe an abstraction that can be used by both logrus and influxdb)
func (gs *HttpServer) LogRequest(c *gin.Context) {
	path := gs.GetPath(c)
	latency := gs.ProfileNextRequest(c)

	shared.GetLoggerFromContext(c).WithFields(log.Fields{
		domain.LogHttpResponseStatusCodeField.String(): c.Writer.Status(),
		domain.LogHttpMethodField.String():             c.Request.Method,
		domain.LogHttpPathField.String():               path,
		domain.LogHttpResponseTimeField.String():       latency.Milliseconds(),
	}).Log(gs.MapHttpStatusToLoggingLevel(c), fmt.Sprintf("%s %s %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status()))

	point := GetMetricsFromContext(c)

	point.Name = domain.HttpMetricPointName
	metrics.LogPoint(point, MetricsDomain.Values{
		domain.LogHttpResponseStatusCodeField.MetricKey(): c.Writer.Status(),
		domain.LogHttpMethodField.MetricKey():             c.Request.Method,
		domain.LogHttpPathField.MetricKey():               path,
		domain.LogHttpResponseTimeField.MetricKey():       latency.Milliseconds(),
	})
}
