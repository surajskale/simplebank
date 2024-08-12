package util

import "github.com/spf13/viper"

// the values are read by viper from config file

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` // viper uses mapstructure that's why this is needed
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //json, xml

	viper.AutomaticEnv() // automatically overrides values it has read from config file with values it has read from corresponding environment variables

	err = viper.ReadInConfig() //start reading config values

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
