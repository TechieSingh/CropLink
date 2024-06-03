package constants

import (
	"time"
)

// Redis Constants
const (
	RedisAddr          = "localhost:6379"
	RedisCacheDuration = 6 * time.Hour
)

// DB Constants
const (
	DB_USER     = "admin"
	DB_PASSWORD = "admin"
	DB_NAME     = "tiny_urls"
)
