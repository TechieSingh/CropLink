package dbservice

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/harivilasp/url-shortner-go/constants"
	_ "github.com/lib/pq"
)

type DBService struct {
	db *sql.DB
}

var dbService = &DBService{}

func InitializeDB() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err.Error())
	}
	log.Println("DB Initialized")
	dbService.db = db
	return db
}

func GetOriginalUrl(shortUrl string) (string, error) {
	dbQuery := "SELECT long_url FROM urls WHERE short_url = $1;"
	var originalUrl string
	err := dbService.db.QueryRow(dbQuery, shortUrl).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func InsertUrl(shortUrl string, longUrl string) error {
	dbQuery := fmt.Sprintf("INSERT INTO urls (short_url, long_url) VALUES ('%s', '%s');", shortUrl, longUrl)
	_, err := dbService.db.Exec(dbQuery)
	if err != nil {
		return err
	}
	return nil
}

func TruncateDB() {
	dbQuery := "TRUNCATE urls;"
	_, err := dbService.db.Exec(dbQuery)
	if err != nil {
		panic(err.Error())
	}
}
