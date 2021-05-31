package mconf

import "github.com/spf13/viper"

var GlobalViper *viper.Viper

func Init() error {
	GlobalViper = viper.New()

	GlobalViper.SetConfigName("dev")
	GlobalViper.SetConfigFile("../../conf/dev.toml")
	err := GlobalViper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
