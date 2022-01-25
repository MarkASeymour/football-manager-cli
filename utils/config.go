package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() AppConfig {
	viper.SetConfigName("appconfig")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./appconfig/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading config \n", err)
		os.Exit(1)
	}
	ac := AppConfig{FootballApiKey: viper.GetString("apiFootball.apiKey")}
	if (viper.GetString("apiFootball.apiKey")) == "" {
		log.Fatal("Please add a key to appconfig")
		os.Exit(1)
	}
	return ac

}

type AppConfig struct {
	FootballApiKey string
}
