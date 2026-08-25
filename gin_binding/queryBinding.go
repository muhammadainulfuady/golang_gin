package gin_binding

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func queryStringBinding(c *gin.Context) {
	var person Person

	err := c.ShouldBindQuery(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Addres,
	})
}

func formViews(c *gin.Context) {
	dst := filepath.Join("..", "views", "index.html")
	c.File(dst)
}
