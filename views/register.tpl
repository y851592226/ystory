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
                url:"/register",  
                data:$('#frmregister').serialize(),// 序列化表单值  
                async: false,  
                error: function(data) {
                    layer.open({
                        content: '网络异常请',
                        btn: '我知道了',
                        yes: function(index){
                            layer.close(index)
                        }   
                    });
                },
                success: function(data) {
                    if (data == 'success'){
                        layer.open({
                            content: '恭喜你，注册成功！！！',
                            btn: '我知道了',
                            yes: function(index){
                                layer.close(index);
                                window.location.href='/'
                            } ,
                            end: function(index){
                                layer.close(index);
                                window.location.href='/'
                            } 

                        });
                    }
                    else{
                         layer.open({
                            content: data,
                            btn: '我知道了',
                            yes: function(index){
                                layer.close(index)
                            }   
                        });
                    }   
                },
                  
            });  
}
</script>
{{template "topbar2" .}}
<form name="frmregister" id="frmregister">
<div class="login">
<p>帐号：</p><p><input id="username" type="text" name="username"  size="20" maxlength="30" class="login_name"></p><br/>
<p>密码：</p><p><input id="regpass" size="20" maxlength="30" type="password" name="password"  class="login_name"></p><br/>
<p>确认密码：</p><p><input name="repassword" type="password"  size="20" maxlength="30" id="repassword" class="login_name"></p><br/>
<p>Email：</p><input type="text" class="login_name" name="email" id="email" size="25" maxlength="60" /></p>
<input type="hidden" name="action" id="action" value="newuser" />
</div>
</form>
<div class="login_btn"><button type="button" name="Submit" onclick=submit() class="btn">注册</button></div>
<div class="login_btn"><a href="/login">已有账号，点击登录</a></div>
{{template "footerbar" .}}
</html>
{{end}}
