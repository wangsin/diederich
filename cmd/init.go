package main

import (
	mdb "github.com/wangsin/diederich/base/myDb"
	mgin "github.com/wangsin/diederich/base/myGin"
	mrds "github.com/wangsin/diederich/base/myRedis"
	mconf "github.com/wangsin/diederich/base/myViper"
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
