package main

import (
	"Eat100/article"
	"Eat100/database"
	"Eat100/entity"
	"Eat100/handler"
	"Eat100/shop_detail_scraping"
	"github.com/gin-gonic/gin"
	"strings"
)

const baseURL string = "https://award.tabelog.com/hyakumeiten/"

func main() {
	Init()
	article := article.New()
	database.DBOpen()
	defer database.DBClose()
	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(article))
	r.Run(":8080")
}

func Init() {
	var detail entity.Restaurant
	db := database.DbInit()
	db.AutoMigrate(&detail)
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
