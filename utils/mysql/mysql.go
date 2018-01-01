package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/astaxie/beego"

type DB struct {
	*sql.DB
}

var (
	DBStory DB
)

func init() {
	db1, err := sql.Open("mysql", beego.AppConfig.String("mysqladdr"))
	if err != nil {
		panic(err)
	}
	DBStory = DB{db1}
}

func (db DB) Select(sql string, values ...interface{}) (result []map[string]interface{}, err error) {
	rows, err := db.Query(sql, values...)
	if err != nil {
		return
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	results := make([]interface{}, len(columns))
	for i := range scanArgs {
		scanArgs[i] = &results[i]
	}
	for rows.Next() {
		Row := map[string]interface{}{}
		if err = rows.Scan(scanArgs...); err != nil {
			return
		}
		for i := range results {
			Row[columns[i]] = results[i]
		}
		result = append(result, Row)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (db DB) Exec(sql string, values ...interface{}) (res sql.Result, err error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	res, err = stmt.Exec(values...)
	return
}
