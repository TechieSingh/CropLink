package shortener

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// Used to create hash of original big url
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encodedData, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encodedData)
}

func GenerateShortUrl(initialUrl string, userId string) string {
	urlHashBytes := sha256Of(initialUrl + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if len(finalString) < 8 {
		panic(errors.New("len of shortened url string is < 8"))
	}
	return finalString[:8]
}
