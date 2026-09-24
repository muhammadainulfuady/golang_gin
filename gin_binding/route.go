package gin_binding

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/query-person", queryStringBinding)
	router.GET("/form", formViews)
	router.POST("/default-person", defaultValudeBinding)
	router.GET("/:name/:id", UriBinding)
}
