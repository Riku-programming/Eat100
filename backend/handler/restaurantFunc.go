package handler

import (
	"Eat100/entity"

	"Eat100/restaurant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestaurantsGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurant := restaurant.New()
		params := entity.SearchResultParams{
			Category:   c.Query("category"),
			Address:    c.Query("address"),
			Reservable: c.Query("reservable"),
			Payment:    c.Query("payment"),
		}
		searchResultMap := entity.StructToMap(&params)
		result := restaurant.GetSearchResult(searchResultMap)
		c.JSON(http.StatusOK, result)
	}
}
