package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Comment struct {
	Data struct {
		GetCompanyStat struct {
			ID     string `json:"id"`
			Good   int    `json:"good"`
			Bad    int    `json:"bad"`
			Normal int    `json:"normal"`
		} `json:"getCompanyStat"`
	} `json:"data"`
}

type Qollie struct {
	Data struct {
		SearchCompanies []struct {
			ID              string      `json:"_id"`
			Authentication  interface{} `json:"authentication"`
			Name            string      `json:"name"`
			Category        interface{} `json:"category"`
			Website         interface{} `json:"website"`
			Introduction    interface{} `json:"introduction"`
			SourcesLinks    []string    `json:"sourcesLinks"`
			CreatedAt       int64       `json:"createdAt"`
			Comments        []string    `json:"comments"`
			EnableNotify    bool        `json:"enableNotify"`
			AuthApplication interface{} `json:"authApplication"`
			Jobs            []struct {
				ID       string `json:"_id"`
				JobTitle string `json:"jobTitle"`
			} `json:"jobs"`
			Announcement    interface{} `json:"announcement"`
			Tags            interface{} `json:"tags"`
			Logo            interface{} `json:"logo"`
			BusinessAddress interface{} `json:"businessAddress"`
			Address         interface{} `json:"address"`
			Facebook        interface{} `json:"facebook"`
			Instagram       interface{} `json:"instagram"`
			Linkedin        interface{} `json:"linkedin"`
			Email           interface{} `json:"email"`
			TaxID           interface{} `json:"taxId"`
		} `json:"searchCompanies"`
	} `json:"data"`
}

var ch chan string

func CrawlNews(s string) {

}

func CrawlPPT(s string) {

}

func CrawlQollie(s string) Comment {
	companyid := getQollieUrl(s)
	var c Comment
	if companyid == "" {
		fmt.Println("empty")
		return c
	}
	url := "https://www.qollie.com/graphql"
	readstr := "{\"query\":\"\\n\\tquery getCompanyStat($id: ID!) {\\n\\t\\tgetCompanyStat(id: $id) {\\n\\t\\t\\tid\\n\\t\\t\\tgood\\n\\t\\t\\tbad\\n\\t\\t\\tnormal\\n\\t\\t}\\n\\t}\\n\",\"variables\":{\"id\":\""
	readstr = readstr + companyid + "\"}}"
	payload := strings.NewReader(readstr)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "70258b06-0758-479d-a9a2-8009b8ea3637")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Defaultclient err:", res.Status, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if body == nil {
		fmt.Println("body nil")
	}
	if len(body) == 0 {
		fmt.Println("body empty")
	}
	if err != nil {
		fmt.Println("ioutil read all err:", res.Status, err)
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("unmarshal json err:", err, body)
	}
	err = res.Body.Close()
	if err != nil {
		fmt.Println("response close err :", err)
	}
	return c
	// fmt.Println(q.Data.SearchCompanies[0].ID)
	// fmt.Println(string(body))
}

func getQollieUrl(s string) string {
	url := "https://www.qollie.com/graphql"
	readstring := "{\"query\":\"\\n  \\nfragment commonFields on Company {\\n  _id\\n  authentication\\n  name\\n  category\\n  website\\n  introduction\\n  sourcesLinks\\n  createdAt\\n  comments\\n  enableNotify\\n  authApplication {\\n    _id\\n    updatedAt\\n  }\\n  jobs {\\n    _id\\n    jobTitle\\n  }\\n  announcement\\n  tags\\n  logo\\n  businessAddress {\\n    kind\\n    county\\n    district\\n    detail\\n  }\\n  address\\n  website\\n  facebook\\n  instagram\\n  linkedin\\n  email\\n  taxId\\n}\\n\\n  query search(\\n    $kind: String\\n    $keyword: String\\n    $page: Int\\n    $limit: Int\\n  ) {\\n    searchCompanies(query: {\\n      kind: $kind\\n      keyword: $keyword\\n      limit: $limit\\n      page: $page\\n    }) {\\n      ... commonFields\\n    }\\n  }\\n  \",\"variables\":{\"kind\":\"company\",\"keyword\":\""
	readstring = readstring + s + "\",\"page\":1,\"limit\":10}}"

	payload := strings.NewReader(readstring)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "*/*")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "8d13357f-599b-47e8-aef9-9013c41f3d3d")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("get defaultclient err", err)
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get ioutilread all err", err)
		fmt.Println(body, res.Body)
		return ""
	}
	var q Qollie
	err = json.Unmarshal(body, &q)
	if err != nil {
		fmt.Println("get unmarshal err", err)
		return ""
	}

	if len(q.Data.SearchCompanies) == 0 {
		return ""
	}
	res.Body.Close()
	return q.Data.SearchCompanies[0].ID
}
