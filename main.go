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
		Message: "Berhasil mengambil data buku",
		Status:  "OK",
		Code:    http.StatusOK,
		Data: []gin.H{
			{
				"id":   1,
				"name": "ilham",
				"nim":  "123432",
			},
			{
				"id":   2,
				"name": "luffy",
				"nim":  "123432",
			},
			{
				"id":   3,
				"name": "ilham",
				"nim":  "123432",
			},
		},
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, apiResponse)
	})

	router.GET("/form", func(c *gin.Context) {
		c.File("./index.html")
	})

	router.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON(http.StatusCreated, gin.H{
			"user": name,
		})
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("message", "anonymus")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/query_post", func(c *gin.Context) {
		name := c.Query("name")
		message := c.PostForm("message")
		c.String(http.StatusCreated, "Halo %s dari query\nHalo %s dari message", name, message)
	})

	router.POST("/query_map", func(c *gin.Context) {
		category_filter := c.QueryMap("category_filter")
		c.JSON(http.StatusCreated, gin.H{
			"category_filter": category_filter,
		})
	})

	err := router.Run(":3000")
	if err != nil {
		return
	}
}
