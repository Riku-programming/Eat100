package restaurant

import (
	"Eat100/database"
	"Eat100/entity"
)

func New() *Restaurants {
	return &Restaurants{}
}

type Restaurants struct {
	Items []entity.Restaurant
}

func (r *Restaurants) GetAll() []entity.Restaurant {
	db := database.GetDBConn().DB
	var restaurants []entity.Restaurant
	// SELECT shop_name FROM restaurants WHERE category like '焼肉%'
	if err := db.Table("restaurants").Select("shop_name").Where("category LIKE ?", "焼肉").Find(&restaurants).Error; err != nil {
		return nil
	}
	return restaurants
}
