package api

import (
	"log"
	"net/http"
	"strings"
	"thesis/connectsql"
	"thesis/crawler"
	"thesis/welfare"

	"github.com/gin-gonic/gin"
)

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

var thesisdb = connectsql.Localdb
var p []int
var err error

func init() {
	p = connectsql.Querypoint()
}

func dealstring(str string) string {
	as := []string{"股份有限公司", "有限公司", "工作室", "事務所", "補習班", "便利店"}
	s := str
	if strings.Contains(s, "(") && strings.Contains(s, ")") {
		_i := strings.Index(s, "(")
		_ii := strings.Index(s, ")")
		if _i < _ii {
			_ai := strings.Split(s, "(")
			_aii := strings.Split(s, ")")
			s = _ai[0] + _aii[1]
		}
	}

	for _, v := range as {
		i := strings.Index(s, v)
		if i > 0 {
			s = s[:i]
			break
		}
	}

	if strings.Contains(s, "_") {
		a := strings.Split(s, "_")
		s = a[len(a)-1]
	}
	return s
}

func Lawsearch(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	company := c.Param("company")
	s := company
	s = dealstring(s)

	str := `SELECT * FROM illegal_record where company like '%` + s + `%'`
	row, err := thesisdb.Query(str)
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

func Qollie(c *gin.Context) {
	company := c.Param("company")
	s := company
	s = dealstring(s)
	ch := make(chan crawler.Comment)
	var qollie crawler.Comment
	go func() {
		qollie := crawler.CrawlQollie(s)
		ch <- qollie
	}()

	qollie = <-ch
	c.JSON(200, gin.H{
		"qollie": qollie,
	})
}

func Makescore(c *gin.Context) {
	wstring := c.Param("welfare")
	var w welfare.Welfarepoint
	w.Match(wstring)
	score := w.Wtoi()
	c.JSON(200, gin.H{
		"message": score,
		"dd":      p,
	})
}

func Salary(c *gin.Context) {
	salary := c.Param("salary")
	c.JSON(200, gin.H{
		"salary": salary,
	})
}

func Postscore(c *gin.Context) {
	wstring := c.PostForm("wdata")
	var w welfare.Welfarepoint
	r := w.Match2(wstring)
	score := w.Wtoi()
	c.JSON(http.StatusOK, gin.H{
		"message": score,
		"dd":      p,
		"r":       r,
	})
}

func Category(c *gin.Context) {
	cstring := c.PostForm("cdata")
	as := strings.Split(cstring, "、")
	result := []string{}
	for _, v := range as {
		strings.Replace(v, "、", "", -1)
		str := `SELECT CategoryId FROM jobcategory where category = '` + v + `' and hiding like'%否%'`
		row, err := thesisdb.Query(str)
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
	for k, v := range result {
		var analysis Analysis
		for i := 0; i < 3; i++ {
			analystr := ""
			if i == 0 {
				analystr = "SELECT jobcategory.category, industry.industry, leftnum, rightnum, middlevalue, average FROM cateindustry, industry, jobcategory where cateindustry.industry = industry.IndustryId and cateindustry.categoryId = jobcategory.CategoryId and cateindustry.categoryId=" + v
				rows, err := thesisdb.Query(analystr)
				if err != nil {
					log.Fatal(err)
				}

				for rows.Next() {
					var ar Analyresult
					rows.Scan(&analysis.Category, &ar.Description, &ar.Left, &ar.Right, &ar.Middle, &ar.Average)
					analysis.Target.Industry = append(analysis.Target.Industry, ar)
				}
				if analysis.Category == "" {
					analysis.Category = as[k]
				}

				rows.Close()
			}
			if i == 1 {
				analystr = `SELECT  district, leftnum, rightnum, middlevalue, average FROM catedistrict, jobcategory where catedistrict.categoryId = jobcategory.CategoryId and catedistrict.categoryId =` + v
				rows, err := thesisdb.Query(analystr)
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
				analystr = `SELECT   exp, leftnum, rightnum, middlevalue, average FROM cateexp, jobcategory where cateexp.categoryId = jobcategory.CategoryId and cateexp.categoryId =` + v
				rows, err := thesisdb.Query(analystr)
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
			// fmt.Println(analysis)
		}
		r = append(r, analysis)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": r,
	})
}
