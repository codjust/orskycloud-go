package models

import{
	"github.com/astaxie/beego"
	"github.com/fzzy/radix/redis"
	"github.com/fzzy/radix/pool"
	"fmt"
	"os"
}

var{
	REDIS_HOST string
	REDIS_DB   int
}

func ErrHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func init(){
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB,_ = beego.AppConfig.Int("REDIS_DB")

	red, err := pool.NewPool("tcp",REDIS_HOST, 10)
	ErrHndlr(err)
}