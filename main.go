package main

import (
	"database/sql"
	"fmt"
	"log"

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

	c.SecureJSON(200, gin.H{
		"records": records,
	})

}

type welfarepoint struct {
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

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	querypoint()
	// r := gin.Default()
	// // r.Use(Cors())
	// r.GET("/welfare/:name", homepage)
	// r.GET("/law/:company", query)
	// r.GET("/query/:company", query)
	// r.Run()
}

func querypoint() {
	str := `SELECT * FROM 104data.welfare`
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows.Scan())
	}
}
