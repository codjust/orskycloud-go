<div class="row-fluid">
		<div class="page-header">
			<h1>My Sensor <small>new create a sensor</small></h1>
		</div>
			<fieldset>
				<div class="control-group">
					<label class="control-label" for="username">传感器标识:</label>
						<div class="controls">
						<input type="text" class="input-xlarge" id="name" value= "" />
						</div>
				</div>
				<div class="control-group">
						<label class="control-label" for="usrkey">传感器名称:</label>
						<div class="controls">
						<input type="text" class="input-xlarge" id="designation"  value= "" />
						</div>
				</div>
				<div class="control-group">
					<label class="control-label" for="phone">单位:</label>
						<div class="controls">
							<input type="text" class="input-xlarge" id="unit" value="" />
						</div>
				</div>
				<div class="control-group">
					<label class="control-label" for="role">所属设备：</label>
					<div class="controls">
						<select id="did">
						{{range .DList}}
							<option value= '{{.Did}}'>{{.DeviceName}}</option>
						{{end}}
						</select>
					</div>
				</div>
				<div class="form-actions">
						<input type="button" class="btn btn-success btn-large" value="确定添加" onclick="SubmitNewSensor()" />
				</div>
			</fieldset>

	<table class="table table-striped table-bordered table-condensed">
	<thead>
	<tr>
		<th>传感器</th>
		<th>
				<div class="control-group">
					<div class="controls">
						<select id="did">
							<option value= '{{.Did}}'>Test</option>
						</select>
					</div>
				</div>
		</th>
		<th>
			<div class="control-group">
					<div class="controls">
						<select id="did">
							<option value= '{{.Did}}'>Test</option>
						</select>
					</div>
				</div>

		</th>
		<th>
			<div class="form-actions">
					<input type="button" class="btn btn-success btn-large" value="查询" />
			</div>
		</th>
		<th>
		<div class="form-actions">
					<input type="button" class="btn btn-success btn-large" value="查询" />
			</div>
		</th>
		<th>操作</th>
		<th></th>
	</tr>
	</thead>

	<tbody>
	{{range .Page.List}}
	<tr class="list-users">
		<td>{{.Name}}</td>
		<td>{{.Device}}</td>
		<td>{{.Designation}}</td>
		<td>{{.Unit}}</td>
		<td>{{.CreateTime}}</td>
		<!-- <td><span class="label label-success">Active</span></td> -->
		<td>
			<div class="btn-group">
			<a class="btn btn-mini dropdown-toggle" data-toggle="dropdown" href="#">Actions <span class="caret"></span></a>
				<ul class="dropdown-menu">
					<li ><a href="/mysensor/edit?did={{.Did}}&&name={{.Name}}"><i class="icon-pencil"></i>编辑</a></li>
					<li onclick="SubmitDeleteSensor({{.Did}},{{.Name}})"><a href="#"><i class="icon-trash"></i> 删除</a></li>
					<li><a href="/mysensor/newdevice"><i class="icon-trash"></i> 新建</a></li>
					<!-- <li><a href="#"><i class="icon-user"></i> Details</a></li> -->
					<!-- <li class="nav-header">Permissions</li>
					<li><a href="#"><i class="icon-lock"></i> Make <strong>Admin</strong></a></li>
					<li><a href="#"><i class="icon-lock"></i> Make <strong>Moderator</strong></a></li>
					<li><a href="#"><i class="icon-lock"></i> Make <strong>User</strong></a></li> -->
				</ul>
			</div>
		</td>
	</tr>
	{{end}}
	</tbody>
</table>
</div>