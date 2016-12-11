package models

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	// "orskycloud-go/cache_module"
	// "orskycloud-go/comm"
	// "orskycloud-go/utils"
	"strings"
	// "time"
)

type DevSenList struct {
	Did      string
	Dev_Name string
	S_Array  []Sensor
}

func GetDevSenList(username string, password string) []DevSenList {
	client, err := red.Get()
	ErrHandlr(err)

	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()

	var ret_array []DevSenList
	var value DevSenList
	device_list_temp, _ := client.Cmd("hget", "uid:"+userkey, "device").Str()
	//beego.Debug("list:", device_list_temp)
	devices_list := strings.Split(device_list_temp, "#")
	for _, did := range devices_list {
		dev_info, _ := client.Cmd("hget", "uid:"+userkey, "did:"+did).Str()
		dev_json, err := simplejson.NewJson([]byte(dev_info))
		ErrHandlr(err)
		value.Dev_Name, err = dev_json.Get("deviceName").String()
		ErrHandlr(err)
		value.Did = did
		s_json := dev_json.Get("Sensor")
		if Get_json_array_len(s_json) == 0 {
			beego.Debug("Len:", Get_json_array_len(s_json))
			ret_array = append(ret_array, value)
			continue
		}
		var s_tmp Sensor
		for i := 0; i < Get_json_array_len(s_json); i++ {
			s_tmp.Name, _ = s_json.GetIndex(i).Get("name").String()
			value.S_Array = append(value.S_Array, s_tmp)
		}

		ret_array = append(ret_array, value)
	}

	red.Put(client)

	return ret_array
}
