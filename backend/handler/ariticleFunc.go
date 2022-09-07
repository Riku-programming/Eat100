package handler

import (
	"Eat100/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := articles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
