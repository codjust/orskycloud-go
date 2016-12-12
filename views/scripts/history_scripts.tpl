<script type="text/javascript">
function SelectTime()
{
	var CurrentSelected = document.getElementById("select").value
	var myDate = new Date();
	myDate.getYear(); //获取当前年份(2位)
	Year = myDate.getFullYear();
	Month = myDate.getMonth() + 1; //获取当前月份(0-11,0代表1月)
	Day = myDate.getDate(); //获取当前日(1-31)

	Hour = myDate.getHours(); //获取当前小时数(0-23)
	Minute = myDate.getMinutes(); //获取当前分钟数(0-59)
	Second = myDate.getSeconds(); //获取当前秒数(0-59)

 /**
* 获取本周、本季度、本月、上月的开端日期、停止日期
*/
var now = new Date(); //当前日期
var nowDayOfWeek = now.getDay(); //今天本周的第几天
var nowDay = now.getDate(); //当前日
var nowMonth = now.getMonth(); //当前月
var nowYear = now.getYear(); //当前年
nowYear += (nowYear < 2000) ? 1900 : 0; //

var lastMonthDate = new Date(); //上月日期
lastMonthDate.setDate(1);
lastMonthDate.setMonth(lastMonthDate.getMonth()-1);
var lastYear = lastMonthDate.getYear();
var lastMonth = lastMonthDate.getMonth();

//格局化日期：yyyy-MM-dd
function formatDate(date) {
var myyear = date.getFullYear();
var mymonth = date.getMonth()+1;
var myweekday = date.getDate();
if(mymonth < 10){
mymonth = "0" + mymonth;
}
if(myweekday < 10){
myweekday = "0" + myweekday;
}
return (myyear+"-"+mymonth + "-" + myweekday);
}

//获得某月的天数
function getMonthDays(myMonth){
var monthStartDate = new Date(nowYear, myMonth, 1);
var monthEndDate = new Date(nowYear, myMonth + 1, 1);
var days = (monthEndDate - monthStartDate)/(1000 * 60 * 60 * 24);
return days;
}



//获得本周的开端日期
function getWeekStartDate() {
var weekStartDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek);
return formatDate(weekStartDate);
}

//获得本周的停止日期
function getWeekEndDate() {
var weekEndDate = new Date(nowYear, nowMonth, nowDay + (6 - nowDayOfWeek));
return formatDate(weekEndDate);
}

//获得本月的开端日期
function getMonthStartDate(){
var monthStartDate = new Date(nowYear, nowMonth, 1);
return formatDate(monthStartDate);
}

//获得本月的停止日期
function getMonthEndDate(){
var monthEndDate = new Date(nowYear, nowMonth, getMonthDays(nowMonth));
return formatDate(monthEndDate);
}

//获得上月开端时候
function getLastMonthStartDate(){
var lastMonthStartDate = new Date(nowYear, lastMonth, 1);
return formatDate(lastMonthStartDate);
}

//获得上月停止时候
function getLastMonthEndDate(){
var lastMonthEndDate = new Date(nowYear, lastMonth, getMonthDays(lastMonth));
return formatDate(lastMonthEndDate);
}


   var StartTime = "";
   var EndTime = "";
   switch(CurrentSelected)
   {
   	case "day":
   		EndTime		= Year + "-" + Month + "-" + Day + " " + Hour + ":" + Minute + ":" + Second;
   		Day = Day - 1 ;
   		StartTime   = Year + "-" + Month + "-" + Day + " " + Hour + ":" + Minute + ":" + Second;
   		document.getElementById("start").value = StartTime
   		document.getElementById("end").value = EndTime
   		break;
   	case "week":
   		StartTime = getWeekStartDate() + " " + Hour + ":" + Minute + ":" + Second;
   		EndTime   = getWeekEndDate() + " " + Hour + ":" + Minute + ":" + Second;
   		document.getElementById("start").value = StartTime
   		document.getElementById("end").value = EndTime
   		break;
   	case "month":
   		StartTime = getMonthStartDate() + " " + Hour + ":" + Minute + ":" + Second;
   		EndTime   = getMonthEndDate() + " " + Hour + ":" + Minute + ":" + Second;
   		document.getElementById("start").value = StartTime
   		document.getElementById("end").value = EndTime
   		break;
   	case "year":
   		EndTime = Year + "-" + Month + "-" + Day + " " + Hour + ":" + Minute + ":" + Second;
   		Year = Year - 1;
   		StartTime   = Year + "-" + Month + "-" + Day + " " + Hour + ":" + Minute + ":" + Second;
   		document.getElementById("start").value = StartTime
   		document.getElementById("end").value = EndTime
   		break;
   	case "self":
   		document.getElementById("start").value = ""
   		document.getElementById("end").value = ""
   		break;
   	default:
   		document.getElementById("start").value = StartTime
   		document.getElementById("end").value = EndTime
   		break;
   }
}


