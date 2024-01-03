package method

import (
	"math/rand"
	"time"
)

// GenerateRandomString generate rant string
func GenerateRandomString(length int) string {

	rand.NewSource(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	charsetLength := len(charset)
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(charsetLength)]
	}

	return string(randomString)
}
