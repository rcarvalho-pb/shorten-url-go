package data

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rcarvalho-pb/link-shorter/internal/generator"
)

type ShortURL struct {
	ID        int       `db:"id"`
	FullURL   string    `db:"full_url"`
	ShortURL  string    `db:"short_url"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *ShortURL) NewShortURL(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	param := generator.RandSeq()
	newUrl := fmt.Sprintf("https://xxx.com/%s", param)

	log.Println(url)
	log.Println(newUrl)

	stmt := "INSERT INTO urls (full_url, short_url) VALUES ($1, $2)"

	_, err := db.ExecContext(ctx, stmt, url, newUrl)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint") {
			query := "SELECT * FROM urls WHERE full_url = $1"
			row := db.QueryRowContext(ctx, query, url)

			var shortUrl ShortURL

			if err = row.Scan(
				&shortUrl.ID,
				&shortUrl.FullURL,
				&shortUrl.ShortURL,
				&shortUrl.CreatedAt,
			); err != nil {
				return "", err
			}

			return shortUrl.ShortURL, nil

		}
	}

	log.Println("Inserted")

	return newUrl, nil
}

func (u *ShortURL) DeleteOldURLs() {
	stmt := "DELETE FROM urls WHERE (created_at < DATETIME('now', '-1 MINUTE'))"

	for {
		log.Println("Cleaning db...")
		if _, err := db.Exec(stmt); err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
		continue
	}
}
