package gin_binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultValudeBinding(c *gin.Context) {
	var persons Persons
	err := c.ShouldBind(&persons)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
}
