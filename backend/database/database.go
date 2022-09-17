package database

import (
	"Eat100/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

//fixme  eat100を手で作る必要あり

func DBOpen() {
	dbConn = NewSQLHandler()
}

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

func NewSQLHandler() *SQLHandler {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	//dsn := "root:password@tcp(Eat100DB)/eat100?charset=utf8mb4&parseTime=true"
	dsn := user + ":" + password + "@tcp(" + host + ")" + "/" + dbName + "?charset=utf8mb4&parseTime=true"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.DB = db
	return sqlHandler
}

func GetDBConn() *SQLHandler {
	return dbConn
}

func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

func RollBack() {
	dbConn.DB.Rollback()
}

func DbInit() *gorm.DB {
	dsn := "root:password@tcp(Eat100DB)/eat100?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Inserts(db *gorm.DB, restaurantDetails []entity.Restaurant) {
	result := db.Create(&restaurantDetails)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}

//todo これ本当に必要かよく考える
//func NewDB(detail entity.restaurantDetail) *gorm.DB {
//	db := DbInit()
//	db.AutoMigrate(&detail)
//	return db
//}

func CreateRestaurantDetails(db *gorm.DB, restaurantDetails []entity.Restaurant) {
	Inserts(db, restaurantDetails)
}
