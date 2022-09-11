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

type SearchResult struct {
	Category    string `json:",omitempty"`
	ShopName    string `json:",omitempty"`
	Reservable  bool   `json:",omitempty"`
	Address     string `json:",omitempty"`
	Time        string `json:",omitempty"`
	Payment     bool   `json:",omitempty"`
	PhoneNumber string `json:",omitempty"`
	Cost        int    `json:",omitempty"`
	URL         string `json:",omitempty"`
}

func (r *Restaurants) GetAll(categoryName, address, reservable, payment string) []SearchResult {
	db := database.GetDBConn().DB
	var searchResult []SearchResult
	// SELECT shop_name FROM restaurants WHERE category like '%' and address like '%' and reservable  like '%' and payment like '%'
	// fixme このクエリ汚いので直す
	if err := db.Table("restaurants").Where("category LIKE ? AND address LIKE ? AND reservable LIKE ? AND payment LIKE ?", categoryName+"%", address+"%", reservable+"%", payment+"%").Find(&searchResult).Error; err != nil {
		return nil
	}
	return searchResult
}
