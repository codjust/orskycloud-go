package models

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"orskycloud-go/cache_module"
	"orskycloud-go/comm"
	"orskycloud-go/utils"
	"strings"
	"time"
)

type Device struct {
	ID          string
	DevName     string
	Description string
	CreateTime  string
}

type DeviceJson struct {
	deviceName  string
	description string
	createTime  string
	Sensor      []*Device
	data        []*Device
}

func ReturnAllDevices(username, password string) ([]Device, int) {
	client, err := red.Get()
	ErrHandlr(err)
	beego.Debug("return...")
	var devices []Device
	var device Device
	count := 0
	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()
	device_list_temp, _ := client.Cmd("hget", "uid:"+userkey, "device").Str()
	devices_list := strings.Split(device_list_temp, "#")

	for _, dev := range devices_list {
		count++
		dev_info, _ := client.Cmd("hget", "uid:"+userkey, "did:"+dev).Str()
		dev_json, err := simplejson.NewJson([]byte(dev_info))
		beego.Debug("error:", err)
		ErrHandlr(err)
		device.ID = dev
		device.DevName, err = dev_json.Get("deviceName").String()
		ErrHandlr(err)
		device.Description, err = dev_json.Get("description").String()
		ErrHandlr(err)
		device.CreateTime, err = dev_json.Get("createTime").String()
		ErrHandlr(err)
		devices = append(devices, device)
	}
	red.Put(client)
	beego.Debug("device[0]:", devices[0])
	return devices, count
}

func PageDevice(pageNo int, username string, password string) utils.Page {
	devices, tp, count, pageSize := ReturnDeviceCacheData(username, password, pageNo)
	beego.Debug("dev:", devices, pageNo)
	return utils.Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: devices}
}

func ReturnDeviceCacheData(username string, password string, pageNum int) (interface{}, int, int, int) {
	key := beego.AppConfig.String("cache.device.key")
	pageSize, _ := beego.AppConfig.Int("page.size")
	var tp int //total page
	var ret_count int
	if cache_module.IsExistCache(key) == false {
		dev_list, count := ReturnAllDevices(username, password)
		ret_count = count
		tp = count / pageSize
		lastPageSize := 0
		if count%pageSize > 0 {
			tp = count/pageSize + 1
			lastPageSize = count % pageSize
		}

		cacheDevice := make([][]Device, tp)
		for i := 0; i < tp; i++ {
			if i == (tp-1) && lastPageSize != 0 {
				cacheDevice[i] = make([]Device, lastPageSize)
				temp := dev_list[(i * pageSize):(i*pageSize + lastPageSize)]
				copy(cacheDevice[i], temp)
			} else {
				cacheDevice[i] = make([]Device, pageSize)
				temp := dev_list[(i * pageSize):(i*pageSize + pageSize)]
				copy(cacheDevice[i], temp)
			}
		}
		cache_module.PutCache(key, cacheDevice, 1000*1000*1000)
	}

	devices := cache_module.GetCache(key).([][]Device)
	return devices[pageNum-1], tp, ret_count, pageSize
}

func CreateNewDevice(username string, password string, dev_info Device) string {
	localtime := time.Now().Format("2006-01-02 15:04:05")
	//	exp_data := DeviceJson{deviceName: dev_info.DevName, description: dev_info.Description, createTime: localtime, Sensor: []*Device{}, data: []*Device{}}
	//	exp_json, _ := json.Marshal(exp_data)

	exp_json := "{\"deviceName\":\"" + dev_info.DevName + "\",\"description\":\"" + dev_info.Description + "\",\"createTime\":\"" + localtime + "\",\"Sensor\":[],\"data\":[]}"

	beego.Debug("json:", exp_json)
	client, err := red.Get()
	ErrHandlr(err)
	//key := username + "#" + comm.Md5_go(password)
	key := username + "#" + password
	userkey, _ := client.Cmd("hget", "User", key).Str()

	// get did
	did := client.Cmd("hincrby", "uid:"+userkey, "nextDeviceId", 1).String()
	did = comm.Md5_go(did)
	device_list := client.Cmd("hget", "uid:"+userkey, "device").String()
	device_list = device_list + "#" + did
	client.Cmd("multi")
	client.Cmd("hset", "uid:"+userkey, "device", device_list)
	client.Cmd("hset", "uid:"+userkey, "did:"+did, exp_json)
	client.Cmd("hincrby", "uid:"+userkey, "count", 1)
	ret := client.Cmd("exec").String()
	red.Put(client)
	var ret_msg string
	ret_msg = "success"
	if ret == "" {
		ret_msg = "failed"
		//ErrHandlr("redis exec failed!")
	}
	return ret_msg
}
