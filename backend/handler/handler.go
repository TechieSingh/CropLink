package handler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/harivilasp/url-shortner-go/cache"
	"github.com/harivilasp/url-shortner-go/dbservice"
	"github.com/harivilasp/url-shortner-go/shortener"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	// Send Bad Request to user if we don't get url and user id
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !IsUrl(creationRequest.LongUrl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Url, Make sure that the url has the scheme and host"})
		return
	}

	shortUrl := shortener.GenerateShortUrl(creationRequest.LongUrl, creationRequest.UserId)

	// Save the mapping in the DB and cache
	dbservice.InsertUrl(shortUrl, creationRequest.LongUrl)
	cache.SaveUrlMapping(creationRequest.LongUrl, shortUrl)

	// Once deployed this url will be changed
	host := "http://localhost:9098/re/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	if cache.IsPresentInCache(shortUrl) {
		log.Printf("Redis Cache Hit Key: %v\n", shortUrl)
		originalUrl, err := cache.RetrieveOriginalUrl(shortUrl)
		if err == nil {
			log.Printf("Redirecting to original url: %v", originalUrl)
			c.Redirect(302, originalUrl)
		}
	} else {
		log.Printf("Redis Cache Miss Key: %v\n", shortUrl)
		originalUrl, err := dbservice.GetOriginalUrl(shortUrl)
		if err == nil {
			c.Redirect(302, originalUrl)
		}
	}
}
