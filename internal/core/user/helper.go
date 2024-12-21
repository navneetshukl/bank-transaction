package user

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateUniqueRandomValue(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)
	if len(randomString) > length {
		randomString = randomString[:length]
	}

	return randomString, nil
}
