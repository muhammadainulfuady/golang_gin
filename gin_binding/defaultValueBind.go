package gin_binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultValudeBinding(c *gin.Context) {
	var persons BindinPerson

	if err := c.ShouldBind(&persons); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": persons.Name,
	})
}
