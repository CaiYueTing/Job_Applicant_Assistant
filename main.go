package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"sort"
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
	c.JSON(200, gin.H{
		"records": records,
	})

}

func makescore(c *gin.Context) {
	welfare := c.Param("welfare")

	//key word regex for money
	three := regexp.MustCompile("三節")
	yearend := regexp.MustCompile("年終")
	birth := regexp.MustCompile("生日")
	marry := regexp.MustCompile("結婚")
	maternity := regexp.MustCompile("生育")
	patent := regexp.MustCompile("專利")
	longterm := regexp.MustCompile("久任")
	insurance := regexp.MustCompile("團保")
	stock := regexp.MustCompile("股票")
	stock1 := regexp.MustCompile("入股")
	annual := regexp.MustCompile("分紅")
	attendance := regexp.MustCompile("全勤")
	performance := regexp.MustCompile("績效")
	travel := regexp.MustCompile("旅遊補助")
	travel1 := regexp.MustCompile("旅遊津貼")
	consolation := regexp.MustCompile("慰問")
	health := regexp.MustCompile("健康檢查")
	health1 := regexp.MustCompile("體檢")

	//key word regex for working time

	flexible := regexp.MustCompile("彈性上下班")
	paternityleave := regexp.MustCompile("陪產假")
	travelleave := regexp.MustCompile("旅遊假")
	physiologyleave := regexp.MustCompile("生理假")
	fullpaysickleave := regexp.MustCompile("全薪病假")
	fullpaysickleave1 := regexp.MustCompile("不扣薪病假")

	//key word regex for infrastructure

	dorm := regexp.MustCompile("宿舍")
	restaurant := regexp.MustCompile("餐廳")
	childcare := regexp.MustCompile("托兒")
	childcare1 := regexp.MustCompile("育兒")
	transport := regexp.MustCompile("交通")
	servemeals := regexp.MustCompile("供餐")
	servemeals1 := regexp.MustCompile("餐點")
	afternoon := regexp.MustCompile("下午茶")
	snack := regexp.MustCompile("點心")
	gym := regexp.MustCompile("健身房")

	//key word regex for entertainment

	education := regexp.MustCompile("教育訓練")
	tail := regexp.MustCompile("尾牙")
	tail1 := regexp.MustCompile("旺年會")
	employeetravel := regexp.MustCompile("員工旅遊")
	society := regexp.MustCompile("社團")

	//key word regex for unusually

	overtime := regexp.MustCompile("加班")
	shift := regexp.MustCompile("輪班")
	permanent := regexp.MustCompile("外派")
	permanent1 := regexp.MustCompile("長駐")

	score :=
		btou(three.MatchString(welfare)) +
			btou(yearend.MatchString(welfare)) +
			btou(birth.MatchString(welfare)) +
			btou(marry.MatchString(welfare)) +
			btou(maternity.MatchString(welfare)) +
			btou(patent.MatchString(welfare)) +
			btou(longterm.MatchString(welfare)) +
			btou(insurance.MatchString(welfare)) +
			btou(stock.MatchString(welfare) || stock1.MatchString(welfare)) +
			btou(annual.MatchString(welfare)) +
			btou(attendance.MatchString(welfare)) +
			btou(performance.MatchString(welfare)) +
			btou(travel.MatchString(welfare) || travel1.MatchString(welfare)) +
			btou(consolation.MatchString(welfare)) +
			btou(health.MatchString(welfare) || health1.MatchString(welfare)) +
			btou(flexible.MatchString(welfare)) +
			btou(paternityleave.MatchString(welfare)) +
			btou(travelleave.MatchString(welfare)) +
			btou(physiologyleave.MatchString(welfare)) +
			btou(fullpaysickleave.MatchString(welfare) || fullpaysickleave1.MatchString(welfare)) +
			btou(dorm.MatchString(welfare)) +
			btou(restaurant.MatchString(welfare)) +
			btou(childcare.MatchString(welfare) || childcare1.MatchString(welfare)) +
			btou(transport.MatchString(welfare)) +
			btou(servemeals.MatchString(welfare) || servemeals1.MatchString(welfare)) +
			btou(snack.MatchString(welfare)) +
			btou(afternoon.MatchString(welfare)) +
			btou(gym.MatchString(welfare)) +
			btou(education.MatchString(welfare)) +
			btou(tail.MatchString(welfare) || tail1.MatchString(welfare)) +
			btou(employeetravel.MatchString(welfare)) +
			btou(society.MatchString(welfare)) +
			btou(overtime.MatchString(welfare)) +
			btou(shift.MatchString(welfare)) +
			btou(permanent.MatchString(welfare) || permanent1.MatchString(welfare))
	c.JSON(200, gin.H{
		"message": score,
	})
}

func main() {

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	// querypoint()
	r := gin.Default()
	r.GET("/welfare/:welfare", makescore)
	r.GET("/law/:company", query)
	r.GET("/query/:company", query)
	r.Run()
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
		w := el.wtoi()
		point = append(point, w)
	}

	sort.Ints(point)
	// fmt.Println(len(a))
	dividindex := len(point) / 10
	fmt.Println("slice index", dividindex)
	fmt.Println(point[dividindex], point[dividindex*2], point[dividindex*3], point[dividindex*4], point[dividindex*5],
		point[dividindex*6], point[dividindex*7], point[dividindex*8], point[dividindex*9], point[len(point)-1],
	)

	end := time.Now()
	fmt.Println("end time: ", end.Sub(start).Seconds())
}

func (w Welfarepoint) wtoi() int {
	result :=
		btou(w.Three)*2 + btou(w.Yearend)*2 + btou(w.Bitrh)*2 + btou(w.Marry)*2 + btou(w.Maternity)*2 +
			btou(w.Patent)*2 + btou(w.Longterm)*2 + btou(w.Insurance)*3 + btou(w.Stock)*3 + btou(w.Annual)*2 +
			btou(w.Attendance) + btou(w.Performance)*2 + btou(w.Travel)*2 + btou(w.Consolation)*2 + btou(w.Health)*3 +
			btou(w.Flexible)*3 + btou(w.Paternityleave)*3 + btou(w.Travelleave)*3 + btou(w.Physiologyleave)*3 + btou(w.Fullpaysickleave)*3 +
			btou(w.Dorm)*3 + btou(w.Restaurant)*2 + btou(w.Childcare)*3 + btou(w.Transport)*2 + btou(w.Servemeals)*2 +
			btou(w.Snack) + btou(w.Afternoon) + btou(w.Gym)*3 + btou(w.Education)*3 + btou(w.Tail)*3 +
			btou(w.Employeetravel)*3 + btou(w.Society)*3 + btou(w.Overtime)*(-1) + btou(w.Shift)*(-1) + btou(w.Permanent)*(-1)
	return result
}

func btou(b bool) int {
	if b {
		return 1
	}
	return 0
}
