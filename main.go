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
	apiResponse := ApiResponse{
		Message: "Berhasil mengambil data",
		Status:  "OK",
		Code:    http.StatusOK,
		Data: gin.H{
			"id":   1,
			"name": "ilham",
			"nim":  "123321",
		},
	}

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, apiResponse)
	})
	router.Run("localhost:2000")
}
