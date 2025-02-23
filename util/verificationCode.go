package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func GenerateVerificationCode() (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var code strings.Builder

	for i := 0; i < 8; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error generating verification code: %v", err)
		}
		code.WriteByte(charset[index.Int64()])
	}

	return code.String(), nil
}
