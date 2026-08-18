package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	MaxUploadSize = 1 << 20
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

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		// src, err := file.Open()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"err": "server err",
		// 	})
		// 	return
		// }
		// defer src.Close()

		// buffer := make([]byte, 512)
		// _, err = src.Read(buffer)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": "Gagal membaca isi file",
		// 	})
		// 	return
		// }

		// realContentType := http.DetectContentType(buffer)
		// allowedType := map[string]bool{
		// 	"image/jpeg": true,
		// 	"image/jpg":  true,
		// 	"image/png":  true,
		// 	"image/webp": true,
		// }

		// if !allowedType[c.ContentType()] {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": "File palsu terdeteksi! Tipe asli: " + realContentType,
		// 	})
		// 	return
		// }
		// c:/users/Asus/index.png

		dst := filepath.Join("./uploadd/file", filepath.Base(file.Filename))
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.MaxMultipartMemory = 1 << 20
	router.POST("/uploads", uploadHandler)

	err := router.Run(":3000")
	if err != nil {
		return
	}
}

func uploadHandler(c *gin.Context) {
	src, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "bad request pak",
		})
		return
	}

	files := src.File["file"]

	for _, file := range files {
		if file.Size > MaxUploadSize {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"message": "filemu kegeden pak",
			})
			return
		}
		dst := filepath.Join("./uploads/files/", filepath.Base(filepath.Base(file.Filename)))
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
