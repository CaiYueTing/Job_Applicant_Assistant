package main

import (
	"database/sql"
	"fmt"
	"log"
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

type Welfarepoint struct {
	Company          string `json:"company"`
	Three            bool   `json:"three"`
	Yearend          bool   `json:"yearend"`
	Bitrh            bool   `json:"bitrh"`
	Marry            bool   `json:"marry"`
	Maternity        bool   `json:"maternity"`
	Patent           bool   `json:"patent"`
	Longterm         bool   `json:"longterm"`
	Insurance        bool   `json:"insurance"`
	Stock            bool   `json:"stock"`
	Annual           bool   `json:"annual"`
	Attendance       bool   `json:"attendance"`
	Performance      bool   `json:"performance"`
	Travel           bool   `json:"travel"`
	Consolation      bool   `json:"consolation"`
	Health           bool   `json:"health"`
	Flexible         bool   `json:"flexible"`
	Paternityleave   bool   `json:"paternityleave"`
	Travelleave      bool   `json:"travelleave"`
	Physiologyleave  bool   `json:"Physiologyleave"`
	Fullpaysickleave bool   `json:"fullpaysickleave"`
	Dorm             bool   `json:"dorm"`
	Restaurant       bool   `json:"restaurant"`
	Childcare        bool   `json:"childcare"`
	Transport        bool   `json:"transport"`
	Servemeals       bool   `json:"servemeals"`
	Snack            bool   `json:"snack"`
	Afternoon        bool   `json:"afternoon"`
	Gym              bool   `json:"gym"`
	Education        bool   `json:"education"`
	Tail             bool   `json:"tail"`
	Employeetravel   bool   `json:"employeetravel"`
	Society          bool   `json:"society"`
	Overtime         bool   `json:"overtime"`
	Shift            bool   `json:"shift"`
	Permanent        bool   `json:"permanent"`
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
	c.SecureJSON(200, gin.H{
		"records": records,
	})

}

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	querypoint()
	// r := gin.Default()
	// r.GET("/welfare/:name", homepage)
	// r.GET("/law/:company", query)
	// r.GET("/query/:company", query)
	// r.Run()
}

func querypoint() {
	start := time.Now()
	str := `SELECT * FROM 104data.welfare`
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}

	welfarepoint := []Welfarepoint{}

	for rows.Next() {
		var w Welfarepoint
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

	// point := []int{}
	// a := []int{}
	vals := []interface{}{}
	sqlStr := `insert into welfarepoint(Custno, point) VALUES`
	for i, el := range welfarepoint {
		p := el.wtoi()
		vals = append(vals, el.Company, p)
		sqlStr += `(?,?),`
		if i%5000 == 0 || i == len(welfarepoint)-1 {
			sqlstart := time.Now()
			sqlStr = sqlStr[0 : len(sqlStr)-1]
			sqlStr = sqlStr + `ON DUPLICATE KEY UPDATE point = values(point)`  
			stmt, err := db.Prepare(sqlStr)
			if err != nil {
				fmt.Println("prepare error ", err)
			}
			_, err = stmt.Exec(vals...)
			if err != nil {
				fmt.Println("exec error", err)
			}
			stmt.Close()
			sqlend := time.Now()
			fmt.Println(i,"complete", sqlend.Sub(sqlstart).Seconds())
			sqlStr = `insert into welfarepoint(Custno, point) VALUES`
			vals = []interface{}{}
		}
	}
	// fmt.Println(c)
	// sort.Ints(point)
	// fmt.Println(len(a))
	// dividindex := len(point) / 10
	// fmt.Println("slice index", dividindex)
	// fmt.Println(point[dividindex], point[dividindex*2], point[dividindex*3], point[dividindex*4], point[dividindex*5],
	// 	point[dividindex*6], point[dividindex*7], point[dividindex*8], point[dividindex*9], point[len(point)-1],
	// )

	end := time.Now()
	fmt.Println("end time: ", end.Sub(start).Seconds())
}

func (w Welfarepoint) wtoi() int {
	result :=
		btou(w.Three) + btou(w.Yearend) + btou(w.Bitrh) + btou(w.Marry) + btou(w.Maternity) +
			btou(w.Patent) + btou(w.Longterm) + btou(w.Insurance) + btou(w.Stock) + btou(w.Annual) +
			btou(w.Attendance) + btou(w.Performance) + btou(w.Travel) + btou(w.Consolation) + btou(w.Health) +
			btou(w.Flexible) + btou(w.Paternityleave) + btou(w.Travelleave) + btou(w.Physiologyleave) + btou(w.Fullpaysickleave) +
			btou(w.Dorm) + btou(w.Restaurant) + btou(w.Childcare) + btou(w.Transport) + btou(w.Servemeals) +
			btou(w.Snack) + btou(w.Afternoon) + btou(w.Gym) + btou(w.Education) + btou(w.Tail) +
			btou(w.Employeetravel) + btou(w.Society) + btou(w.Overtime) + btou(w.Shift) + btou(w.Permanent)
	return result
}

func btou(b bool) int {
	if b {
		return 1
	}
	return 0
}
