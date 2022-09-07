package article

import "Eat100/database"

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func New() *Articles {
	return &Articles{}
}

type Articles struct {
	Items []Article
}

func (r *Articles) GetAll() []Article {
	db := database.GetDBConn().DB
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		return nil
	}
	return articles
}
