package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnvs(envPath string) {
	// viper.AddConfigPath(envPath)

	viper.AddConfigPath(".") // for docker
	// viper.AddConfigPath("../..")

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	viper.AutomaticEnv()

	dsn := viper.GetString("DSN")
	red := viper.GetString("REDIS")

	_, _ = dsn, red

}
