package models

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	// "orskycloud-go/cache_module"
	"orskycloud-go/comm"
	// "orskycloud-go/utils"
	"strings"
	// "time"
)

type DevSenList struct {
	Did      string
	Dev_Name string
	S_Array  []Sensor
}

type HistoryData struct {
	Name        string //传感器标识
	Designation string //描述或者别名
	Timestamp   string //上传时间
	Value       string //值
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

type S_List struct {
	Name string
}

func GetSenSor(username string, password string, Did string) []S_List {
	client, err := red.Get()
	ErrHandlr(err)

	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	dev_info := client.Cmd("hget", "uid:"+userkey, "did:"+Did).String()
	dev_json, err := simplejson.NewJson([]byte(dev_info))
	ErrHandlr(err)
	s_json := dev_json.Get("Sensor")
	var s_list []S_List
	if Get_json_array_len(s_json) == 0 {
		beego.Debug("Len:", Get_json_array_len(s_json))
		return s_list
	}
	var s_tmp S_List
	for i := 0; i < Get_json_array_len(s_json); i++ {
		s_tmp.Name, _ = s_json.GetIndex(i).Get("name").String()
		beego.Debug("Name:", s_tmp)
		s_list = append(s_list, s_tmp)
	}

	red.Put(client)

	return s_list
}

func ReturnSelectHistory(username, password, Did, Name, Start, End string) {
	client, err := red.Get()
	ErrHandlr(err)

	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	dev_info := client.Cmd("hget", "uid:"+userkey, "did:"+Did).String()
	dev_json, err := simplejson.NewJson([]byte(dev_info))
	ErrHandlr(err)
	var Data []HistoryData
	var tmp_data HistoryData
	data_json := dev_json.Get("data")
	for i := 0; i < Get_json_array_len(data_json); i++ {
		tmp, _ := data_json.GetIndex(i).Get("name")
		if tmp == Name {
			timestamp, _ := data_json.GetIndex(i).Get("timestamp")
			if comm.CompareTime(Start, timestamp) == true && comm.CompareTime(timestamp, End) == true {
				value, _ := data_json.GetIndex(i).Get("value")
				tmp_data.Name = tmp
				tmp_data.timestamp = timestamp

			}
		}
	}
}

func GetHistory(username, password, Did, Name, Start, End string) {
	client, err := red.Get()
	ErrHandlr(err)

	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	dev_info := client.Cmd("hget", "uid:"+userkey, "did:"+Did).String()
	dev_json, err := simplejson.NewJson([]byte(dev_info))
	ErrHandlr(err)
	data_json := dev_json.Get("data")
	// {
	//            "sensor": "weight",
	//            "timestamp": "2016-10-20 14:50:30",
	//            "value": 78
	//    }

}
