package HttpConfig

import "github.com/yeencloud/lib-shared/config"

type HttpServerConfig struct {
	// Bind Address
	Host string `config:"HTTP_BIND_ADDRESS" default:"0.0.0.0"`
	Port int    `config:"HTTP_BIND_PORT" default:"8080"`

	// CORS
	AllowedOrigins string `config:"HTTP_ALLOWED_ORIGINS" default:"*"`

	// Encryption
	TLS     bool   `config:"HTTP_TLS" default:"false"`
	TLSCert string `config:"HTTP_TLS_CERT" default:""`
	TLSKey  string `config:"HTTP_TLS_KEY" default:""`

	CookieSecret config.Secret `config:"HTTP_COOKIE_SECRET" default:""`
	//TODO: Do not forget about changing session name key
	CookieName  string `config:"HTTP_COOKIE_STORE_NAME" default:"session_key_name_to_change_at_some_point"`
	CookeMaxAge int    `config:"HTTP_COOKIE_MAX_AGE" default:"604800"` // 7 days
}
