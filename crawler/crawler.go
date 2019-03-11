package crawler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// func main() {
// 	CrawlQollie("饗賓餐旅")
// }

type C interface {
	CrawlNews()
}

func CrawlNews(s string) {

}

func CrawlPPT(s string) {

}

func CrawlQollie(s string) {
	uri := "https://www.qollie.com/graphql"
	enuri := "https://www.qollie.com/search?keyword=" + s + "&kind=company&from=normal"
	graphqlreader := "{\"query\":\"\\n  \\nfragment commonFields on Company {\\n  _id\\n  authentication\\n  name\\n  category\\n  website\\n  introduction\\n  sourcesLinks\\n  createdAt\\n  comments\\n  enableNotify\\n  authApplication {\\n    _id\\n    updatedAt\\n  }\\n  jobs {\\n    _id\\n    jobTitle\\n  }\\n  announcement\\n  tags\\n  logo\\n  businessAddress {\\n    kind\\n    county\\n    district\\n    detail\\n  }\\n  address\\n  website\\n  facebook\\n  instagram\\n  linkedin\\n  email\\n  taxId\\n}\\n\\n  query search(\\n    $kind: String\\n    $keyword: String\\n    $page: Int\\n    $limit: Int\\n  ) {\\n    searchCompanies(query: {\\n      kind: $kind\\n      keyword: $keyword\\n      limit: $limit\\n      page: $page\\n    }) {\\n      ... commonFields\\n    }\\n  }\\n  \",\"variables\":{\"kind\":\"company\",\"keyword\":\""
	g := graphqlreader + s + "\",\"page\":1,\"limit\":10}}"
	payload := strings.NewReader(g)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", uri, payload)
	cookie := "__cfduid=d8fad860ff02636d5016f163d30bfaf8d1537509983;_ga=GA1.2.542764508.1538989300;_ym_uid=1538989300959422428;_ym_d=1538989300;__qca=P0-2091743772-1538989299985;_ZB_STATIC_DR_widgetsUpdateTime={'296887':1513584046};_ZB_STATIC_DR_firstTimeVisit=1538989301473;_ZB_STATIC_VIEW_THROUGH_WIDGETS=[296887];__smVID=2647928e5647eeaec69c936d990774139ae8158a3c386c7ff64192a1f7daea7b;_gid=GA1.2.511954536.1552197765;_ym_wasSynced={'time':1552197771241,'params':{'eu':0},'bkParams':{}};_ZB_STATS_VISIT=true;_ZB_STATIC_DR_currentSessionTimeVisit=1552197777421;_ym_isad=2;_ZB_STATS_IMPRESSION.2340986f=true;_ZB_STATS_IMPRESSION_FREEMIUM_=true;_hp2_ses_props.2694192674={'r':'https://www.google.com/','ts':1552225115880,'d':'www.qollie.com','h':'/search'};_ym_visorc_47843744=w;__smToken=RWCWjyY31IYekTZiNwmZhFPI;connect.sid=s:v_WzLOWInUCdHDyvfpJTMqJldUZEoEDb.PaFvFasO3MmGP/5h09YAFUQYQFQJz50WQnL5CfiCfSI;_ZB_STATIC_296887_STATUS=closed;_ZB_ADMIN_TIME_STAMP_=1552229730344;_ZB_ADMIN_LAST_URL_=https://www.qollie.com/;_dc_gtm_UA-82256880-2=1;mp_9106e91156ed01460ec46c42109de098_mixpanel={'distinct_id': '16652e919002c-056fee21ee497f-8383268-1fa400-16652e919017a','$initial_referrer': 'https://www.104.com.tw/jobbank/custjob/index.php?r=cust&j=4a4048713a3c446d363840693e443c1f22f2f2f2a463c482624j97','$initial_referring_domain': 'www.104.com.tw','$search_engine': 'google'};_hp2_id.2694192674={'userId':'6924317099599276','pageviewId':'7019326964272686','sessionId':'3728285964555434','identity':null,'trackerVersion':'4.0'};_gat_UA-82256880-2=1"
	req.Header.Add("x-recaptcha", "undefined")
	req.Header.Add("origin", "https://www.qollie.com")
	req.Header.Add("accept-encoding", "")
	req.Header.Add("accept-language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfaWQiOiI1YWYxMTc1NTQ2MzNhNDAxMTYwOTRmMmEiLCJpYXQiOjE1NTIyMjgwMTUsImV4cCI6MTU1MjMxNDQxNX0.Q0AqstPn3GlkGj7S5Ksjgj5FCYJHjKhjiFZJlTw28vE")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	req.Header.Add("content-type", "application/json; ")
	req.Header.Add("accept", "*/*")
	req.Header.Add("referer", url.QueryEscape(enuri))
	req.Header.Add("authority", "www.qollie.com")
	req.Header.Add("cookie", cookie)
	req.Header.Add("x-ssr-key", "undefined")
	req.Header.Add("cache-control", "no-cache")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		fmt.Println("status code:", res.StatusCode, res.Status)
	}
	dom, _ := goquery.NewDocumentFromResponse(res)
	dom.Find("div.src-client-components-SearchResultList-___style__container___2x5X7")
	if dom.Nodes == nil {
		fmt.Println("not found in div")
	}
	dom.Each(func(i int, s *goquery.Selection) {
		s.Find("a")
		if s.Nodes == nil {
			fmt.Println("not found in a href")
		}
		s.Each(func(j int, ss *goquery.Selection) {
			t := ss.Text()

			fmt.Println(t)
		})
	})

}

// func startcrawler(url string) (*goquery.Document, error) {

// }

func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}
