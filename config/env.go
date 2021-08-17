package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	ServerPort         int    `env:"SERVER_PORT" envDefault:"8080"`
	PostgresDBHost     string `env:"POSTGRES_DB_HOST" envDefault:"127.0.0.1"`
	PostgresDBDriver   string `env:"POSTGRES_DB_DRIVER" envDefault:"postgres"`
	PostgresDBUser     string `env:"POSTGRES_DB_USER" envDefault:"root"`
	PostgresDBPassword string `env:"POSTGRES_DB_PASSWORD" envDefault:"12345678"`
	PostgresDBName     string `env:"POSTGRES_DB_NAME" envDefault:"12sq"`
	PostgresDBPort     int    `env:"POSTGRES_DB_PORT" envDefault:"5432"`
	PostgresSSLMode    string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	RedisHost          string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort          int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisUsername      string `env:"REDIS_USERNAME" envDefault:""`
	RedisPassword      string `env:"REDIS_PASSWORD" envDefault:""`
	AccessTokenSecret  string `env:"ACCESS_SECRET" envDefault:"access_token_secret"`
	RefreshTokenSecret string `env:"REFRESH_SECRET" envDefault:"refresh_token_secret"`
}

// AppEnv provide config values
type AppEnv interface {
	loadEnvFromFile()
	GetServerPort() int
	GetPostgresDBHost() string
	GetPostgresDBDriver() string
	GetPostgresDBUser() string
	GetPostgresDBPassword() string
	GetPostgresDBName() string
	GetPostgresDBPort() int
	GetPostgresDBSSLMode() string
	GetRedisHost() string
	GetRedisPort() int
	GetRedisUsername() string
	GetRedisPassword() string
	GetAccessTokenSecret() string
	GetRefreshTokenSecret() string
}

type appEnv struct {
	config envConfig
}

func (ae *appEnv) GetServerPort() int {
	return ae.config.ServerPort
}

func (ae *appEnv) GetPostgresDBHost() string {
	return ae.config.PostgresDBHost
}

func (ae *appEnv) GetPostgresDBDriver() string {
	return ae.config.PostgresDBDriver
}

func (ae *appEnv) GetPostgresDBUser() string {
	return ae.config.PostgresDBUser
}

func (ae *appEnv) GetPostgresDBPassword() string {
	return ae.config.PostgresDBPassword
}

func (ae *appEnv) GetPostgresDBName() string {
	return ae.config.PostgresDBName
}

func (ae *appEnv) GetPostgresDBSSLMode() string {
	return ae.config.PostgresSSLMode
}

func (ae *appEnv) GetPostgresDBPort() int {
	return ae.config.PostgresDBPort
}

func (ae *appEnv) GetRedisHost() string {
	return ae.config.RedisHost
}

func (ae *appEnv) GetRedisUsername() string {
	return ae.config.RedisUsername
}

func (ae *appEnv) GetRedisPort() int {
	return ae.config.RedisPort
}

func (ae *appEnv) GetRedisPassword() string {
	return ae.config.RedisPassword
}

func (ae *appEnv) GetAccessTokenSecret() string {
	return ae.config.AccessTokenSecret
}

func (ae *appEnv) GetRefreshTokenSecret() string {
	return ae.config.RefreshTokenSecret
}

var _env AppEnv

// init initialize ONLY ONCE default AppEnv instance
// this function will be called when import the package
// https://golang.org/doc/effective_go.html#init
func init() {
	_env = &appEnv{}
	_env.loadEnvFromFile()
	Logger().Info("Initialized app environment successfully!")
}

// Env get app configs. Example: constant.Env().GetPort()
func Env() AppEnv {
	return _env
}

// loadEnvFromFile load and parse environment variables
func (ae *appEnv) loadEnvFromFile() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	loadEnvErr := godotenv.Load(envFile)
	if loadEnvErr != nil {
		Logger().Fatal(loadEnvErr)
	}

	parseErr := env.Parse(&ae.config)
	if parseErr != nil {
		Logger().Fatal(parseErr)
	}
}
