package db

import (
	"time"
)

/*
  * @brief SaveURL saves a URL and its short code to the database and cache.
  *
  * This function saves a URL and its short code to the database and cache. If the
  * short code is custom, it is stored in the database with the custom_alias flag
  * set to true. The URL is also cached with a 24 hour expiration time.
  *
  * @param originalURL The original URL.
  * @param shortCode The short code for the URL.
  * @param custom Whether the short code is custom.
  *
  * @return An error if the save operation fails.
*/
func SaveURL(originalURL, shortCode string, custom bool) error {
	_, err := PG.Exec(`INSERT INTO urls (original_url, short_code, custom_alias) VALUES ($1, $2, $3)`, originalURL, shortCode, custom)
	if err == nil {
		RDB.Set(ctx, shortCode, originalURL, 24*time.Hour)
	}
	return err
}

/*
 * @brief GetOriginalURL retrieves the original URL for a given short code.
 *
 * This function attempts to retrieve the original URL from the cache using the provided
 * short code. If the URL is not found in the cache, it queries the database for the original
 * URL, caches the result, and returns it. If the URL cannot be found in both the cache and
 * the database, an error is returned.
 *
 * @param shortCode The short code for which to retrieve the original URL.
 *
 * @return The original URL if found, otherwise an error.
 */
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
