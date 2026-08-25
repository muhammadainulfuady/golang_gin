package gin_binding

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/person", queryStringBinding)
	router.GET("/form", formViews)
}
