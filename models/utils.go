package models

import (
    "html/template"

)

type Item struct {
    Url template.URL
    Name string
}

type Topbar struct {
    Item [3]*Item
}

type Subnavbar struct {
    Item [5]*Item
}

type Searchbar struct {
    Tip1 string
    Tip2 string
}

type Chapter struct {
    Book_id      int    `json:"book_id"`
    Book_name    string `json:"book_name"`
    Chapter_num  int    `json:"chapter_num"`
    Chapter_name string `json:"chapter_name"`
    Chapter_url  string `json:"chapter_url"`
    Href_mulu    string `json:"href_mulu"`
    Href_before  string `json:"href_before"`
    Href_after   string `json:"href_after"`
    Content      []string `json:"content"`
}

type Catalog struct {
    Chapters []Chapter `json:"chapters"`
}

type Book struct {
    Book_id           int      `json:"book_id"`
    Book_name         string   `json:"book_name"`
    Book_url          string   `json:"book_url"`
    Book_author       string   `json:"book_author"`
    Book_intro        string   `json:"book_intro"`
    Book_type1        string   `json:"book_type1"`
    Book_type2        string   `json:"book_type2"`
    Book_img          string   `json:"book_img"`
    Last_read_chapter *Chapter `json:"last_read_chapter"`
    Latest_chapter    *Chapter `json:"latest_chapter"`
    Catalog           *Catalog `json:"catalog"`
}

type Book_list struct {
    Title string `json:"Title"`
    Books []Book `json:"books"`
}

type Reading_history struct {
    Books []Book `json:"books"`
}

type Index_list struct {
    Index_lists []Index_list `json:"index_lists"`
}