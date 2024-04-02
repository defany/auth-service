package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/defany/slogger/pkg/logger/sl"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "Dev"
)

type Metrics struct {
	ServiceName string `json:"service_name" env-default:"auth_service"`
}

type Server struct {
	GRPC       GRPC       `json:"grpc"`
	HTTP       HTTP       `json:"http"`
	Swagger    Swagger    `json:"swagger"`
	Prometheus Prometheus `json:"prometheus"`
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

type App struct {
	Name string `json:"name" env:"APP_NAME" env-default:"auth_service"`
}

type Config struct {
	Env      string   `json:"env" env-required:"true" env:"ENV"`
	App      App      `json:"app"`
	Metrics  Metrics  `json:"metrics"`
	Server   Server   `json:"server"`
	JWT      JWT      `json:"jwt"`
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
	// TODO: почему-то конфликтует со statik, видимо flag - глобальная тема и ее надо заменить, иначе ошибка redefined
	//confPath := flag.String("config", "", "sets the config path for application")
	//if confPath != nil {
	//	if *confPath != "" {
	//		return *confPath
	//	}
	//}

	envConfPath := os.Getenv("CONFIG_PATH")

	return envConfPath
}
