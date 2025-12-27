package httpserver

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (hs *HttpServer) setupSession() {
	store := cookie.NewStore([]byte(hs.config.CookieSecret.Value))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   hs.config.CookeMaxAge,
		HttpOnly: true,
		Secure:   hs.config.TLS,
		SameSite: http.SameSiteLaxMode,
	})
	hs.Gin.Use(sessions.Sessions(hs.config.CookieName, store))
}

func (hs *HttpServer) SessionSet(c *gin.Context, key string, value interface{}) error {
	session := sessions.Default(c)
	session.Set(key, value)
	return session.Save()
}

func (hs *HttpServer) SessionGet(c *gin.Context, key string) (interface{}, bool) {
	session := sessions.Default(c)
	value := session.Get(key)
	return value, value != nil
}

func (hs *HttpServer) SessionDelete(c *gin.Context, key string) error {
	session := sessions.Default(c)
	session.Delete(key)
	return session.Save()
}
