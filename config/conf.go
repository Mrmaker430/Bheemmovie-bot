package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	BotToken string `mapstructure:"BOT_TOKEN"`
}

var (
	cfg Config
)

func InitConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	log.Printf(time.Now().Format("2006/01/02 15:04:05") + " : " + "Configuration Success")
	return &cfg
}
