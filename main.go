package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"thesis/welfare"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var err error
var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:qaz741236985@tcp(localhost:3306)/104data?charset=utf8")
}

type Record struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	Publicdate  string `json:"publicdate"`
	Company     string `json:"company"`
	Dealdate    string `json:"dealdate"`
	Govnumber   string `json:"govnumber"`
	Law         string `json:"law"`
	Description string `json:"description"`
	Ps          string `json:"ps"`
}

func homepage(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "hello" + name,
	})
}

func query(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	company := c.Param("company")
	str := `SELECT * FROM 104data.illegal_record where company like '%` + company + `%'`
	fmt.Println(str)
	row, err := db.Query(str)
	defer row.Close()
	if err != nil {
		log.Fatal(err)
	}

	records := []Record{}

	for row.Next() {
		var record Record
		row.Scan(&record.ID, &record.Location, &record.Publicdate, &record.Company, &record.Dealdate, &record.Govnumber, &record.Law, &record.Description, &record.Ps)
		records = append(records, record)
	}
	if err = row.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("")
	fmt.Println(records)
	c.JSON(200, gin.H{
		"records": records,
	})

}

func makescore(c *gin.Context) {
	wstring := c.Param("welfare")
	var w welfare.Welfarepoint
	w.Match(wstring)
	score := w.Wtoi()
	c.JSON(200, gin.H{
		"message": score,
	})
}

func salary(c *gin.Context) {
	salary := c.Param("salary")
	c.JSON(200, gin.H{
		"salary": salary,
	})
}

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	// dividpoint := querypoint()
	// fmt.Println(dividpoint)

	r := gin.Default()
	r.GET("/welfare/:welfare", makescore)
	r.GET("/law/:company", query)
	r.GET("/query/:company", query)
	r.GET("/salary/:salary", salary)
	r.Run()
}

func querypoint() []int {
	start := time.Now()
	str := `SELECT * FROM 104data.welfare`
	rows, err := db.Query(str)
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

	//write score to db
	// vals := []interface{}{}
	// sqlStr := `insert into welfarepoint(Custno, point) VALUES`
	// for i, el := range welfarepoint {
	// 	p := el.wtoi()
	// 	vals = append(vals, el.Company, p)
	// 	sqlStr += `(?,?),`
	// 	if i%5000 == 0 || i == len(welfarepoint)-1 {
	// 		sqlstart := time.Now()
	// 		sqlStr = sqlStr[0 : len(sqlStr)-1]
	// 		sqlStr = sqlStr + `ON DUPLICATE KEY UPDATE point = values(point)`
	// 		stmt, err := db.Prepare(sqlStr)
	// 		if err != nil {
	// 			fmt.Println("prepare error ", err)
	// 		}
	// 		_, err = stmt.Exec(vals...)
	// 		if err != nil {
	// 			fmt.Println("exec error", err)
	// 		}
	// 		stmt.Close()
	// 		sqlend := time.Now()
	// 		fmt.Println(i, "complete", sqlend.Sub(sqlstart).Seconds())
	// 		sqlStr = `insert into welfarepoint(Custno, point) VALUES`
	// 		vals = []interface{}{}
	// 	}
	// }
	// end write to db

	point := []int{}
	// a := []int{}
	for _, el := range welfarepoint {
		w := el.Wtoi()

		point = append(point, w)

	}

	sort.Ints(point)
	// fmt.Println(len(a))
	dividindex := len(point) / 10
	fmt.Println("slice index", dividindex)
	fmt.Println(point[dividindex], point[dividindex*2], point[dividindex*3], point[dividindex*4], point[dividindex*5],
		point[dividindex*6], point[dividindex*7], point[dividindex*8], point[dividindex*9], point[len(point)-1],
	)
	var result []int
	for i := 0; i < 10; i++ {
		result = append(result, point[dividindex*i])
	}

	end := time.Now()
	fmt.Println("end time: ", end.Sub(start).Seconds())
	return result
}
