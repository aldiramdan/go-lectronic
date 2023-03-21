package libs

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
)

func CodeCrypt(len int) (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("could not has password %w", err)
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:len], nil
}