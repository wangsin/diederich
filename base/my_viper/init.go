package mconf

import "github.com/spf13/viper"

var Viper *viper.Viper

func Init() error {
	Viper = viper.New()

	Viper.SetConfigName("dev")
	Viper.SetConfigFile("conf/dev.toml")
	err := Viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
