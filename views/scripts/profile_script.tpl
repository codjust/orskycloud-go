<script type="text/javascript">

function SubmitProfileInfo(){
	var Username = document.getElementById("username").value();
	var Phone    = document.getElementById("phone").value();
	var Email    = document.getElementById("email").value();

	if(Username.length < 4)
		alert("用户名长度不合法！");
	else if(Phone.length < 8 || Phone.length > 11)
		alert("电话号码长度不合法");
	else if(Email.length<8)
		alert("E-Mail 格式不合法");

	$.ajax({
		async: false,
		type:"POST",
		url:"/myprofile/update",
		data:{"username": Username,"phone": Phone, "email" : Email}
		}).done(function(msg){
		if(msg.Val == "success")
		{
			alert("更新成功！");
			window.location.href = "/myprofile";
		}
		else if(msg.Val == "failed")
		{
			alert("更新失败，数据库操作错误，请重试！");
		}

	});
}

</script>