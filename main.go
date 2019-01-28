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
	id          int
	location    string
	publicdate  string
	company     string
	dealdate    string
	govnumber   string
	law         string
	description string
	ps          string
}

func homepage(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "hello" + name,
	})
}

func lawcount(c *gin.Context) {
	count := c.Param("company")
	c.JSON(200, gin.H{
		"message": len(count),
	})
}

func query(c *gin.Context) {
	company := c.Param("company")
	str := `SELECT * FROM 104data.illegal_record where company like '%` + company + `%'`
	fmt.Println(str)
	row, err := db.Query(str)
	defer row.Close()
	if err != nil {
		panic(err)
	}

	records := []Record{}

	for row.Next() {
		var record Record
		row.Scan(&record.id, &record.location, &record.publicdate, &record.company, &record.dealdate, &record.govnumber, &record.law, &record.description, &record.ps)
		records = append(records, record)
	}
	fmt.Println(records)
	if err = row.Err(); err != nil {
		log.Fatalln(err)
	}
	c.JSON(200, gin.H{
		"records": records,
	})
}

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.GET("/welfare/:name", homepage)
	r.GET("/law/:company", lawcount)
	r.GET("/query/:company", query)
	r.Run()
}
