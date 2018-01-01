package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"ystory/models"
	"ystory/utils"
)

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

type ShujiaController struct {
	beego.Controller
}

type AddBookcaseController struct {
	beego.Controller
}

type DeleteBookcaseController struct {
	beego.Controller
}

func (this *DeleteBookcaseController) Get() {
	BookId, err := strconv.Atoi(this.Ctx.Input.Param(":BookId"))
	utils.CkeckError(this, err)
	userName := this.Ctx.Input.GetData("userName").(string)
	err = models.DeleteBookcase(userName,BookId)
	utils.CkeckError(this, err)
	this.Ctx.Redirect(302, "/shujia")
}

func (this *AddBookcaseController) Get() {
	BookId, err := strconv.Atoi(this.Ctx.Input.Param(":BookId"))
	utils.CkeckError(this, err)
	userName := this.Ctx.Input.GetData("userName").(string)
	result := models.AddBookcase(userName,BookId)
	this.Ctx.WriteString(result)
}

func (this *ShujiaController) Get() {
	userName := this.Ctx.Input.GetData("userName").(string)
	shujia, err := models.GetShujia(userName)
	utils.CkeckError(this, err)
	this.Data["Data"] = shujia
	this.TplName = "shujia.tpl"
}


func (this *LoginController) Get() {
	this.Data["Data"] = map[string]string{"BookName": "用户登录", "Title": "用户登录"}
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	result := models.CheckPassword(username, password)
	if result == "success"{
		encodeString := utils.Sign(username)
		this.Ctx.Output.Cookie("_user", encodeString, 8640000)
	}	
	this.Ctx.WriteString(result)
}

func (this *RegisterController) Get() {
	this.Data["Data"] = map[string]string{"BookName": "用户注册", "Title": "用户注册"}
	this.TplName = "register.tpl"
}

func (this *RegisterController) Post() {
	this.TplName = "register.tpl"
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	email := this.GetString("email")
	var result string
	if password != repassword {
		result = "两次密码输入不一致"
	} else {
		result = models.Register(username, password, email)
	}
	if result == "success"{
		encodeString := utils.Sign(username)
		this.Ctx.Output.Cookie("_user", encodeString, 8640000)
	}	
	this.Ctx.WriteString(result)
}
