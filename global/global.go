package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"lancer/data"
)

//database

var (
	DB    *gorm.DB
	Redis *redis.Client
)

//logger

var (
	Logger *zap.SugaredLogger
)

// Snowflake

var (
	Snowflake *sonyflake.Sonyflake
)

//language

var (
	TranslateMessage map[string]string
)

//data

var (
	MysqlData *data.MysqlData
	RedisData *data.RedisData
	JwtData   *data.JwtData
	ModeData  *data.ModeData
)
