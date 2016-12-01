package models

import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	//	"orskycloud-go/cache_module"
	//"orskycloud-go/utils"
	"strings"
)

type Sensor struct {
	Name        string
	Device      string
	Designation string
	Unit        string
	CreateTime  string
}

func Get_json_array_len(data *simplejson.Json) int {
	data_arr, err := data.Array()
	if err != nil {
		panic(err.Error())
	}

	return len(data_arr)
}

func ReturnSensorInfo(username string, password string) []Sensor {
	client, err := red.Get()
	ErrHandlr(err)

	var SensorInfo []Sensor
	var temp_sensor Sensor
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	device_list_temp, _ := client.Cmd("hget", "uid:"+userkey, "device").Str()
	devices_list := strings.Split(device_list_temp, "#")
	for _, did := range devices_list {
		dev_info, _ := client.Cmd("hget", "uid:"+userkey, "did:"+did).Str()
		dev_json, err := simplejson.NewJson([]byte(dev_info))
		ErrHandlr(err)
		dev_name, _ := dev_json.Get("deviceName").String()
		sensor := dev_json.Get("Sensor")
		beego.Debug("len:", Get_json_array_len(sensor))
		if Get_json_array_len(sensor) == 1 {
			continue
		}
		for i := 0; i < Get_json_array_len(sensor); i++ {
			temp_sensor.Device = dev_name
			temp_sensor.Name, _ = sensor.GetIndex(i).Get("name").String()
			temp_sensor.Designation, _ = sensor.GetIndex(i).Get("designation").String()
			temp_sensor.Unit, _ = sensor.GetIndex(i).Get("unit").String()
			temp_sensor.CreateTime, _ = sensor.GetIndex(i).Get("createTime").String()
			SensorInfo = append(SensorInfo, temp_sensor)
			beego.Debug("data1:", temp_sensor)
		}
	}

	beego.Debug("data:", SensorInfo)
	return SensorInfo

}
