package database

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedis() *redis.Client {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	poolsize := viper.GetInt("redis.poolsize")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // 密码
		DB:       db,       // 数据库
		PoolSize: poolsize, // 连接池大小
	})

	return rdb
}
