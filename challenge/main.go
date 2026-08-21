package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Message string
	Code    int
	Status  string
	Data    any
}

type Articles struct {
	ID      int
	Title   string
	Content string
}

func main() {
	route := gin.Default()

	// route
	route.GET("/articles", func(c *gin.Context) {
		articles := []Articles{}
		for i := 0; i < 20; i++ {
			article := &Articles{
				ID:      i,
				Title:   "Prabowo gibran",
				Content: "Gibran prabowo",
			}
			articles = append(articles, *article)
		}

		apiResponse := &ApiResponse{
			Message: "Berhasil mengambil semua articles",
			Code:    http.StatusOK,
			Status:  "OK",
			Data:    articles,
		}
		c.JSON(http.StatusOK, apiResponse)
	})

	route.GET("/articles/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			response := ApiResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Data:    nil,
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		article := &Articles{
			ID:      id,
			Title:   "Prabowo gibran",
			Content: "Gibran prabowo",
		}
		response := ApiResponse{
			Message: "Article berhasil di ambil",
			Code:    http.StatusOK,
			Status:  "OK",
			Data:    article,
		}
		c.JSON(http.StatusBadRequest, response)
	})

	route.POST("/articles", func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)
		err := json.NewEncoder(c.Writer).Encode(decoder)
		if err != nil {
			response := ApiResponse{
				Message: "Gagal menambah data baru",
				Code:    http.StatusBadRequest,
				Status:  "fail",
				Data:    nil,
			}
			c.JSON(http.StatusCreated, response)
			return
		}
		c.Request.Header.Add("Content-Type", "application/json")
		response := ApiResponse{
			Message: "Berhasil menambah data",
			Code:    http.StatusCreated,
			Status:  "OK",
			Data:    "c.Request",
		}
		c.JSON(http.StatusCreated, response)
	})

	route.Run(":3000")
}
