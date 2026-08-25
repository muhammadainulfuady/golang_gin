package main

import (
	"golang_gin/gin_binding"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	gin_binding.SetupRoutes(route)

	err := route.Run(":3000")
	if err != nil {
		return
	}
}
