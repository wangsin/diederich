package main

import (
	"fmt"
	mdb "github.com/wangsin/diederich/base/myDb"
	mgin "github.com/wangsin/diederich/base/myGin"
	mrds "github.com/wangsin/diederich/base/myRedis"
	mconf "github.com/wangsin/diederich/base/myViper"
	"io"
	"os"
)

func initHandler() {
	err := mconf.Init()
	if err != nil {
		panic(err)
	}

	err = mdb.Init()
	if err != nil {
		panic(err)
	}

	err = mgin.Init()
	if err != nil {
		panic(err)
	}

	err = mrds.Init()
	if err != nil {
		panic(err)
	}
}

func initSign() {
	file, err := os.OpenFile("conf/sign.txt", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	sign, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sign))
	fmt.Println("Diederich Start Listening On Port " + mconf.Viper.GetString("main.port"))
}
