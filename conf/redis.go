package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"lancer/global"
)

func InitRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     global.RedisData.Url,
		Password: global.RedisData.Password,
		DB:       global.RedisData.Db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, err
}
