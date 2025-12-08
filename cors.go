package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func SetCors(router *gin.Engine, allowedOrigins []string) {
	// Check if wildcard is used
	isWildcard := len(allowedOrigins) == 1 && allowedOrigins[0] == "*"

	config := cors.Config{
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type", "X-Request-ID", "X-Correlation-ID"},
		ExposeHeaders: []string{"Content-Length", "X-Request-ID", "X-Correlation-ID"},
		MaxAge:        12 * time.Hour,
	}

	if isWildcard {
		// When using wildcard, we cannot use AllowCredentials
		// Use AllowAllOrigins instead
		config.AllowAllOrigins = true
	} else {
		// When using specific origins, we can use AllowCredentials
		config.AllowOrigins = allowedOrigins
		config.AllowCredentials = true
	}

	router.Use(cors.New(config))
}
