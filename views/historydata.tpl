<div class="row-fluid">
		<div class="page-header">
			<h1>My Sensor <small>new create a sensor</small></h1>
		</div>
	<table >
	<div>
	<tr>
		<th><p style = "font-size: 15px">传感器：</p></th>
		<td><select id="did" style="width:120px"><option value=>Test</option></select></td>
		<td><select id="did" style="width:120px"><option value=>Test</option></select></td>
<!-- 	</tr>
	<tr> -->
		<td>&nbsp;&nbsp;&nbsp;</td>
		<td><select id="did" style="width:100px"><option value=>自定义时间</option></select></td>
		<td><input type="text" class="input-xlarge" id="phone" value='{{.Profile.Phone}}' style="width:110px"/></td>
		<td><p>-</p></td>
		<td><input type="text" class="input-xlarge" id="phone" value='{{.Profile.Phone}}' style="width:110px"/></td>
		<td>&nbsp;&nbsp;&nbsp;</td>
		<div>
		<td><input type="button" value="查询" class="btn btn-success btn-large" /></td>
		<td>&nbsp;&nbsp;</td>
		<td><input type="button" value="删除" class="btn btn-success btn-large" /></td>
		</div>
	</tr>
	</div>
<!-- 	<tr>
		<td>传感器</td>
		<td>
			<div class="control-group">
				<div class="controls">
					<select id="did">
						<option value= '{{.Did}}'>Test</option>
					</select>
				</div>
			</div>
		</td>
		<td>
		<div class="control-group">
					 <div class="controls">
						<select id="did">
							<option value= '{{.Did}}'>Test</option>
						</select>
				</div>
		</div>

		</td>
	</tr>
	<tr>
		<td>自定义时间</td>
		<td>
			<div class="control-group">
				<div class="controls">
					<select id="did">
						<option value= '{{.Did}}'>Time</option>
					</select>
				</div>
			</div>
		</td>
		<td>
			<div class="control-group">
					<div class="controls">
						<select id="did">
							<option value= '{{.Did}}'>Time</option>
						</select>
					</div>
				</div>

		</td>
	</tr>
		<td>
			<div >
					<input type="button" class="btn btn-success btn-large" value="查询" />
			</div>
		</td>
		<td>
		<div >
					<input type="button" class="btn btn-success btn-large" value="删除" />
		</div>
		</td> -->
  </table>
<table class="table table-striped table-bordered table-condensed">
	<tr class="list-users">
		<th>标识</th>
		<th>名称</th>
		<th>更新时间</th>
		<th>值</th>
		<!-- <td><span class="label label-success">Active</span></td> -->
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