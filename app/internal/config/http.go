package config

import "net"

type HTTP struct {
	Host string `json:"host" env:"SERVER_HTTP_HOST" env-default:"localhost"`
	Port string `json:"port" env:"SERVER_HTTP_PORT" env-default:"8000"`
}

func (h *HTTP) Addr() string {
	return net.JoinHostPort(h.Host, h.Port)
}
