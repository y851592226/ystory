package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"ystory/controllers"
	"ystory/utils"
)

var CheckUser = func(ctx *context.Context) {
	user := ctx.Input.Cookie("_user")
	if user != "" {
		userName, err := utils.UnSign(user)
		if err == nil {
			ctx.Output.Cookie("_user", user, 8640000)
			ctx.Input.SetData("userName", userName)
		}

	} else {
		encodeString := utils.Sign(utils.GetAnonymousName())
		ctx.Output.Cookie("_user", encodeString, 8640000)
		ctx.Redirect(302, "/")
	}
	return
}

var LogAccess = func(ctx *context.Context) {
	userName := ctx.Input.GetData("userName")
	if userName == nil {
		panic("userName error")
	}
	utils.LogAccess(userName.(string), ctx.Input.URL())
	return
}

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, LogAccess)
	beego.InsertFilter("/*", beego.BeforeRouter, CheckUser)
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/shujia", &controllers.ShujiaController{})
	beego.Router("/addbookcase/book/:BookId:int", &controllers.AddBookcaseController{})
	beego.Router("/deletebookcase/book/:BookId:int", &controllers.DeleteBookcaseController{})
	


	// beego.Router("/logout", &controllers.LogoutController{})

	beego.Router("/get/book/:BookId:int", &controllers.BookInfoController{})
	beego.Router("/get/book/:BookId:int/chapter/:ChapterNum:int", &controllers.ChapterController{})
	beego.Router("/book/search", &controllers.SearchController{})
	// beego.Router("/search/auto/:name:string", &controllers.SearchBookNameController{})
	// beego.Router("/get/classify", &controllers.BookInfoController{})
	// beego.Router("/get/classify/:name:string", &controllers.BookInfoController{})
	// beego.Router("/get/rank", &controllers.BookInfoController{})
	// beego.Router("/get/rank/:name:string", &controllers.BookInfoController{})
	// beego.Router("/get/history", &controllers.BookInfoController{})
	// beego.Router("/get/sanjiang", &controllers.BookInfoController{})
	// beego.Router("/get/strongrec", &controllers.BookInfoController{})
	beego.SetStaticPath("/file", "static")

}
