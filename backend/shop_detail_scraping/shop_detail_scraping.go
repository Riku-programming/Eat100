package shop_detail_scraping

import (
	"Eat100/entity"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

const (
	Reservable      = "予約可"
	AppointmentOnly = "完全予約制"
	NotReservable   = "予約不可"
)

const (
	CardAccepted = "カード可"
)

// スクレイピングする際の並列実行数
const parallelism int = 128

func ShopDetailScraping(URL string, categoryName string) []entity.Restaurant {
	shopDetails := make([]entity.Restaurant, 0)
	// collyインスタンス
	c := *colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)
	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: parallelism})
	if err != nil {
		return nil
	}
	detailCollector := c.Clone()
	// html -> .hyakumeiten-shop__item要素にアクセス
	c.OnHTML(".hyakumeiten-shop__item", func(e *colly.HTMLElement) {
		//詳細リンク
		link := e.ChildAttr(".hyakumeiten-shop__target", "href")
		detailCollector.Visit(e.Request.AbsoluteURL(link))

	})
	detailCollector.OnHTML(".rstinfo-table", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.URL.String())
		shopDetail := entity.Restaurant{
			//Category:    category.ClassifyCategory(categoryName),
			Category:    categoryName,
			ShopName:    e.ChildText(".rstinfo-table__name-wrap"),
			Reservable:  isReservable(e.ChildText(".rstinfo-table__reserve-status")),
			Address:     e.ChildText(".listlink"),
			Time:        e.ChildText(".rstinfo-table__subject-text"),
			Payment:     isPayment(e.ChildText(".rstinfo-table__pay-item")),
			PhoneNumber: e.DOM.Find(".rstinfo-table__tel-num").First().Text(),
			Cost:        costToInt(e.DOM.Find(".gly-b-dinner").First().Text()),
			URL:         e.Request.URL.String(),
		}
		fmt.Println(shopDetail)
		shopDetails = append(shopDetails, shopDetail)
	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("ここにリクエストするよ", request.URL.String())
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	// スクレイピング
	err = c.Visit(URL)
	if err != nil {
		return nil
	}
	c.Wait()
	detailCollector.Wait()
	//// 後処理
	fmt.Println(shopDetails)
	return shopDetails
}

func isReservable(status string) bool {
	switch status {
	case Reservable:
		return true
	case NotReservable:
		return false
	case AppointmentOnly:
		return false
	default:
		return false
	}
}

func isPayment(status string) bool {
	if strings.Contains(status, CardAccepted) {
		return true
	}
	return false
}

// CostToInt スクレイピングしたstring型のCostをint型に変換する
// example ￥30,000～￥39,999 → 30000
func costToInt(costString string) int {
	replace1 := strings.Replace(costString, "￥", "", -1)
	replace2 := strings.Replace(replace1, ",", "", -1)
	replace3 := strings.Split(replace2, "～")
	cost, _ := strconv.Atoi(replace3[0])
	return cost
}
