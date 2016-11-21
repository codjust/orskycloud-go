package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fzzy/radix/pool"
	//"github.com/fzzy/radix/redis"
	"os"
	//"strings"
)

var (
	REDIS_HOST string
	REDIS_DB   int
	red        *pool.Pool
)

func ErrHandlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func init() {
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("REDIS_DB")

	var err error
	red, err = pool.NewPool("tcp", REDIS_HOST, 10)
	ErrHandlr(err)
}

func HandleRegist(username, password string) string {
	// REDIS_HOST = beego.AppConfig.String("redis.host")
	// REDIS_DB, _ = beego.AppConfig.Int("REDIS_DB")

	// red, err := pool.NewPool("tcp", REDIS_HOST, 10)
	// ErrHandlr(err)
	//key := username + "#" + password
	client, err := red.Get()
	ErrHandlr(err)
	res, err := client.Cmd("get", "name").Str()
	ErrHandlr(err)
	red.Put(client)
	// if res != nil {
	// 	return res.Str()
	// }
	// return "some"
	return res
}
