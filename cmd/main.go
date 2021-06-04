package main

import (
	mgin "github.com/wangsin/diederich/base/myGin"
	mconf "github.com/wangsin/diederich/base/myViper"
)

func main() {
	initHandler()
	initSign()

	err := mgin.GinEngine.Run(mconf.Viper.GetString("main.port"))
	if err != nil {
		panic(err)
	}
}
