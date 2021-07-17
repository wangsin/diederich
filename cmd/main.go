package main

import (
	mgin "github.com/wangsin/diederich/base/my_gin"
	mconf "github.com/wangsin/diederich/base/my_viper"
)

func main() {
	initHandler()
	initSign()

	err := mgin.GinEngine.Run(mconf.Viper.GetString("main.port"))
	if err != nil {
		panic(err)
	}
}
