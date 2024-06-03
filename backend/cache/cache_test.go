package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	cacheService = InitializeCacheService()
}

func TestCacheStoreInit(t *testing.T) {
	assert.True(t, cacheService != nil)
}

func TestUrlPresentInCache(t *testing.T) {
	originalUrl := "some_url.com"
	shortUrl := "short_url"
	err := SaveUrlMapping(originalUrl, shortUrl)
	assert.True(t, err == nil)
	assert.True(t, IsPresentInCache(shortUrl))
}

func TestRetrieveOriginalUrlWithError(t *testing.T) {
	url := "some_unsaved_url.com"
	_, err := RetrieveOriginalUrl(url)
	assert.True(t, err != nil)
}

func TestRetrieveOriginalUrlWithoutError(t *testing.T) {
	originalUrl := "saving_new_url.com"
	shortUrl := "short_url_sample.com"

	// Save the Url in cache
	saveErr := SaveUrlMapping(originalUrl, shortUrl)

	assert.True(t, saveErr == nil)

	// Retrieve Url from cache
	retrievedUrl, err := RetrieveOriginalUrl(shortUrl)

	assert.True(t, err == nil)
	assert.True(t, retrievedUrl == originalUrl)
}
