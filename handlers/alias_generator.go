package handlers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomAlias(length int) (string, error) {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	alias := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", fmt.Errorf("GenerateRandomAlias:error &w", err)
		}
		alias[i] = charset[randomIndex.Int64()]
	}
	return string(alias), nil
}
