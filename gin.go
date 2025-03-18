package httpserver

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/yeencloud/lib-httpserver/domain"
	metrics "github.com/yeencloud/lib-metrics"
	MetricsDomain "github.com/yeencloud/lib-metrics/domain"
	sharedMetrics "github.com/yeencloud/lib-shared/metrics"
)

type HttpServer struct {
	config *domain.HttpServerConfig

	Gin *gin.Engine
}

type ContextKey struct {
	name string
}

func (ck ContextKey) String() string {
	return "context." + ck.name
}

func debugPrintRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	entry := log.NewEntry(log.StandardLogger())

	entry = domain.LogHttpMethodField.WithValue(httpMethod).AsField(entry)
	entry = domain.LogHttpPathField.WithValue(absolutePath).AsField(entry)
	entry = domain.LogHttpHandlerCountField.WithValue(nuHandlers).AsField(entry)
	entry = domain.LogHttpHandlerNameField.WithValue(handlerName).AsField(entry)

	entry.Info(fmt.Sprintf("%s %s", httpMethod, absolutePath))
}

func NewHttpServer(config *domain.HttpServerConfig) *HttpServer {
	gin.DebugPrintRouteFunc = debugPrintRoutes
	gin.DebugPrintFunc = func(format string, values ...interface{}) {}

	r := gin.New()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}
	SetCors(r, strings.Split(config.AllowedOrigins, ","))

	gs := &HttpServer{
		Gin: r,

		config: config,
	}

	gs.handleMiddleware()
	gs.handleErrorRoutes()
	r.Use(gs.LogRequest)

	return gs
}

func (gs *HttpServer) Run() error {
	return gs.Gin.Run(fmt.Sprintf("%s:%d", gs.config.Host, gs.config.Port))
}

func (gs *HttpServer) LogRequest(c *gin.Context) {
	path := gs.GetPath(c)
	latency := gs.ProfileNextRequest(c)

	GetLoggerFromContext(c).WithFields(log.Fields{
		"status":  c.Writer.Status(),
		"method":  c.Request.Method,
		"path":    path,
		"latency": latency.Milliseconds(),
	}).Log(gs.MapHttpStatusToLoggingLevel(c), fmt.Sprintf("%s %s %d", c.Request.Method, c.Request.URL.Path, c.Writer.Status()))

	point, ok := c.MustGet(sharedMetrics.MetricsPointKey).(MetricsDomain.Point)
	if !ok {
		point = MetricsDomain.Point{
			Tags: map[string]string{},
		}
	}
	point.Name = "http"
	err := metrics.LogPoint(point, MetricsDomain.Values{
		"response_status": c.Writer.Status(),
		"route_method":    c.Request.Method,
		"route_path":      path,
		"response_ms":     latency.Milliseconds(),
	})

	if err != nil {
		return
	}
}
