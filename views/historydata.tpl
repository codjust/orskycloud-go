<div class="row-fluid">
		<div class="page-header">
			<h1>My Sensor <small>new create a sensor</small></h1>
		</div>
	<table >
	<div>
	<tr>
		<th><p style = "font-size: 15px">传感器：</p></th>
		<td><select id="did" style="width:120px">
		{{range .Data}}
		<option value={{.Did}}>{{.Dev_Name}}</option>
		{{end}}
		</select></td>
		<td><select id="s_name" style="width:120px">
		<option value=>{{.S_Array.Name}}</option>
		</select></td>
<!-- 	</tr>
	<tr> -->
		<td>&nbsp;&nbsp;&nbsp;</td>
		<td><select id="select" style="width:100px" onclick="SelectTime()">
		<option value="self">自定义时间</option>
		<option value="day">最近一天</option>
		<option value="week">最近一周</option>
		<option value="month">最近一月</option>
		<option value="year">最近一年</option>
		</select></td>
		<td><input type="text" class="input-xlarge" id="start" style="width:130px"/></td>
		<td><p>-</p></td>
		<td><input type="text" class="input-xlarge" id="end"  style="width:130px"/></td>
		<td>&nbsp;&nbsp;&nbsp;</td>
		<div>
		<td><input type="button" value="查询" class="btn btn-success btn-large" /></td>
		<td>&nbsp;&nbsp;</td>
		<td><input type="button" value="删除" class="btn btn-success btn-large" /></td>
		</div>
	</tr>
	</div>
  </table>
<table class="table table-striped table-bordered table-condensed">
	<tr class="list-users">
		<th>标识</th>
		<th>名称</th>
		<th>更新时间</th>
		<th>值</th>
	</tr>
	<tr>
		<td>1</td>
		<td>2</td>
		<td>3</td>
		<td>4</td>
		<td>
			<div class="btn-group">
			<a class="btn btn-mini dropdown-toggle" data-toggle="dropdown" href="#">Actions <span class="caret"></span></a>
				<ul class="dropdown-menu">
					<li ><a href="/mysensor/edit?did={{.Did}}&&name={{.Name}}"><i class="icon-pencil"></i>编辑</a></li>
					<li onclick="SubmitDeleteSensor({{.Did}},{{.Name}})"><a href="#"><i class="icon-trash"></i> 删除</a></li>
					<li><a href="/mysensor/newdevice"><i class="icon-trash"></i> 新建</a></li>
				</ul>
			</div>
		</td>
	</tr>
</table>
</div>