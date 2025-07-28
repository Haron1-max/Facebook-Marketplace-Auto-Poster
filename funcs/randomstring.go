package funcs

import (
	"crypto/rand"
	"math/big"
)

// Function to generate a cryptographically secure random string of a given length
func GenerateSecureRandomName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[randomInt.Int64()]
	}
	return string(b)
}
