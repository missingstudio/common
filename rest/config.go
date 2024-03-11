package rest

import "time"

// Config for a HTTP server
type Config struct {
	HTTPAddress                   string        `yaml:"address"`
	HTTPPort                      int           `yaml:"port"`
	HTTPServerReadTimeout         time.Duration `yaml:"readTimeout"`
	HTTPServerWriteTimeout        time.Duration `yaml:"writeTimeout"`
	HTTPServerIdleTimeout         time.Duration `yaml:"idleTimeout"`
	ServerGracefulShutdownTimeout time.Duration `yaml:"gracefulShutdownTimeout"`
}

func (c *Config) SetDefaults() {
	c.HTTPAddress = ""
	c.HTTPPort = 8080
	c.HTTPServerIdleTimeout = 30 * time.Second
	c.HTTPServerReadTimeout = 30 * time.Second
	c.HTTPServerWriteTimeout = 30 * time.Second
	c.ServerGracefulShutdownTimeout = 30 * time.Second
}
