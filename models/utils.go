package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"ystory/utils"
	"ystory/utils/mysql"
)

func httpPost(url string, data *map[string]interface{}) (result *map[string]interface{}, err error) {
	application := "application/json"
	requestData, err := json.Marshal(data)
	if err != nil {
		return
	}
	resp, err := http.Post(url, application, bytes.NewBuffer(requestData))
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = new(map[string]interface{})
	err = json.Unmarshal(body, result)
	if err != nil {
		return
	}
	if (*result)["error_code"].(float64) != 0 {
		err = utils.SPE
		return
	}
	content := (*result)["content"].(map[string]interface{})
	result = &content
	return
}

func getChapterUrl(BookId, ChapterNum int) (url string) {
	url = "/get/book/" + strconv.Itoa(BookId) + "/chapter/" + strconv.Itoa(ChapterNum)
	return
}

func getCatalogUrl(BookId int) (url string) {
	url = "/get/book/" + strconv.Itoa(BookId)
	return
}

func getRomoveUrl(BookId int) (url string) {
	url = "/deletebookcase/book/" + strconv.Itoa(BookId)
	return
}

func getChapterFromSpider(BookId, ChapterNum int) (chapter *Chapter, err error) {
	sql := `select chapter_url from catalog where book_id= ? and chapter_num= ?`
	result, err := mysql.DBStory.Select(sql, BookId, ChapterNum)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	data := make(map[string]interface{})
	data["book_id"] = BookId
	data["chapter_url"] = string(result[0][`chapter_url`].([]byte))
	data["chapter_num"] = ChapterNum
	resp, err := httpPost(beego.AppConfig.String("spiderurl"), &data)
	if err != nil {
		return
	}
	chapter = new(Chapter)
	chapter.BookId = BookId
	chapter.BookName = (*resp)["book_name"].(string)
	chapter.ChapterNum = ChapterNum
	chapter.ChapterName = (*resp)["chapter_name"].(string)
	chapter.ChapterUrl = string(result[0][`chapter_url`].([]byte))
	chapter.Content = strings.Split(strings.Replace((*resp)["content"].(string), " ", "&nbsp;", -1), "\n")
	return
}

func getChapterFromMysql(BookId, ChapterNum int) (chapter *Chapter, err error) {
	sql := `SELECT id,book_id,book_name,chapter_num,chapter_name,chapter_url,content 
        FROM chapter where book_id = ? and chapter_num = ?`
	result, err := mysql.DBStory.Select(sql, BookId, ChapterNum)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	chapter = new(Chapter)
	chapter.Id = int(result[0][`id`].(int64))
	chapter.BookId = int(result[0][`book_id`].(int64))
	chapter.BookName = string(result[0][`book_name`].([]byte))
	chapter.ChapterNum = int(result[0][`chapter_num`].(int64))
	chapter.ChapterName = string(result[0][`chapter_name`].([]byte))
	chapter.ChapterUrl = string(result[0][`chapter_url`].([]byte))
	chapter.Content = strings.Split(strings.Replace(string(result[0][`content`].([]byte)), " ", "&nbsp;", -1), "\n")
	return
}

func getMaxChapterFromMysql(BookId int) (maxChapter int, err error) {
	sql := `select max(chapter_num) maxChapter from catalog where book_id = ?`
	result, err := mysql.DBStory.Select(sql, BookId)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	maxChapter = int(result[0][`maxChapter`].(int64))
	return
}

func getBookFromMysql(BookId int) (book *Book, err error) {
	sql := `select id,book_name,book_author,book_status,book_chapter_num,update_time,book_intro,book_img
        from book where id= ?`
	result, err := mysql.DBStory.Select(sql, BookId)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	book = new(Book)
	book.BookId = int(result[0][`id`].(int64))
	book.BookName = string(result[0][`book_name`].([]byte))
	book.BookAuthor = string(result[0][`book_author`].([]byte))
	book.BookStatus = string(result[0][`book_status`].([]byte))
	book.BookChapterNum = int(result[0][`book_chapter_num`].(int64))
	book.UpdateTime = string(result[0][`update_time`].([]byte))
	book.BookIntro = string(result[0][`book_intro`].([]byte))
	book.BookImg = string(result[0][`book_img`].([]byte))
	return
}

func getCatalogFromMysql(BookId int) (catalog *Catalog, err error) {
	book, err := getBookFromMysql(BookId)
	if err != nil {
		return
	}
	sql := `select chapter_name,chapter_num from catalog where book_id= ? 
        order by chapter_num`
	result, err := mysql.DBStory.Select(sql, BookId)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	catalog = new(Catalog)
	catalog.Book = book
	chapterNames := make([]string, len(result), len(result))
	chapterNums := make([]int, len(result), len(result))
	for index, item := range result {
		chapterNames[index] = string(item[`chapter_name`].([]byte))
		chapterNums[index] = int(item[`chapter_num`].(int64))
	}
	catalog.ChapterNames = &chapterNames
	catalog.ChapterNums = &chapterNums
	return
}

func getIndexItemFromMysql(BookType string) (indexItem *IndexItem, err error) {
	sql := `select id,book_name,book_author,book_type1 from book where book_type1=? limit 10`
	result, err := mysql.DBStory.Select(sql, BookType)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	indexItem = new(IndexItem)
	indexItem.Title = BookType
	for _, item := range result {
		var bookItem BookItem
		bookItem.New = false
		bookItem.BookName = string(item[`book_name`].([]byte))
		bookItem.BookAuthor = string(item[`book_author`].([]byte))
		bookItem.BookType1 = string(item[`book_type1`].([]byte))
		bookId := int(item[`id`].(int64))
		bookItem.BookUrl = getCatalogUrl(bookId)
		indexItem.BookItems = append(indexItem.BookItems, bookItem)
	}
	return
}

