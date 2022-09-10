package main

import (
	"Eat100/database"
	"Eat100/entity"
	"Eat100/handler"
	"Eat100/restaurant"
	"Eat100/shop_detail_scraping"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

const baseURL string = "https://award.tabelog.com/hyakumeiten/"

func main() {
	setEnv()
	//Init()
	restaurant := restaurant.New()
	database.DBOpen()
	defer database.DBClose()
	r := gin.Default()
	r.GET("/restaurant", handler.RestaurantsGet(restaurant))
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

func setEnv() {
	env := os.Getenv("ENV")
	fmt.Println(env)
	if "" == env {
		env = "development"
	}
	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		fmt.Println("Its not test environment")
	}

	if "test" != env {
		err := godotenv.Load(".env.local")
		if err != nil {
			fmt.Println("cant load .env.local")
		}
	}
	godotenv.Load(".env." + env)
	fmt.Println(os.Getenv("PORT"))
	godotenv.Load()
}
