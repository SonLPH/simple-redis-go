package config

import "github.com/spf13/viper"

type Config struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	RedisURL         string `mapstructure:"REDIS_URL"`

	ServerPort   string `mapstructure:"SERVER_PORT"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
}

var _config = Config{}

func LoadConfig(path string) (conf Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&_config)
	_config = conf
	return
}

func GetConfig() Config {
	return _config
}
