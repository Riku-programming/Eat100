package handler

import (
	"Eat100/restaurant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestaurantsGet(restaurants *restaurant.Restaurants) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := restaurants.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
