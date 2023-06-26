package configs

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) error {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	return viper.ReadInConfig()
}
