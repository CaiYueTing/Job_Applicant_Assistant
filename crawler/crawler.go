package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
			CreatedAt       string      `json:"createdAt"`
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

type C interface {
	CrawlNews()
}

func CrawlNews(s string) {

}

func CrawlPPT(s string) {

}

func CrawlQollie(s string) string {
	companyid := getQollieUrl(s)
	url := "https://www.qollie.com/companies/" + companyid

	return url
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

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	var q Qollie
	err := json.Unmarshal(body, &q)

	if err != nil {
		fmt.Println(err)
	}
	if len(q.Data.SearchCompanies) == 0 {
		return ""
	}
	res.Body.Close()
	return q.Data.SearchCompanies[0].ID
}
