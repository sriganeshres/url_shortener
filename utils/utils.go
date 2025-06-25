package utils

import (
    "math/rand"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShortCode() string {
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}