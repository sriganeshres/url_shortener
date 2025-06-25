package utils

import (
    "math/rand"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
 * @brief GenerateShortCode generates a random 6-character short code.
 *
 * This function generates a random 6-character short code using a combination of
 * uppercase and lowercase letters, as well as digits.
 *
 * @return The generated short code.
 */
func GenerateShortCode() string {
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}