package utils

import (
	"math/rand"
	"time"
)

func GenerateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(charset))
		result = append(result, charset[idx])
	}

	return string(result)
}
