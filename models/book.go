package models

type bookItem struct {
	ChapterNum int    `json:"chapter_num"`
	BookName   string `json:"chapter_name"`
	BookUrl    string `json:"book_url"`
}

type Book struct {
	BookId         int    `json:"book_id"`
	BookName       string `json:"book_name"`
	BookAuthor     string `json:"book_author"`
	BookStatus     string `json:"book_status"`
	BookChapterNum int    `json:"book_chapter_num"`
	UpdateTime     string `json:"update_time"`
	BookIntro      string `json:"book_intro"`
	BookImg        string `json:"book_img"`
}

