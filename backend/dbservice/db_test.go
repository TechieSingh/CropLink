package dbservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitializeDB()
}

func TestDBInit(t *testing.T) {
	assert.True(t, dbService != nil)
}

func TestInsertIntoDB(t *testing.T) {
	// Delete all the data that has been inserted while previous test methods
	TruncateDB()

	shortUrl := "short_url"
	longUrl := "long_url"

	err := InsertUrl(shortUrl, longUrl)
	assert.True(t, err == nil)
}

func TestRetrieveFromDB(t *testing.T) {
	// Delete all the data that has been inserted while previous test methods
	TruncateDB()

	shortUrl := "short_url"
	longUrl := "long_url"

	err := InsertUrl(shortUrl, longUrl)
	assert.True(t, err == nil)

	retrievedUrl, err := GetOriginalUrl(shortUrl)

	assert.True(t, err == nil)
	assert.True(t, retrievedUrl == longUrl)
}
