package models


type BookItem struct {
	New        bool   `json:"new"`
	BookType1  string `json:"book_type1"`
	BookAuthor string `json:"book_author"`
	BookName   string `json:"chapter_name"`
	BookUrl    string `json:"book_url"`
}

type IndexItem struct {
	Title     string     `json:"title"`
	TitleUrl  string     `json:"title_url"`
	BookItems []BookItem `json:"book_item"`
	More      bool
	MoreUrl   string
}

type Index struct {
	Title      string      `json:"title"`
	IndexItems []IndexItem `json:"index_item"`
}

func GetIndex(userName string) (index *Index, err error) {
	index = new(Index)
	index.Title = `小说阅读`
	bookTypes := []string{`游戏`, `历史`, `武侠`}

	IndexItem, err := getBookAccessHistory(userName)
	if err == nil {
		index.IndexItems = append(index.IndexItems, *IndexItem)
	}
	err = nil
	for _, bookType := range bookTypes {
		IndexItem, err := getIndexItemFromMysql(bookType)
		if err != nil {
			return index, err
		}
		index.IndexItems = append(index.IndexItems, *IndexItem)
	}
	return
}
