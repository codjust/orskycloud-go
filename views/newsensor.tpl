<div class="row-fluid">
	<div class="page-header">
		<h1>My Sensor <small>new create a sensor</small></h1>
	</div>
		<form class="form-horizontal" method = "post" action = >
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
							<input type="text" class="input-xlarge" id="againsure" value="" />
						</div>
				</div>
				<div class="control-group">
					<label class="control-label" for="role">所属设备：</label>
					<div class="controls">
						<select id="role">
							<option value="admin">Admin</option>
							<option value="mod">Moderator</option>
							<option value="user" selected>User</option>
							</select>
						</div>
					</div>
				<div class="form-actions">
						<input type="button" class="btn btn-success btn-large" value="确定修改" onclick="SubmitModifyPwd()" />
				</div>
				</fieldset>
		</form>
</div>