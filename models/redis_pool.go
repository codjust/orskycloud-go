package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fzzy/radix/pool"
	//"github.com/fzzy/radix/redis"
	"crypto/md5"
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
	key := username + "#" + password
	client, err := red.Get()
	ErrHandlr(err)
	res, err := client.Cmd("hget", "User", key).Str()
	ErrHandlr(err)
	if res == "" {
		return "exist"
	}
	// srcData := []byte("iyannik0215")
	// cipherText1 := md5.Sum(srcData)
	// fmt.Println(cipherText1)

	uid, err := client.Cmd("incr", "nextUserId").Str()
	ErrHandlr(err)
	srcData := []byte(uid)
	uid = md5.Sum(srcData)
	UserList, err := client.Cmd("get", "UserList").Str()
	ErrHandlr(err)
	UserList = UserList + uid + "#"
	client.Cmd("set", "UserList", UserList)
	client.Cmd("hset", "User", key, uid)
	client.Cmd("incr", "UserCount")

	client.Cmd("hset", "uid:"+uid, "username", username)
	client.Cmd("hset", "uid:"+uid, "password", password)
	localtime := time.Now().Unix()
	client.Cmd("hset", "uid:"+uid, "sign_up_time", localtime)

	red.Put(client)
	// if res != nil {
	// 	return res.Str()
	// }
	// return "some"
	return res
}
