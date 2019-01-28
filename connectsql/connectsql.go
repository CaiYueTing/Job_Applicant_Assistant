package connectsql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var err error
var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:qaz741236985@tcp(localhost:3306)/104data?charset=utf8")
}

func Querystr(str string) (interface{}, error) {
	res, err := db.Exec(str)
	if err != nil {
		panic(err)
	}
	fmt.Println("in package :", res)
	return res, err
}
