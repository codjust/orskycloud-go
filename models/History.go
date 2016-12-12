package models

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	"orskycloud-go/cache_module"
	"orskycloud-go/comm"
	// "orskycloud-go/utils"
	"strings"
	// "time"
	"strconv"
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
	Name        string
	Designation string
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
		s_tmp.Designation, _ = s_json.GetIndex(i).Get("designation").String()
		beego.Debug("Name:", s_tmp)
		s_list = append(s_list, s_tmp)
	}

	red.Put(client)

	return s_list
}

func ReturnSelectHistory(username, password, Did, Name, Start, End string) ([]HistoryData, int) {
	beego.Debug("ReturnSelectHistory:")
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
	var designation string
	var Count = 0
	sensor_json := dev_json.Get("Sensor")
	for i := 0; i < Get_json_array_len(sensor_json); i++ {
		tmp, _ := sensor_json.GetIndex(i).Get("name").String()
		if tmp == Name {
			designation, _ = sensor_json.GetIndex(i).Get("designation").String()
			beego.Debug("ReturnSelectHistory:Designation", designation)
			break
		}
	}
	data_json := dev_json.Get("data")
	for i := 0; i < Get_json_array_len(data_json); i++ {
		tmp, _ := data_json.GetIndex(i).Get("sensor").String()
		beego.Debug("ReturnSelectHistory: Name", tmp)
		if tmp == Name {
			timestamp, _ := data_json.GetIndex(i).Get("timestamp").String()
			beego.Debug("ReturnSelectHistory:in for", i)
			if comm.CompareTime(Start, timestamp) == true && comm.CompareTime(timestamp, End) == true {
				//value, _ := data_json.GetIndex(i).Get("value").String()
				v_tmp, _ := data_json.GetIndex(i).Get("value").Int()
				value := strconv.Itoa(v_tmp)
				beego.Debug("ReturnSelectHistory:value", v_tmp)
				tmp_data.Name = Name
				tmp_data.Timestamp = timestamp
				tmp_data.Value = value
				tmp_data.Designation = designation
				Data = append(Data, tmp_data) //save value
				beego.Debug("ReturnSelectHistory:tmp_data", tmp_data)
				Count++
			}
		}
	}

	red.Put(client)
	return Data, Count

}
func PaginationSelectData(username, password, Did, Name, Start, End string, Page int) ([]HistoryData, int, int) {
	key := beego.AppConfig.String("cache.historydata.key")
	pageSize, _ := beego.AppConfig.Int("history.page.size")
	var tp int //total page
	var ret_count int
	if cache_module.IsExistCache(key) == false {
		beego.Debug("history data cache not exist.")
		Data, count := ReturnSelectHistory(username, password, Did, Name, Start, End)
		ret_count = count
		tp = count / pageSize
		lastPageSize := 0
		if count%pageSize > 0 {
			tp = count/pageSize + 1
			lastPageSize = count % pageSize
		}

		cacheHistoryData := make([][]HistoryData, tp)
		for i := 0; i < tp; i++ {
			if i == (tp-1) && lastPageSize != 0 {
				cacheHistoryData[i] = make([]HistoryData, lastPageSize)
				temp := Data[(i * pageSize):(i*pageSize + lastPageSize)]
				copy(cacheHistoryData[i], temp)
			} else {
				cacheHistoryData[i] = make([]HistoryData, pageSize)
				temp := Data[(i * pageSize):(i*pageSize + pageSize)]
				copy(cacheHistoryData[i], temp)
			}
		}
		cache_module.PutCache(key, cacheHistoryData, 1000*1000)
	}

	ret_data := cache_module.GetCache(key).([][]HistoryData)
	return ret_data[Page-1], tp, ret_count

}

type Pagination struct {
	TotalPage   int
	CurrentPage int
	Count       int
	Data        []HistoryData
}

func GetHistory(username, password, Did, Name, Start, End string, Page string) Pagination {

	page, _ := strconv.Atoi(Page)
	//返回的数据：CurrentPage, TotalPage, 选中页的数据
	data, totalpage, count := PaginationSelectData(username, password, Did, Name, Start, End, page)

	ret_data := Pagination{totalpage, page, count, data}

	beego.Debug("GetHistory:ret_data->:", ret_data)
	return ret_data
}
