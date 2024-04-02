package config

import "net"

type Prometheus struct {
	Host string `json:"host" env:"SERVER_PROMETHEUS_HOST" env-default:"localhost"`
	Port string `json:"port" env:"SERVER_PROMETHEUS_PORT" env-default:"10000"`
}

func (p *Prometheus) Addr() string {
	return net.JoinHostPort(p.Host, p.Port)
}
