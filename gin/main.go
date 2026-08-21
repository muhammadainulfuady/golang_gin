package main

import (
	"encoding/json"
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

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 1 << 20

	setupRoutes(router)

	err := router.Run(":3000")
	if err != nil {
		return
	}
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", getBooks)
	router.GET("/form", getForm)
	router.POST("/users", createUser)
	router.POST("/form_post", postForm)
	router.POST("/query_post", queryPost)
	router.POST("/query_map", queryMapHandler)
	router.POST("/upload", uploadFile)
	router.POST("/uploads", uploadHandler)
	router.POST("/bindig", modelBindingJson)
	router.POST("/encode", modelBindingJsonDecodeEncode)

	// tanpa middleware
	{
		v1 := router.Group("/api/v1")
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)

	}

	// dengan middleware
	v2 := router.Group("/api/v2")
	v2.Use(AuthMiddleware())
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}

	// nested route
	api := router.Group("/api")
	{
		v3 := api.Group("/v3")
		{
			v3.GET("/login", loginEndpoint)
			v3.GET("/submit", submitEndpoint)
			v3.GET("/read", readEndpoint)
			v3.GET("/users/:name", getUserByName)
		}
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		messageFrom(c, "dari middlewara bolo")
		c.Next()
	}
}

func getBooks(c *gin.Context) {
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
	c.JSON(http.StatusOK, apiResponse)
}

func getForm(c *gin.Context) {
	c.File("./index.html")
}

func createUser(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(http.StatusCreated, gin.H{
		"user": name,
	})
}

func postForm(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("message", "anonymus")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func queryPost(c *gin.Context) {
	name := c.Query("name")
	message := c.PostForm("message")
	c.String(http.StatusCreated, "Halo %s dari query\nHalo %s dari message", name, message)
}

func queryMapHandler(c *gin.Context) {
	category_filter := c.QueryMap("category_filter")
	c.JSON(http.StatusCreated, gin.H{
		"category_filter": category_filter,
	})
}

func uploadFile(c *gin.Context) {
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

func messageFrom(c *gin.Context, messageTo string) {
	c.JSON(http.StatusOK, gin.H{
		"message": &messageTo,
	})
}

func loginEndpoint(c *gin.Context) {
	messageFrom(c, "login")
}

func submitEndpoint(c *gin.Context) {
	messageFrom(c, "submit")
}

func readEndpoint(c *gin.Context) {
	messageFrom(c, "read")
}

func getUserByName(c *gin.Context) {
	name := c.Param("name")
	messageFrom(c, name)
}

func modelBindingJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "selamat anda login loh yah"})
}

func modelBindingJsonDecodeEncode(c *gin.Context) {
	var loginJSON Login

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&loginJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if loginJSON.User != "ilham" || loginJSON.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "username dan password salah",
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	encoder.Encode(gin.H{
		"status": "login sukses bolo",
	})
}
