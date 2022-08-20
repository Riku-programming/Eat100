package main

import (
	"Eat100/category"
	"Eat100/database"
	"Eat100/entity"
	"Eat100/shop_detail_scraping"
	"strings"
)

const baseURL string = "https://award.tabelog.com/hyakumeiten/"

var keyWord = "sushi_tokyo"
var URL = baseURL + keyWord

func main() {
	var detail entity.ShopDetail
	db := database.NewDB(detail)
	categoryName := category.ExtractCategoryFromURL(URL)
	if category.IsPopular(URL) == true {
		for _, v := range []string{"tokyo", "east", "west"} {
			URL := baseURL + strings.Split(categoryName, "_")[0] + "_" + v
			popular := shop_detail_scraping.ShopDetailScraping(URL, categoryName)
			database.CreateShopDetails(db, popular)
		}
		return
	}
	tonkatu := shop_detail_scraping.ShopDetailScraping(URL, categoryName)
	database.CreateShopDetails(db, tonkatu)
}
