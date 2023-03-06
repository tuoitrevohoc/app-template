package config

import (
	"github.com/spf13/viper"
)

type Configurations struct {
	Port int `mapstructure:"PORT"`
	DBURL        string `mapstructure:"DB_URL"`
}

func NewConfig() (Configurations, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	var c Configurations
	err := viper.ReadInConfig()
	if err != nil {
		return Configurations{}, nil
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return Configurations{}, nil
	}

	return c, nil
}
