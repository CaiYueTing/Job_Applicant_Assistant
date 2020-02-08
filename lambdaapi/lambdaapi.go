package lambdaapi

import (
	"fmt"
	"jobassistant-server/crawler"
	"jobassistant-server/welfare"
	"net/http"
	"strings"

	dynamo "github.com/CaiYueTing/dynamoHelper"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

type Analysis struct {
	CategoryID   string
	CategoryName string
	Target       struct {
		Industry []Analyresult
		Exp      []Analyresult
		District []Analyresult
	}
}

type Analyresult struct {
	Description string
	Rightnum    int
	Leftnum     int
	Middlevalue int
	Average     int
}

type Welfare struct {
	Custno      string
	Score       []int
	LastUpdated string
}

type IllegalRecord struct {
	Company     string
	Location    string
	Publicdate  string
	Dealdate    string
	Govnumber   string
	Law         string
	Description string
	Ps          string
}

func Category(c *gin.Context) {
	cstring := c.PostForm("cdata")
	as := strings.Split(cstring, "、")
	resultstring := []string{}
	for _, v := range as {
		strings.Replace(v, "、", "", -1)
		resultstring = append(resultstring, v)
	}
	db := dynamo.NewDynamo("ap-northeast-2", "cate_salary")
	var ws []Analysis
	for _, v := range resultstring {
		result, err := db.QueryTableWithIndex(
			"categoryName",
			"categoryName-index",
			v,
			"EQ",
		)
		if err != nil {
			checkerr(err)
		}
		var w Analysis
		err = dynamodbattribute.UnmarshalMap(result[0], &w)
		ws = append(ws, w)
	}
	c.JSON(200, gin.H{
		"message": ws,
	})
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
	company := c.Param("company")
	s := company
	s = dealstring(s)

	db := dynamo.NewDynamo("ap-northeast-2", "illegal_record")
	size, _ := db.GetTableSize()

	size = (size / 1048576) + 1

	result := db.ScanTable(size, "Company", s)
	var records []IllegalRecord
	for _, v := range result {
		var w IllegalRecord
		dynamodbattribute.UnmarshalMap(v, &w)
		records = append(records, w)
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

func Salary(c *gin.Context) {
	salary := c.Param("salary")
	c.JSON(200, gin.H{
		"salary": salary,
	})
}

func Postscore(c *gin.Context) {
	var p []int
	var ud string
	db := dynamo.NewDynamo("ap-northeast-2", "welfarepoint")
	result, err := db.GetItem("Custno", "all")
	if err != nil {
		checkerr(err)
	} else {
		var w Welfare
		err = dynamodbattribute.UnmarshalMap(result.Item, &w)
		p = w.Score
	}

	wstring := c.PostForm("wdata")
	var w welfare.Welfarepoint
	r := w.Match2(wstring)
	score := w.Wtoi()
	c.JSON(http.StatusOK, gin.H{
		"message": score,
		"dd":      p,
		"r":       r,
		"update":  ud,
	})
}

func checkerr(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case dynamodb.ErrCodeProvisionedThroughputExceededException:
			fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
		case dynamodb.ErrCodeResourceNotFoundException:
			fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
		case dynamodb.ErrCodeRequestLimitExceeded:
			fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
		case dynamodb.ErrCodeInternalServerError:
			fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {

		fmt.Println(err.Error())
	}
	return
}
