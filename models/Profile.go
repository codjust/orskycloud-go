package models

import (
	"github.com/astaxie/beego"
)

type Profile struct {
	UserName string
	UserKey  string
	Phone    string
	EMail    string
	DevCount string
	SignTime string
}

func ReturnProfileInfo(username string, password string) Profile {
	client, err := red.Get()
	ErrHandlr(err)
	beego.Debug("return...")

	var ProfileInfo Profile
	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	ProfileInfo.UserName, _ = client.Cmd("hget", "uid:"+userkey, "username").Str()
	ProfileInfo.UserKey = userkey
	ProfileInfo.Phone, _ = client.Cmd("hget", "uid:"+userkey, "phone").Str()
	ProfileInfo.EMail, _ = client.Cmd("hget", "uid:"+userkey, "email").Str()
	ProfileInfo.DevCount, _ = client.Cmd("hget", "uid:"+userkey, "count").Str()
	ProfileInfo.SignTime, _ = client.Cmd("hget", "uid:"+userkey, "sign_up_time").Str()

	return ProfileInfo
}

func UpdataProfileInfo(username string, password string, profile Profile) {
	client, err := red.Get()
	ErrHandlr(err)

	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	client.Cmd("multi")
	client.Cmd("hset", "uid:"+userkey, "username", profile.UserName)
	client.Cmd("hset", "uid:"+userkey, "phone", profile.Phone)
	client.Cmd("hset", "uid:"+userkey, "email", profile.EMail)
	ret := client.Cmd("exec").String()
	var ret_msg string
	ret_msg = "success"
	if ret == nil {
		ret_msg = "failed"
		//ErrHandlr("redis exec failed!")
	}
	return ret_msg
}
