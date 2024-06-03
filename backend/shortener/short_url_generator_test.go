package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortUrlGenerator(t *testing.T) {

	initialLink_1 := "https://leetcode.com/problemset/all/"
	shortLink_1 := "HPWqVZ1j"

	initialLink_2 := "https://admission.asu.edu/international/student-visa"
	shortLink_2 := "QiFVRpmY"

	initialLink_3 := "https://aayushimalviya.wordpress.com/2021/08/20/googles-hiring-process-for-software-engineering/"
	shortLink_3 := "SbhXm5AK"

	assert.True(t, shortLink_1 == GenerateShortUrl(initialLink_1, UserId))
	assert.True(t, shortLink_2 == GenerateShortUrl(initialLink_2, UserId))
	assert.True(t, shortLink_3 == GenerateShortUrl(initialLink_3, UserId))

}
