package models

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义常量
var (
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

func init() {
	//从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB = beego.AppConfig.Int("redis.db")

	//建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			client.Do("SELECT", REDIS_DB)
			beego.Info("create redis pool success,select db:", REDIS_DB)
			return client, nil
		},
	}
}

func RedisDo(client redis.Conn, args ...string) interface{} {
	switch len(args) {
	case 1:
		res, err := client.Do(args[0])
		if err != nil {
			panic(err.Error())
		}
		return value
	case 2:
		res, err := client.Do(args[0], args[1])
		if err != nil {
			panic(err.Error())
		}
		return value
	case 3:
		res, err := client.Do(args[0], args[1], args[2])
		if err != nil {
			panic(err.Error())
		}
		return value
	}
}
