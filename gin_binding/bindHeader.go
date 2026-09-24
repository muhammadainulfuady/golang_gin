package gin_binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindHeader(c *gin.Context) {
	h := TestHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Rate":   h.Rate,
		"Domain": h.Domain,
	})
}
