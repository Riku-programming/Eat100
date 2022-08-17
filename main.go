package main

import (
	"Eat100/database"
	"Eat100/entity"
	"Eat100/shop_detail_scraping"
	"fmt"
	"strings"
)

func main() {
	//var detail entity.ShopDetail
	//db := database.NewDB(detail)
	//shopDetails := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
	//izakaya := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
	//unagi := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/unagi/")
	//sushi := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/sushi_tokyo/")
	PopularCategoryProcessing("https://award.tabelog.com/hyakumeiten/sushi_tokyo/")
	//database.CreateShopDetails(db, sushi)
	//database.CreateShopDetails(db, unagi)
}

// todo 関数名ちゃんと考える, この実装綺麗にすること
// _が含まれていたら自動的にtokyo,west,east３つを足してスクレイピングするようにする

func PopularCategoryProcessing(category string) {
	var detail entity.ShopDetail
	db := database.NewDB(detail)
	for _, v := range []string{"tokyo", "east", "west"} {
		path := strings.Split(category, "_")[0] + "_" + v
		fmt.Println(path)
		three := shop_detail_scraping.ShopDetailScraping(path)
		fmt.Println(three)
		database.CreateShopDetails(db, three)
	}
}