func getSearchRresult(serarchKey string) (indexItem *IndexItem, err error) {
	serarchKey = fmt.Sprintf("%%%s%%", serarchKey)
	sql := `select id,book_name,book_author,book_type1 from book where book_name like ? limit 10`
	result, err := mysql.DBStory.Select(sql, serarchKey)
	if err != nil {
		return
	}
	indexItem = new(IndexItem)
	indexItem.Title = "搜索结果"
	for _, item := range result {
		var bookItem BookItem
		bookItem.New = false
		bookItem.BookName = string(item[`book_name`].([]byte))
		bookItem.BookAuthor = string(item[`book_author`].([]byte))
		bookItem.BookType1 = string(item[`book_type1`].([]byte))
		bookId := int(item[`id`].(int64))
		bookItem.BookUrl = getCatalogUrl(bookId)
		indexItem.BookItems = append(indexItem.BookItems, bookItem)
	}
	return
}

func checkUpdate(lastAccessTime string, bookId int) (chapterNum int) {
	if lastAccessTime == "" {
		return -1
	}
	sql := `select chapter_num as chapter_num from catalog where book_id =?
		and create_time >? order by chapter_num limit 1`
	result, err := mysql.DBStory.Select(sql, bookId, lastAccessTime)
	if err != nil {
		return -1
	}
	if len(result) == 0 {
		return -1
	}
	return int(result[0][`chapter_num`].(int64))
}

func getBookAccessHistory(userName string) (indexItem *IndexItem, err error) {
	sql := `select a.id id,book_name,book_author,book_type1,b.ct access_time from book a inner join 
		(select book_id,max(create_time)as ct from book_access_history where user_name = ? group by book_id order by ct desc limit 20) b 
		on a.id=b.book_id;`
	result, err := mysql.DBStory.Select(sql, userName)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = utils.NFD
		return
	}
	indexItem = new(IndexItem)
	indexItem.Title = "阅读历史"
	for _, item := range result {
		var bookItem BookItem
		bookItem.BookName = string(item[`book_name`].([]byte))
		bookItem.BookAuthor = string(item[`book_author`].([]byte))
		bookItem.BookType1 = string(item[`book_type1`].([]byte))
		bookId := int(item[`id`].(int64))
		chapterNum := checkUpdate(string(item[`access_time`].([]byte)), bookId)
		if chapterNum > 0 {
			bookItem.New = true
		}
		bookItem.BookUrl = getCatalogUrl(bookId)
		indexItem.BookItems = append(indexItem.BookItems, bookItem)
	}
	return
}

func getLastReadHistory(userName string, bookId int) (chapterNum int, accessTime string) {
	sql := `select chapter_num,create_time access_time from book_access_history where user_name=? and book_id=? and chapter_num!=0 order by id desc limit 1`
	result, err := mysql.DBStory.Select(sql, userName, bookId)
	if err != nil {
		panic(err)
		return 0, ""
	}
	if len(result) == 0 {
		return 0, ""
	}
	accessTime = string(result[0][`access_time`].([]byte))
	chapterNum = int(result[0][`chapter_num`].(int64))
	return
}

func CheckPassword(username, password string) string {
	sql := `select id from user where username=? and password=?`
	result, err := mysql.DBStory.Select(sql, username, password)
	if err != nil {
		return "登录失败"
	}
	if len(result) == 0 {
		return "登录失败"
	}
	return "success"
}

func Register(username, password, email string) string {
	sql := `insert into  user(username,password,email) values(?,?,?)`
	_, err := mysql.DBStory.Exec(sql, username, password, email)
	if err != nil {
		return "注册失败"
	}
	return "success"
}

type ShujiaItem struct {
	New        bool   `json:"new"`
	BookType1  string `json:"book_type1"`
	BookAuthor string `json:"book_author"`
	BookName   string `json:"book_name"`
	BookUrl    string `json:"book_url"`
	RemoveUrl  string `json:"RemoveUrl"`
}

type Shujia struct {
	BookName   string `json:"book_name"`
	Title     string     `json:"title"`
	TitleUrl  string     `json:"title_url"`
	ShujiaItems []ShujiaItem `json:"book_item"`
	More      bool
	MoreUrl   string
}

func GetShujia(username string) (shujia *Shujia, err error) {
	sql := `select a.id id,book_name,book_author,book_type1 from book a inner join 
		shujia b on a.id = b.book_id where b.user_name = ?`
	result, err := mysql.DBStory.Select(sql, username)
	if err != nil {
		return
	}
	shujia = new(Shujia)
	shujia.BookName = "我的书架"
	shujia.Title = "我的书架"
	for _, item := range result {
		var shujiaItem ShujiaItem
		shujiaItem.BookName = string(item[`book_name`].([]byte))
		bookId := int(item[`id`].(int64))
		shujiaItem.BookUrl = getCatalogUrl(bookId)
		shujiaItem.RemoveUrl = getRomoveUrl(bookId)
		shujia.ShujiaItems = append(shujia.ShujiaItems, shujiaItem)
	}
	return
}

func AddBookcase(username string,bookId int)  string{
	sql := `insert into shujia(user_name,book_id) values(?,?)`
	_, err := mysql.DBStory.Exec(sql, username, bookId)
	if err != nil {
		return "加入失败"
	}
	return "success"
}

func DeleteBookcase(username string,bookId int)  error{
	sql := `delete from shujia where user_name = ? and book_id = ?`
	_, err := mysql.DBStory.Exec(sql, username, bookId)
	return err
}

type Item struct {
	Url  template.URL
	Name string
}
