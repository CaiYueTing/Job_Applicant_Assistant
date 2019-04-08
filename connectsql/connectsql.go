package connectsql

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"thesis/welfare"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var err error
var Db = &sql.DB{}
var Localdb = &sql.DB{}

func init() {
	// Localdb, _ = sql.Open("mysql", "root:qaz741236985@tcp(localhost:3306)/104data?charset=utf8")
	Localdb, _ = sql.Open("mysql", "jobhelper:qaz741236985@tcp(jobhelper.ck1vznvje3ei.ap-northeast-2.rds.amazonaws.com:3306)/jobdata?charset=utf8")
	if err = Localdb.Ping(); err != nil {
		log.Fatal(err)
	}
}

func getwelfare() []welfare.Welfarepoint {
	str := `SELECT * FROM welfare`
	rows, err := Localdb.Query(str)
	if err != nil {
		log.Fatal(err)
	}

	welfarepoint := []welfare.Welfarepoint{}

	for rows.Next() {
		var w welfare.Welfarepoint
		rows.Scan(
			&w.Company,
			&w.Three, &w.Yearend, &w.Bitrh, &w.Marry, &w.Maternity,
			&w.Patent, &w.Longterm, &w.Insurance, &w.Stock, &w.Annual,
			&w.Attendance, &w.Performance, &w.Travel, &w.Consolation, &w.Health,
			&w.Flexible, &w.Paternityleave, &w.Travelleave, &w.Physiologyleave, &w.Fullpaysickleave,
			&w.Dorm, &w.Restaurant, &w.Childcare, &w.Transport, &w.Servemeals,
			&w.Snack, &w.Afternoon, &w.Gym, &w.Education, &w.Tail,
			&w.Employeetravel, &w.Society, &w.Overtime, &w.Shift, &w.Permanent,
		)

		welfarepoint = append(welfarepoint, w)
	}

	rows.Close()
	return welfarepoint

}

func writepoint() {
	w := getwelfare()

	vals := []interface{}{}
	sqlStr := `insert into welfarepoint(Custno, point) VALUES`
	for i, el := range w {
		p := el.Wtoi()
		vals = append(vals, el.Company, p)
		sqlStr += `(?,?),`
		if i%5000 == 0 || i == len(w)-1 {
			sqlstart := time.Now()
			sqlStr = sqlStr[0 : len(sqlStr)-1]
			sqlStr = sqlStr + `ON DUPLICATE KEY UPDATE point = values(point)`
			stmt, err := Localdb.Prepare(sqlStr)
			if err != nil {
				fmt.Println("prepare error ", err)
			}
			_, err = stmt.Exec(vals...)
			if err != nil {
				fmt.Println("exec error", err)
			}
			stmt.Close()
			sqlend := time.Now()
			fmt.Println(i, "complete", sqlend.Sub(sqlstart).Seconds())
			sqlStr = `insert into welfarepoint(Custno, point) VALUES`
			vals = []interface{}{}
		}
	}
}

func Querypoint() []int {
	welfarepoint := getwelfare()
	point := []int{}
	for _, el := range welfarepoint {
		w := el.Wtoi()
		if w > 0 {
			point = append(point, w)
		}
	}
	sort.Ints(point)
	dividindex := len(point) / 10
	fmt.Println("slice index", dividindex)
	fmt.Println(point[dividindex], point[dividindex*2], point[dividindex*3], point[dividindex*4], point[dividindex*5],
		point[dividindex*6], point[dividindex*7], point[dividindex*8], point[dividindex*9], point[len(point)-1],
	)
	var result []int
	for i := 1; i < 10; i++ {
		result = append(result, point[dividindex*i])
	}
	return result
}

// func Querystr(str string) (interface{}, error) {
// 	res, err := localdb.Exec(str)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("in package :", res)
// 	return res, err
// }
