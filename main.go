package main

import (
	"Eat100/database"
	"Eat100/entity"
	"Eat100/shop_detail_scraping"
)

func main() {
	var detail entity.ShopDetail
	db := database.NewDB(detail)
	//shopDetails := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
	izakaya := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
	unagi := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/unagi/")
	//fmt.Println(shopDetails)
	database.CreateShopDetails(db, izakaya)
	database.CreateShopDetails(db, unagi)
}
