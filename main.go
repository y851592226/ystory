package main

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
	_ "ystory/routers"
)

func main() {
    beego.BConfig.Log.AccessLogs = true
    logs.SetLogFuncCallDepth(3)
    beego.SetLogger("file", `{"filename":"log/test.log"}`)
	beego.Run()
}

// beego 存储用户数据
// 已完成 细节优化 登录注册 书架管理

// 自动更新，代码结构优化，代码格式处理优化，代码测试 日志初步处理tag 0.1
// tag 0.2 起点分类，各种榜单，目录，章节，自动换源，优化展示
// 
// tag 0.3 错误报警，增加功能，完善日志
// ，，，起点爬取，书籍索引，分类,排行
// 书本换源，定时更新数据，优化展示效果

// 密码修改，图片显示，增加功能，日志处理
