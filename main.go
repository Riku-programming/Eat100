package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

type ShopDetail struct {
	PhoneNumber   string
	ReserveStatus string
	Address       string
	Time          string
	Payment       string
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
	detailCollector.OnHTML(".rstinfo-table__table", func(e *colly.HTMLElement) {
		shopDetail := ShopDetail{
			PhoneNumber:   e.ChildText(".rstinfo-table__tel-num"),
			ReserveStatus: e.ChildText(".rstinfo-table__reserve-status"),
			Address:       e.ChildText(".rstinfo-table__reserve-address"),
			Time:          e.ChildText(".rstinfo-table__subject-text"),
			Payment:       e.ChildText(".rstinfo-table__pay-item"),
		}
		shopDetails = append(shopDetails, shopDetail)
	})
	// スクレイピング
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("ここにリクエストするよ", request.URL.String())

	})
	start := time.Now()
	c.Visit(URL)

	// 後処理
	fmt.Println(shopDetails)
	end := time.Now()
	log.Println("Scraping Complete")
	log.Printf("Time:%v", end.Sub(start))
	log.Println(c)
}
