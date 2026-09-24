package gin_binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UriBinding(c *gin.Context) {
	var person UriPerson
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
}
