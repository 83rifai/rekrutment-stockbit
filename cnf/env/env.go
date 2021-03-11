package env

import "github.com/caarlos0/env/v6"

type ConfigurationEnvironment struct {
	ApplicationEnvironment
	DatabaseEnvironment
}

type ApplicationEnvironment struct {
	AppEnv      string `env:"APP_ENV" envDefault:"development"`
	Development string `env:"DEVELOPMENT" envDefault:"1"`
	HttpPort    string `env:"PORT" envDefault:"8080"`
}

type DatabaseEnvironment struct {
	MaxIdle     int    `env:"DB_MAX_IDLE_CONN" envDefault:"10"`
	MaxOpenConn int    `env:"DB_MAX_OPEN_CONN" envDefault:"100"`
	DBName      string `env:"DB_NAME_MYSQL"`
	DBUser      string `env:"DB_USER_MYSQL"`
	DBPass      string `env:"DB_PASSWORD_MYSQL"`
	DBHost      string `env:"DB_HOST_MYSQL"`
	DBPort      string `env:"DB_PORT_MYSQL"`
}

var Conf = ConfigurationEnvironment{}

func LoadEnv() {
	if err := env.Parse(&Conf); err != nil {
		panic(err.Error())
	}
}
