package database

import (
	"Eat100/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DbInit() *gorm.DB {
	dsn := "root:password@tcp(shop-db)/shop_db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Inserts(db *gorm.DB, shopDetails []entity.ShopDetail) {
	result := db.Create(&shopDetails)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}
func NewDB(detail entity.ShopDetail) *gorm.DB {
	db := DbInit()
	db.AutoMigrate(&detail)
	return db
}

func CreateShopDetails(db *gorm.DB, shopDetails []entity.ShopDetail) {
	Inserts(db, shopDetails)
}
