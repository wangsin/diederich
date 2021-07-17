package main

import (
	"fmt"
	mdb "github.com/wangsin/diederich/base/my_db"
	mgin "github.com/wangsin/diederich/base/my_gin"
	mrds "github.com/wangsin/diederich/base/my_redis"
	mconf "github.com/wangsin/diederich/base/my_viper"
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

	//tracer, closer := base.InitJaeger()
	//defer closer.Close()

	//tracer.StartSpan()

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
