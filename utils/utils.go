package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"
	"ystory/utils/mysql"
)

func LogAccess(userName, url string) {
	sql := `insert into access_history(user_name,access_url) values(?,?)`
	_, err := mysql.DBStory.Exec(sql, userName, url)
	if err != nil {
		panic(err)
	}
}

func LogBookAccess(userName string, bookId, chapterNum int) {
	sql := `insert into book_access_history(user_name,book_id,chapter_num) values(?,?,?)`
	_, err := mysql.DBStory.Exec(sql, userName, bookId, chapterNum)
	if err != nil {
		panic(err)
	}
}

func GetAnonymousName() (name string) {
	now := time.Now()
	name = now.Format("2006-01-02 15:04:05") + "." + strconv.FormatInt(time.Now().UnixNano(), 10)
	return
}

func Sign(userName string) (encodeString string) {
	s := sha1.New()
	io.WriteString(s, userName+"MDZZ")
	encodeString = fmt.Sprintf("%s%x", userName, s.Sum(nil))
	return
}

func UnSign(encodeString string) (userName string, err error) {
	if len(encodeString) < 40 {
		err = errors.New("encodeString error")
		return
	}
	userName = encodeString[0 : len(encodeString)-40]
	if Sign(userName) != encodeString {
		userName = ""
		err = errors.New("encodeString error")
	}
	return
}
