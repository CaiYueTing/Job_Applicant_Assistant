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

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	// r.Use(Cors())
	r.GET("/welfare/:name", homepage)
	r.GET("/law/:company", query)
	r.GET("/query/:company", query)
	r.Run()
}
