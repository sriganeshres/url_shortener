package db

import (
	"time"
)

func SaveURL(originalURL, shortCode string, custom bool) error {
	_, err := PG.Exec(`INSERT INTO urls (original_url, short_code, custom_alias) VALUES ($1, $2, $3)`, originalURL, shortCode, custom)
	if err == nil {
		RDB.Set(ctx, shortCode, originalURL, 24*time.Hour)
	}
	return err
}

func GetOriginalURL(shortCode string) (string, error) {
	url, err := RDB.Get(ctx, shortCode).Result()
	if err == nil {
		return url, nil
	}

	var original string
	err = PG.QueryRow(`SELECT original_url FROM urls WHERE short_code=$1`, shortCode).Scan(&original)
	if err == nil {
		RDB.Set(ctx, shortCode, original, 24*time.Hour)
	}
	return original, err
}
