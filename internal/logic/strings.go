package logic

import (
	"crypto/rand"
	"fmt"
)

func RandomString(length int) (string, error) {
	data := make([]byte, length)
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", data), nil
}
