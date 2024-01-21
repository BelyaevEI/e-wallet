package config

import (
	"github.com/BelyaevEI/e-wallet/internal/models"
	"github.com/spf13/viper"
)

type Config struct {
	DSN  string `mapstructure:"DSN"`
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

// Reading config file for setting application
func LoadConfig(path string) (Config, error) {

	conf := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(models.ConfigName)
	viper.SetConfigType(models.ConfigType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
