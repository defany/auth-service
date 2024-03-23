package config

import "net"

type GRPC struct {
	Host string `json:"host" env:"SERVER_GRPC_HOST" env-default:"localhost"`
	Port string `json:"port" env:"SERVER_GRPC_PORT" env-default:"7000"`
}

func (g *GRPC) Addr() string {
	return net.JoinHostPort(g.Host, g.Port)
}
