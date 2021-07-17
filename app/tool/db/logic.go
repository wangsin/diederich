package db

import (
	dbreq "github.com/wangsin/diederich/app/tool/db/req"
	dbresp "github.com/wangsin/diederich/app/tool/db/resp"
	mconf "github.com/wangsin/diederich/base/my_viper"
)

func TestDB(req *dbreq.TestDBRequest) *dbresp.EchoDBConf {
	return &dbresp.EchoDBConf{
		Echo: req,
		Conf: mconf.Viper.GetStringMap("db_dsn"),
	}
}
