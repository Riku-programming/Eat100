package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type ShopDetail struct {
	ShopName      string
	ReserveStatus string
	Address       string
	Time          string
	Payment       string
	PhoneNumber   string
	Cost          string
}

func main() {
	scraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
}

func scraping(URL string) {
	shopDetails := make([]ShopDetail, 0)
	// collyインスタンス
	c := colly.NewCollector()
	detailCollector := c.Clone()
	// html -> .hyakumeiten-shop__item要素にアクセス
	c.OnHTML(".hyakumeiten-shop__item", func(e *colly.HTMLElement) {
		//詳細リンク
		link := e.ChildAttr(".hyakumeiten-shop__target", "href")
		detailCollector.Visit(e.Request.AbsoluteURL(link))
	})
	detailCollector.OnHTML(".rstinfo-table", func(e *colly.HTMLElement) {
		shopDetail := ShopDetail{
			ShopName:      e.ChildText(".rstinfo-table__name-wrap"),
			ReserveStatus: e.ChildText(".rstinfo-table__reserve-status"),
			Address:       e.ChildText(".listlink"),
			Time:          e.ChildText(".rstinfo-table__subject-text"),
			Payment:       e.ChildText(".rstinfo-table__pay-item"),
			PhoneNumber:   e.DOM.Find(".rstinfo-table__tel-num").First().Text(),
			Cost:          e.DOM.Find(".gly-b-dinner").First().Text(),
		}
		shopDetails = append(shopDetails, shopDetail)
	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("ここにリクエストするよ", request.URL.String())
	})
	// スクレイピング
	c.Visit(URL)
	// 後処理
	fmt.Println(shopDetails)
}
