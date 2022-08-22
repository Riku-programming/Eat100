package main

import (
	"Eat100/database"
	"Eat100/entity"
	"Eat100/shop_detail_scraping"
	"strings"
)

const baseURL string = "https://award.tabelog.com/hyakumeiten/"

func main() {
	var detail entity.ShopDetail
	db := database.NewDB(detail)
	categoryList := entity.CreateCategoryList()
	for _, v := range categoryList {
		categoryName := v.CategoryName()
		searchWord := v.SearchWord()
		if strings.Contains(searchWord, "_") == true {
			for _, v := range []string{"east", "west"} {
				URL := baseURL + strings.Split(searchWord, "_")[0] + "_" + v
				scrapingResult := shop_detail_scraping.ShopDetailScraping(URL, categoryName)
				database.CreateShopDetails(db, scrapingResult)
			}
		}
		scrapingResult := shop_detail_scraping.ShopDetailScraping(baseURL+searchWord, categoryName)
		database.CreateShopDetails(db, scrapingResult)
	}
}
