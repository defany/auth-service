package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/defany/slogger/pkg/logger/sl"
	"github.com/ilyakaznacheev/cleanenv"
)

type Metrics struct {
	ServiceName string `json:"service_name" env-default:"auth_service"`
}

type Server struct {
	Port int `json:"port" env:"SERVER_PORT" env-default:"50001"`
}

type Database struct {
	Username             string        `json:"username" env:"DATABASE_USERNAME" env-required:"true"`
	Password             string        `json:"password" env:"DATABASE_PASSWORD" env-required:"true"`
	Host                 string        `json:"host" env:"DATABASE_HOST" env-required:"true"`
	Port                 string        `json:"port" env:"DATABASE_PORT" env-required:"true"`
	Database             string        `json:"database" env:"DATABASE" env-required:"true"`
	MigrationsDir        string        `json:"migrations_dir" env:"DATABASE_MIGRATIONS_DIR" env-default:"migrations"`
	ConnectAttempts      int           `json:"connect_attempts" env:"DATABASE_CONNECT_ATTEMPTS" env-default:"3"`
	ConnectAttemptsDelay time.Duration `json:"connect_attempts_delay" env:"DATABASE_CONNECT_ATTEMPTS_DELAY" env-default:"5s"`
}

type Config struct {
	Env      string   `json:"env" env-required:"true" env:"ENV"`
	Metrics  Metrics  `json:"metrics"`
	Server   Server   `json:"server"`
	Database Database `json:"database"`
	Logger   sl.Slog  `json:"logger"`
}

func MustLoad() *Config {
	confPath := configPath()
	if confPath == "" {
		log.Fatalln("config path cannot be empty set them by `config` flag or `CONFIG_PATH` in env")
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Fatalln(fmt.Sprintf("file with this path was not found: %s", confPath))
	}

	var cfg Config

	if envErr := cleanenv.ReadEnv(&cfg); envErr != nil {
		if err := cleanenv.ReadConfig(confPath, &cfg); err != nil {
			log.Fatalf("cannot read config: %s", errors.Join(err, envErr))
		}
	}

	if err := cleanenv.ReadConfig(confPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func configPath() string {
	confPath := flag.String("config", "", "sets the config path for application")
	if confPath != nil {
		if *confPath != "" {
			return *confPath
		}
	}

	envConfPath := os.Getenv("CONFIG_PATH")

	return envConfPath
}
