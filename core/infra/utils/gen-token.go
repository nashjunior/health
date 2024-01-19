package utils

import (
	"math/rand"
	"time"
)

func GenerateCode(size int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define the characters allowed in the random string
	const charset = "0123456789"

	result := make([]byte, size)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)

}
