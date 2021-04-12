package redis

import (
	gredis "github.com/go-redis/redis/v8"
)

var RedisConn *gredis.Client

//初始化redis
// url 连接地址  如：192.168.200.159：6379
// pwd 密码  如：123456
// db 几号数据库 如：3
func Setup(url,pwd string,db int) {
	RedisConn = gredis.NewClient(
		&gredis.Options{
			Addr: url,
			Password: pwd,
			DB:db,
		})
}