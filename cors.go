package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func setCors(router *gin.Engine, allowedOrigins []string) {
	isWildcard := len(allowedOrigins) == 1 && allowedOrigins[0] == "*"

	config := cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type", "X-Request-ID", "X-Correlation-ID"},
		ExposeHeaders: []string{"Content-Length", "X-Request-ID", "X-Correlation-ID"},
		MaxAge:        12 * time.Hour,
	}

	if isWildcard {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = allowedOrigins
		config.AllowCredentials = true
	}

	router.Use(cors.New(config))
}
