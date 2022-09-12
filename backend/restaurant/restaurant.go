package restaurant

import (
	"Eat100/database"
	"Eat100/entity"
	"fmt"
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

func (r *Restaurants) GetSearchResult(fields map[string]string) []SearchResult {
	var searchResult []SearchResult
	db := database.GetDBConn().DB
	queryDB := db.Table("restaurants")
	for key, value := range fields {
		if value != "" {
			s := fmt.Sprintf("%s LIKE ?", key)
			queryDB = queryDB.Where(s, value+"%")
		}
	}
	if err := queryDB.Find(&searchResult).Error; err != nil {
		return nil
	}
	return searchResult
}
