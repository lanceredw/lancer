package cmd

import (
	"lancer/conf"
	"lancer/global"
)

func StartMigrate() {

	// init settings
	conf.InitSettings()

	//init config
	conf.IniConfig()

	// init logger
	global.Logger = conf.InitLogger()

	// init Snowflake
	global.Snowflake = conf.InitSnowflake()

	// init mysql database
	db, err := conf.InitDB()
	if err != nil {
		panic(err)
	}
	global.DB = db

}
