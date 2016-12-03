package models

import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"orskycloud-go/cache_module"
	"orskycloud-go/utils"
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

func ReturnSensorInfo(username string, password string) ([]Sensor, int) {
	client, err := red.Get()
	ErrHandlr(err)

	var SensorInfo []Sensor
	var temp_sensor Sensor
	count := 0
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
			count++
		}
	}

	red.Put(client)
	beego.Debug("data:", SensorInfo)
	return SensorInfo, count

}

func ReturnSensorCacheData(username string, password string, pageNum int) (interface{}, int, int, int) {
	key := beego.AppConfig.String("cache.sensor.key")
	pageSize, _ := beego.AppConfig.Int("page.size")
	var tp int //total page
	var ret_count int
	if cache_module.IsExistCache(key) == false {
		dev_list, count := ReturnSensorInfo(username, password)
		ret_count = count
		tp = count / pageSize
		lastPageSize := 0
		if count%pageSize > 0 {
			tp = count/pageSize + 1
			lastPageSize = count % pageSize
		}

		cacheDevice := make([][]Sensor, tp)
		for i := 0; i < tp; i++ {
			if i == (tp-1) && lastPageSize != 0 {
				cacheDevice[i] = make([]Sensor, lastPageSize)
				temp := dev_list[(i * pageSize):(i*pageSize + lastPageSize)]
				copy(cacheDevice[i], temp)
			} else {
				cacheDevice[i] = make([]Sensor, pageSize)
				temp := dev_list[(i * pageSize):(i*pageSize + pageSize)]
				copy(cacheDevice[i], temp)
			}
		}
		cache_module.PutCache(key, cacheDevice, 1000*1000*1000)
	}

	devices := cache_module.GetCache(key).([][]Sensor)
	return devices[pageNum-1], tp, ret_count, pageSize
}

func PageSensor(pageNo int, username string, password string) utils.Page {
	sensors, tp, count, pageSize := ReturnSensorCacheData(username, password, pageNo)
	beego.Debug("dev:", sensors, pageNo)
	return utils.Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: sensors}
}
