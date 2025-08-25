package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl        string `mapstructure:"DB_URL"`
	GinHttpPort  string `mapstructure:"GIN_HTTP_PORT"`
	GrpcHttpPort string `mapstructure:"GRPC_HTTP_PORT"`
	GrpcPort     string `mapstructure:"GRPC_PORT"`
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
