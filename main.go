package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	// data.csvを作成する
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	defer file.Close()

	// csv操作インスタンス
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ShopName", "Area", "Holiday", "Rating", "Target"})

	// collyインスタンス
	c := colly.NewCollector(
	//colly.AllowedDomains("tabelog.com"),
	)

	// html -> .internship_meta要素にアクセス
	c.OnHTML(".hyakumeiten-shop__item", func(e *colly.HTMLElement) {
		writer.Write([]string{
			//店名
			e.ChildText(".hyakumeiten-shop__name"),
			//定休日
			e.ChildText(".hyakumeiten-shop__holiday"),
			//住所
			e.ChildText(".hyakumeiten-shop__area"),
			// 評価
			e.ChildText(".hyakumeiten-shop__rating"),
			//詳細リンク
			e.ChildAttr(".hyakumeiten-shop__target", "href"),
		})
	})

	start := time.Now()
	// スクレイピング
	c.Visit("https://award.tabelog.com/hyakumeiten/izakaya/2021/")

	// 後処理
	end := time.Now()
	log.Println("Scraping Complete")
	log.Printf("Time:%v", end.Sub(start))
	log.Println(c)
}