function AddSensorItem()
{
	var Did = document.getElementById("did").value;
	//alert(Did)
	$.ajax({
			async: false,
            url: "/history/list",    //后台webservice里的方法名称
            type: "post",
            data:{"did": Did},
            traditional: true,
            success: function (data) {
            	 var optionstring = "";
                for (var i in data) {
                    var jsonObj =data[i];
                        optionstring += "<option value=\"" + jsonObj.Name + "\" >" + jsonObj.Designation + "</option>";
                        $("#s_name").html("<option value='请选择'>请选择...</option> "+optionstring);
                    }
                },
                error: function (msg) {
                    alert("出错了！");
                }
            });
}


document.onload = AddSensorItem()  //页面加载完自动执行此方法

function SearchHistory(Page){
	//声明两个全局变量
var TotalPage
var CurrentPage
var page
	if(Page){
		page = Page
	}
	else{
		page = 1
	}

	//alert(page)
	var h_did  = document.getElementById("did").value
	var h_name = document.getElementById("s_name").value
	var start  = document.getElementById("start").value
	var end    = document.getElementById("end").value

	//alert(h_name)
	if(h_name == "请选择"){
		alert("请选择要查询的传感器！")
		return
	}
	//2015-12-11 19:27:57
	var pattern = /\d{4}-\d{2}-\d{2}\s?\d{2}:\d{2}:(\d+)/
	var r1 = pattern.test(start)
	var r2 = pattern.test(end)
	if(r1 == false || r2 == false)
	{
		alert("时间格式错误，请设置：2015-12-1 12:12:12")
		return;
	}

	$.ajax({
			async: false,
            url: "/history/data",    //后台webservice里的方法名称
            type: "post",
            data:{"did": h_did, "name": h_name, "start":start, "end":end, "page":page},
            traditional: true,
            success: function (data) {
            	TotalPage = data.TotalPage
            	CurrentPage = data.CurrentPage
            	//alert("success")
            	//alert(TotalPage)
            	//alert(CurrentPage)
            	var tablestring = "";
            	var arr = data.Data
                for (var i in arr) {
                    var jsonObj = arr[i];
                    tablestring += "<tr><td>" + jsonObj.Name + "</td><td>" + jsonObj.Designation + "</td><td>" + jsonObj.Timestamp + "</td><td>" + jsonObj.Value + "</td></tr>"
                    }
                $("#tb").html("<tr class='list-users'><th>标识</th><th>名称</th><th>更新时间</th><th>值</th></tr>"+tablestring);
                },
                error: function (msg) {
                    alert("出错了！");
                }
            });

$(function () {
	//alert("pagination1")
	//alert(TotalPage)
	// var pageno = CurrentPage
	// var allpage = TotalPage
    $("#pagination1").bootstrapPaginator({
      currentPage: CurrentPage,
      totalPages: TotalPage,
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        //window.location.href = "/mysensor/" + page
        SearchHistory(page);
      }
    });
  });

}

// $(function () {
// 	alert("pagination1")
// 	alert(TotalPage)
// 	// var pageno = CurrentPage
// 	// var allpage = TotalPage
//     $("#pagination1").bootstrapPaginator({
//       currentPage: CurrentPage,
//       totalPages: TotalPage,
//       bootstrapMajorVersion: 3,
//       size: "small",
//       onPageClicked: function(e,originalEvent,type,page){
//         //window.location.href = "/mysensor/" + page
//         SearchHistory(page);
//       }
//     });
//   });

</script>