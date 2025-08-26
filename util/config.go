package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl              string        `mapstructure:"DB_URL"`
	GinHttpPort        string        `mapstructure:"GIN_HTTP_PORT"`
	GrpcHttpPort       string        `mapstructure:"GRPC_HTTP_PORT"`
	GrpcPort           string        `mapstructure:"GRPC_PORT"`
	GoogleClientID     string        `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string        `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURL  string        `mapstructure:"GOOGLE_REDIRECT_URL"`
	PasetoSymmetricKey string        `mapstructure:"PASETO_SYMMETRIC_KEY"`
	TokenDuration      time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}
