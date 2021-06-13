package mdb

import (
	"errors"
	mconf "github.com/wangsin/diederich/base/myViper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	dbConnMap map[string]*gorm.DB

	key = []string{
		"main",
	}
)

func Init() error {
	dbConnMap = make(map[string]*gorm.DB)
	for _, k := range key {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       mconf.Viper.GetString("db_dsn." + k),
			SkipInitializeWithVersion: true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			DontSupportForShareClause: true,
		}))
		if err != nil {
			// todo log
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			// todo log
			return err
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)

		dbConnMap[k] = db
	}

	return nil
}

func GetDBConnection(k string) (*gorm.DB, error) {
	conn, ok := dbConnMap[k]
	if !ok || conn == nil {
		return nil, errors.New("fucked")
	}

	return conn.Session(&gorm.Session{NewDB: true}), nil
}
