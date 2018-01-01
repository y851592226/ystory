package models

type chapterItem struct {
	New         bool   `json:"new"`
	ChapterNum  int    `json:"chapter_num"`
	ChapterName string `json:"chapter_name"`
	ChapterUrl  string `json:"chapter_url"`
}
type Catalog struct {
	*Book
	ChapterNames *[]string `json:"chapter_names"`
	ChapterNums  *[]int    `json:"chapterNums"`
}

type RenderCatalog struct {
	*Catalog
	Title               string
	LastChapterUrl      string
	LastChapterName     string
	LastReadChapterUrl  string
	LastReadChapterName string
	Chapters            *[]*chapterItem
}

func GetCatalog(BookId int) (catalog *Catalog, err error) {
	catalog, err = getCatalogFromMysql(BookId)
	return
}

func (catalog *Catalog) GetRenderCatalog(userName string) (renderCatalog *RenderCatalog, err error) {
	renderCatalog = new(RenderCatalog)
	renderCatalog.Title = catalog.BookName
	chapters := make([]*chapterItem, len(*catalog.ChapterNums), len(*catalog.ChapterNums))
	lastReadChapterNum, accessTime := getLastReadHistory(userName, catalog.BookId)
	if lastReadChapterNum != 0 {
		renderCatalog.LastReadChapterUrl = getChapterUrl(catalog.BookId, (*catalog.ChapterNums)[lastReadChapterNum-1])
		renderCatalog.LastReadChapterName = (*catalog.ChapterNames)[lastReadChapterNum-1]
	}
	chapterNum := checkUpdate(accessTime, catalog.BookId)
	for index, _ := range *catalog.ChapterNums {
		chapterItem := new(chapterItem)
		if chapterNum > 0 && index >= chapterNum-1 {
			chapterItem.New = true
		}
		chapterItem.ChapterNum = (*catalog.ChapterNums)[index]
		chapterItem.ChapterName = (*catalog.ChapterNames)[index]
		chapterItem.ChapterUrl = getChapterUrl(catalog.BookId, (*catalog.ChapterNums)[index])
		chapters[index] = chapterItem
	}
	renderCatalog.LastChapterUrl = getChapterUrl(catalog.BookId, (*catalog.ChapterNums)[len(*catalog.ChapterNums)-1])
	renderCatalog.LastChapterName = (*catalog.ChapterNames)[len(*catalog.ChapterNums)-1]
	renderCatalog.Catalog = catalog
	renderCatalog.Chapters = &chapters
	return
}
