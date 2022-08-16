package main

import (
	"Eat100/database"
	"Eat100/entity"
	"Eat100/shop_detail_scraping"
	"fmt"
)

func main() {
	var detail entity.ShopDetail
	db := database.NewDB(detail)
	shopDetails := shop_detail_scraping.ShopDetailScraping("https://award.tabelog.com/hyakumeiten/izakaya/2021/")
	fmt.Println(shopDetails)
	database.CreateShopDetails(db, shopDetails)
}
