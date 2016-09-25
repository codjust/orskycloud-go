package models

import (
	"github.com/garyburd/redigo/redis"
	 "time"
)

func AddUser(username,password string)string,error{
	var err error
	rc := RedisClient.Get()  //从连接池获取一个redis连接
	defer rc.close()

	//判断是否被注册
	name := RedisDo(rc,"HGET","UserId",username)
	if name != nil {
		return "exist",err
	}

	id := RedisDo(rc,"INCR","nextUserId")
	uid,err := redis.string(id,err)
	if err  != nil {
		panic(err.Error())
	}

	_ = RedisDo(rc,"HSET","UserId",username,uid)
	_ = RedisDo(rc, "INCR","UserCount") //用户数量增1

	//密码加盐值
	//初始化数据库用户信息结构
	_ = RedisDo(rc,"HSET","uid:"+uid,"username",username)
	_ = RedisDo(rc,"HSET","uid:"+uid,"password",password)
	_ = RedisDo(rc,"HSET","uid:"+uid,"nextDeviceId ",0)
	_ = RedisDo(rc,"HSET","uid:"+uid,"count ",0)
	_ = RedisDo(rc,"HSET","uid:"+uid,"createtime ",time.Now().Unix())

	return "success",err
}

