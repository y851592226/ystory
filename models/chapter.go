package models

type Chapter struct {
	Id          int      `json:"id"`
	BookId      int      `json:"book_id"`
	BookName    string   `json:"book_name"`
	ChapterNum  int      `json:"chapter_num"`
	ChapterName string   `json:"chapter_name"`
	ChapterUrl  string   `json:"chapter_url"`
	Content     []string `json:"content"`
}

type RenderChapter struct {
	*Chapter
	Title            string
	CatalogUrl       string
	ChapterUrlBefore string
	ChapterUrlAfter  string
}

func GetChapter(BookId, ChapterNum int) (chapter *Chapter, err error) {
	chapter, err = getChapterFromMysql(BookId, ChapterNum)
	if err == nil {
		return
	}
	chapter, err = getChapterFromSpider(BookId, ChapterNum)
	return
}

func (chapter *Chapter) GetMaxChapter() (maxChapter int, err error) {
	BookId := chapter.BookId
	maxChapter, err = getMaxChapterFromMysql(BookId)
	return
}

func (chapter *Chapter) GetRenderChapter() (renderChapter *RenderChapter, err error) {
	maxChapter, err := chapter.GetMaxChapter()
	if err != nil {
		return
	}
	renderChapter = new(RenderChapter)
	renderChapter.Chapter = chapter
	renderChapter.Title = chapter.BookName
	renderChapter.CatalogUrl = getCatalogUrl(chapter.BookId)
	if chapter.ChapterNum > 1 {
		renderChapter.ChapterUrlBefore = getChapterUrl(chapter.BookId, chapter.ChapterNum-1)
	} else {
		renderChapter.ChapterUrlBefore = renderChapter.CatalogUrl
	}
	if chapter.ChapterNum < maxChapter {
		renderChapter.ChapterUrlAfter = getChapterUrl(chapter.BookId, chapter.ChapterNum+1)
	} else {
		renderChapter.ChapterUrlAfter = renderChapter.CatalogUrl
	}
	return
}
