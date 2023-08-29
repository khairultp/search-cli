package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

var AppConfig Config

type Config struct {
	ApiUrls string `mapstructure:"API_URLS"`
	CsvFile string `mapstructure:"CSV_FILE"`
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	err = viper.Unmarshal(&AppConfig)
}
