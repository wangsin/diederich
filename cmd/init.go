package main

import mconf "github.com/wangsin/diederich/base/myViper"

func initHandler() {
	err := mconf.Init()
	if err != nil {
		panic(err)
	}

}
