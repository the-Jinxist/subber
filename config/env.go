package config

import "github.com/spf13/viper"

func LoadEnvs(envPath string) {
	// viper.AddConfigPath(envPath)
	viper.AddConfigPath("../../")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")

	// viper.AddConfigPath("./.")
	// viper.AddConfigPath("././.")

	// viper.AddConfigPath("../..")
	// viper.AddConfigPath("..")

	// viper.AddConfigPath("../../")
	// viper.AddConfigPath("../../..")
	// viper.AddConfigPath("../../../")
	// viper.AddConfigPath("../../../..")

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	dsn := viper.GetString("DSN")
	red := viper.GetString("REDIS")

	_, _ = dsn, red

}
