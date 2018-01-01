{{with .Data}}
<!DOCTYPE html>
<html>
<head>
{{template "head" .}}
</head>
<script type="text/javascript">
submit = function(){
$.ajax({  
                type: "POST",  
                url:"/login",  
                data:$('#frmlogin').serialize(),// 序列化表单值  
                async: false,  
                success: function(data) {
                    if (data == 'success'){
                        window.location.href='/';
                    }
                    else{
                         layer.open({
                            content: data,
                            btn: '我知道了', 
                        });
                    }   
                }  
            });  
}
</script>
{{template "topbar2" .}}
<form name="frmlogin" id="frmlogin">
<div class="login">
<p>帐号：</p><p><input id="username" type="text" name="username"  size="20" maxlength="30" class="login_name"></p><br/>
<p>密码：</p><p><input id="regpass" size="20" maxlength="30" type="password" name="password"  class="login_name"></p><br/>
</div>
</form>
<div class="login_btn"><button type="button" name="Submit" onclick=submit() class="btn">登陆</button></div>
<div class="login_btn"><a href="/register">没有账号？点击注册</a></div>
{{template "footerbar" .}}
</html>
{{end}}

