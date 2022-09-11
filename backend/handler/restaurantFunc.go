package handler

import (
	"Eat100/restaurant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestaurantsGet(restaurants *restaurant.Restaurants) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		address := c.Query("address")
		reservable := c.Query("reservable")
		payment := c.Query("payment")
		result := restaurants.GetAll(category, address, reservable, payment)
		c.JSON(http.StatusOK, result)
	}
}
