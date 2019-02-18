package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"thesis/welfare"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var err error
var db = &sql.DB{}
var p []int

func init() {
	db, _ = sql.Open("mysql", "root:qaz741236985@tcp(localhost:3306)/104data?charset=utf8")

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	p = querypoint()
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

type Analysis struct {
	Category string `json:"category"`
	Target   struct {
		Industry []Analyresult `json:"industry"`
		Exp      []Analyresult `json:"exp"`
		District []Analyresult `json:"district"`
	} `json:"target"`
}

type Analyresult struct {
	Description string `json:"description"`
	Right       int    `json:"right"`
	Left        int    `json:"left"`
	Middle      int    `json:"middle"`
	Average     int    `json:"average"`
}

func lawsearch(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	company := c.Param("company")
	str := `SELECT * FROM 104data.illegal_record where company like '%` + company + `%'`
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
		"dd":      p,
	})
}

func salary(c *gin.Context) {
	salary := c.Param("salary")
	c.JSON(200, gin.H{
		"salary": salary,
	})
}

func postscore(c *gin.Context) {
	wstring := c.PostForm("wdata")
	var w welfare.Welfarepoint
	w.Match(wstring)
	score := w.Wtoi()
	c.JSON(http.StatusOK, gin.H{
		"message": score,
		"dd":      p,
	})
}

func category(c *gin.Context) {
	cstring := c.PostForm("cdata")
	as := strings.Split(cstring, "、")
	result := []string{}
	for _, v := range as {
		strings.Replace(v, "、", "", -1)
		str := `SELECT CategoryId FROM 104data.jobcategory where category = '` + v + `'`
		row, err := db.Query(str)
		if err != nil {
			log.Fatal(err)
		}
		for row.Next() {
			var c string
			row.Scan(&c)
			result = append(result, c)
		}
	}
	r := []Analysis{}
	for _, v := range result {
		var analysis Analysis
		for i := 0; i < 3; i++ {
			analystr := ""
			if i == 0 {
				analystr = "SELECT jobcategory.category, industry.industry, leftnum, rightnum, middlevalue, average FROM 104data.cateindustry, industry, jobcategory where cateindustry.industry = industry.IndustryId and cateindustry.categoryId = jobcategory.CategoryId and cateindustry.categoryId=" + v
				rows, err := db.Query(analystr)
				if err != nil {
					log.Fatal(err)
				}

				for rows.Next() {
					var ar Analyresult
					rows.Scan(&analysis.Category, &ar.Description, &ar.Left, &ar.Right, &ar.Middle, &ar.Average)
					analysis.Target.Industry = append(analysis.Target.Industry, ar)
				}
				rows.Close()
			}
			if i == 1 {
				analystr = `SELECT  district, leftnum, rightnum, middlevalue, average FROM 104data.catedistrict, jobcategory where catedistrict.categoryId = jobcategory.CategoryId and catedistrict.categoryId =` + v
				rows, err := db.Query(analystr)
				if err != nil {
					log.Fatal(err)
				}

				for rows.Next() {
					var ar Analyresult
					rows.Scan(&ar.Description, &ar.Left, &ar.Right, &ar.Middle, &ar.Average)
					analysis.Target.District = append(analysis.Target.District, ar)
				}
				rows.Close()
			}
			if i == 2 {
				analystr = `SELECT   exp, leftnum, rightnum, middlevalue, average FROM 104data.cateexp, jobcategory where cateexp.categoryId = jobcategory.CategoryId and cateexp.categoryId =` + v
				rows, err := db.Query(analystr)
				if err != nil {
					log.Fatal(err)
				}

				for rows.Next() {
					var ar Analyresult
					rows.Scan(&ar.Description, &ar.Left, &ar.Right, &ar.Middle, &ar.Average)
					analysis.Target.Exp = append(analysis.Target.Exp, ar)
				}
				rows.Close()
			}
			fmt.Println(analysis)
		}
		r = append(r, analysis)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": r,
	})
}

func main() {

	r := gin.Default()

	cardAPI := r.Group("/card")
	{
		cardAPI.POST("/welfare", postscore)
		cardAPI.GET("/law/:company", lawsearch)
		cardAPI.GET("/salary/:salary", salary)
		cardAPI.POST("/category", category)
	}

	r.Run()
}

func getwelfare() []welfare.Welfarepoint {
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
			fmt.Println(i, "complete", sqlend.Sub(sqlstart).Seconds())
			sqlStr = `insert into welfarepoint(Custno, point) VALUES`
			vals = []interface{}{}
		}
	}
}

func querypoint() []int {
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
