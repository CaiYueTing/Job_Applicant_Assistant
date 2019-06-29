package lambdaapi

import (
	"fmt"
	"net/http"
	"strings"
	"thesis/crawler"
	"thesis/welfare"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
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
	var ddb *dynamodb.DynamoDB
	ddb = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-2"))

	cstring := c.PostForm("cdata")
	as := strings.Split(cstring, "、")
	resultstring := []string{}
	for _, v := range as {
		strings.Replace(v, "、", "", -1)
		resultstring = append(resultstring, v)
	}

	// var a []Analysis
	var ws []Analysis
	for _, v := range resultstring {

		input := &dynamodb.QueryInput{
			KeyConditions: map[string]*dynamodb.Condition{
				"categoryName": {
					AttributeValueList: []*dynamodb.AttributeValue{{S: aws.String(v)}},
					ComparisonOperator: aws.String("EQ"),
				},
			},
			TableName: aws.String("cate_salary"),
			IndexName: aws.String("categoryName-index"),
		}
		result, err := ddb.Query(input)
		if err != nil {
			checkerr(err)
		}
		// fmt.Println(result.)
		var w Analysis
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &w)
		ws = append(ws, w)
	}

	// fmt.Println("with fmt println: ", w)
	// c.String(200, w.Custno, w.Score)
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

	var ddb *dynamodb.DynamoDB
	ddb = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-2"))
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String("illegal_record"),
	}
	result, err := ddb.DescribeTable(input)
	if err != nil {
		fmt.Println(err)
	}
	n := aws.Int64Value(result.Table.TableSizeBytes)

	n = (n / 1048576) + 1

	inputs := initInput(n, s) // all file divid into 16 segment (16MB)

	var records []IllegalRecord

	chrecord := make(chan []IllegalRecord, len(inputs))
	for i := 0; i < len(inputs); i++ {
		go func(i int) {
			record := scanner(inputs[i])
			chrecord <- record
		}(i)
	}
	for i := 0; i < len(inputs); i++ {
		record := <-chrecord
		for _, v := range record {
			records = append(records, v)
		}
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
	var ddb *dynamodb.DynamoDB
	var p []int
	var ud string
	ddb = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-2"))
	input := &dynamodb.GetItemInput{
		TableName: aws.String("welfarepoint"),
		Key:       map[string]*dynamodb.AttributeValue{"Custno": {S: aws.String("all")}},
	}
	result, err := ddb.GetItem(input)
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

func scanner(input *dynamodb.ScanInput) []IllegalRecord {
	var ddb *dynamodb.DynamoDB
	ddb = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-2"))
	result, err := ddb.Scan(input)
	if err != nil {
		checkerr(err)
	}
	var record []IllegalRecord
	for _, v := range result.Items {
		var w IllegalRecord
		err = dynamodbattribute.UnmarshalMap(v, &w)
		record = append(record, w)
	}

	return record
}

func initInput(totalSegment int64, search string) []*dynamodb.ScanInput {
	var inputs []*dynamodb.ScanInput
	var i int64
	for i = 0; i < totalSegment; i++ {
		input := &dynamodb.ScanInput{
			ScanFilter: map[string]*dynamodb.Condition{
				"Company": {
					ComparisonOperator: aws.String("CONTAINS"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(search),
						},
					},
				},
			},
			TableName:     aws.String("illegal_record"),
			Segment:       aws.Int64(int64(i)),
			TotalSegments: aws.Int64(totalSegment),
		}
		inputs = append(inputs, input)
	}
	return inputs
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
