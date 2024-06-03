package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harivilasp/url-shortner-go/cache"
	"github.com/harivilasp/url-shortner-go/dbservice"
	"github.com/harivilasp/url-shortner-go/handler"
)

func main() {
	r := gin.Default()
	// r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to url shortener API 🚀",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/re/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	dbservice.InitializeDB()
	cache.InitializeCacheService()

	err := r.Run(":9098")

	if err != nil {
		panic(err.Error())
	}
}
