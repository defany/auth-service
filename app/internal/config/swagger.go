package config

import "net"

type Swagger struct {
	Host string `json:"host" env:"SERVER_SWAGGER_HOST" env-default:"localhost"`
	Port string `json:"port" env:"SERVER_SWAGGER_PORT" env-default:"9000"`
}

func (s *Swagger) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}
