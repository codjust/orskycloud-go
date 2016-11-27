package models

import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"orskycloud-go/utils"
	"strings"
)

type Device struct {
	ID          string
	DevName     string
	Description string
	CreateTime  string
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
	return devices, count
}

func PageUser(p int, size int, username string, password string) utils.Page {
	devices, count := ReturnAllDevices(username, password)
	// var user User
	// var list []User
	// qs := o.QueryTable(user)
	// count, _ := qs.Limit(-1).Count()
	// qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	// c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(count, p, size, devices)
}
