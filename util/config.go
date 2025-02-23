package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	EmailPort           string        `mapstructure:"EMAIL_PORT"`
	Environment         string        `mapstructure:"ENVIRONMENT"`
	AllowHeaders        []string      `mapstructure:"ALLOW_HEADERS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	SmtpUser            string        `mapstructure:"SMTP_USER"`
	SmtpPassword        string        `mapstructure:"SMTP_PASSWORD"`
	SmtpHost            string        `mapstructure:"SMTP_HOST"`
	SmtpPort            string        `mapstructure:"SMTP_PORT"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	RedisPassword       string        `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, nil
}
