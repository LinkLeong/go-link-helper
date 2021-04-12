package redis

import (
	gredis "github.com/go-redis/redis/v8"
)

var RedisConn *gredis.Client

//初始化redis
func Setup() {
	RedisConn = gredis.NewClient(
		&gredis.Options{
			Addr: "",
			Password: "",
			DB:3,
		})
}