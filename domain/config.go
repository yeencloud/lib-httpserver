package domain

type HttpServerConfig struct {
	// Bind Address
	Host string `config:"HTTP_BIND_ADDRESS" default:"0.0.0.0"`
	Port int    `config:"HTTP_BIND_PORT" default:"8080"`

	// CORS
	AllowedOrigins string `config:"HTTP_ALLOWED_ORIGINS" default:"*"`
}
