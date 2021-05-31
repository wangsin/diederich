package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mconf "github.com/wangsin/diederich/base/myViper"
)

func main() {
	initHandler()
	fmt.Println(`
________  ___  _______   ________  _______   ________  ___  ________  ___  ___     
|\   ___ \|\  \|\  ___ \ |\   ___ \|\  ___ \ |\   __  \|\  \|\   ____\|\  \|\  \    
\ \  \_|\ \ \  \ \   __/|\ \  \_|\ \ \   __/|\ \  \|\  \ \  \ \  \___|\ \  \\\  \   
 \ \  \ \\ \ \  \ \  \_|/_\ \  \ \\ \ \  \_|/_\ \   _  _\ \  \ \  \    \ \   __  \  
  \ \  \_\\ \ \  \ \  \_|\ \ \  \_\\ \ \  \_|\ \ \  \\  \\ \  \ \  \____\ \  \ \  \ 
   \ \_______\ \__\ \_______\ \_______\ \_______\ \__\\ _\\ \__\ \_______\ \__\ \__\
    \|_______|\|__|\|_______|\|_______|\|_______|\|__|\|__|\|__|\|_______|\|__|\|__|
                                                                                    `)
	fmt.Println("Diederich Start Listening On Port " + mconf.GlobalViper.GetString("main.port"))

	err := gin.New().Run(mconf.GlobalViper.GetString("main.port"))
	if err != nil {
		panic(err)
	}
}
