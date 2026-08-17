package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

func main() {
	router := gin.Default()
	router.GET("/alamat", func(c *gin.Context) {
		alamat := c.Query("search")
		c.String(http.StatusOK, "Alamat oke %s", alamat)
	})
	err := router.Run(":3000")
	if err != nil {
		return
	}
}
