package conf

import (
	"github.com/spf13/viper"
	"lancer/data"
	"lancer/global"
)

func IniConfig() {

	global.MysqlData = &data.MysqlData{
		Dsn:         viper.GetString("db.dsn"),
		MaxIdleConn: viper.GetInt("db.maxIdleConn"),
		MaxOpenConn: viper.GetInt("db.maxOpenConn"),
	}

	global.RedisData = &data.RedisData{
		Url:      viper.GetString("redis.url"),
		Password: viper.GetString("redis.password"),
		Db:       viper.GetInt("redis.db"),
	}

	global.JwtData = &data.JwtData{
		Secret: viper.GetString("jwt.secret"),
		Expire: viper.GetInt64("jwt.expire"),
	}

	global.ModeData = &data.ModeData{
		Develop:  viper.GetBool("mode.develop"),
		Language: viper.GetString("mode.language"),
	}

}
