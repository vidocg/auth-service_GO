package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost                string `mapstructure:"POSTGRES_HOST"`
	DBUserName            string `mapstructure:"POSTGRES_USER_NAME"`
	DBUserPassword        string `mapstructure:"POSTGRES_PASS"`
	DBName                string `mapstructure:"POSTGRES_DB"`
	DBPort                string `mapstructure:"POSTGRES_PORT"`
	ServerPort            string `mapstructure:"PORT"`
	MigrationFolder       string `mapstructure:"MIGRATION_FOLDER"`
	GoogleClientId        string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret    string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleAuthCallbackUrl string `mapstructure:"GOOGLE_AUTH_CALLBACK"`
}

func LoadConfig(path string) (config Config, err error) {
	profile := os.Getenv("PROFILE")
	if profile == "" {
		profile = "local"
	}
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(profile)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Print(err)
		return
	}

	err = viper.Unmarshal(&config)

	return
}
