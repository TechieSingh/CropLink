package cache

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/harivilasp/url-shortner-go/constants"
)

type CacheService struct {
	client *redis.Client
}

var (
	cacheService = &CacheService{}
)

func InitializeCacheService() *CacheService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("Error initializaing Redis Client %v", err)
		return nil
	}
	log.Printf("Redis Client Initialized Successfully! %v\n", pong)
	cacheService.client = redisClient
	return cacheService
}

func IsPresentInCache(url string) bool {
	_, err := cacheService.client.Get(url).Result()
	return err == nil
}

func SaveUrlMapping(originalUrl string, shortUrl string) error {
	err := cacheService.client.Set(shortUrl, originalUrl, constants.RedisCacheDuration).Err()
	if err != nil {
		return errors.New("error saving key-value to redis")
	}
	log.Printf("Saved key-value %v - %v\n", shortUrl, originalUrl)
	return nil
}

func RetrieveOriginalUrl(shortUrl string) (string, error) {
	if !IsPresentInCache(shortUrl) {
		log.Printf("Redis Caced: url %v not found\n", shortUrl)
		return "", errors.New("404: key not found")
	}
	val, err := cacheService.client.Get(shortUrl).Result()
	if err != nil {
		return "", errors.New("unknown error occured when retrieving original url")
	}
	return val, nil
}
