package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"ystory/models"
	"ystory/utils"
)

type ChapterController struct {
	beego.Controller
}

type BookInfoController struct {
	beego.Controller
}

type IndexController struct {
	beego.Controller
}

type SearchController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	userName := this.Ctx.Input.GetData("userName").(string)
	Index, err := models.GetIndex(userName)
	utils.CkeckError(this, err)
	this.Data["Data"] = Index
	this.TplName = "index.tpl"

}

func (this *ChapterController) Get() {
	BookId, err := strconv.Atoi(this.Ctx.Input.Param(":BookId"))
	utils.CkeckError(this, err)
	ChapterNum, err := strconv.Atoi(this.Ctx.Input.Param(":ChapterNum"))
	utils.CkeckError(this, err)
	userName := this.Ctx.Input.GetData("userName").(string)
	utils.LogBookAccess(userName, BookId, ChapterNum)
	Chapter, err := models.GetChapter(BookId, ChapterNum)
	utils.CkeckError(this, err)
	RenderChapter, err := Chapter.GetRenderChapter()
	utils.CkeckError(this, err)
	this.Data["Data"] = RenderChapter
	this.TplName = "chapter.tpl"
}

func (this *BookInfoController) Get() {
	BookId, err := strconv.Atoi(this.Ctx.Input.Param(":BookId"))
	utils.CkeckError(this, err)
	userName := this.Ctx.Input.GetData("userName").(string)
	utils.LogBookAccess(userName, BookId, 0)
	Catalog, err := models.GetCatalog(BookId)
	utils.CkeckError(this, err)
	RenderCatalog, err := Catalog.GetRenderCatalog(userName)
	utils.CkeckError(this, err)
	this.Data["Data"] = RenderCatalog
	this.TplName = "catalog.tpl"
}

func (this *SearchController) Get() {
	searchKey := this.GetString("searchkey")
	Index, err := models.GetSearchRresult(searchKey)
	utils.CkeckError(this, err)
	this.Data["Data"] = Index
	this.TplName = "index.tpl"
}
