package logic

import (
	"crypto/sha256"
	"fmt"

	"github.com/cufee/aftermath/internal/encoding"
)

/*
returns a sha256 hash of the input string
*/
func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

/*
gob encode the input and sha256 hashes the resulting string
*/
func HashAny(input any) (string, error) {
	data, err := encoding.EncodeGob(input)
	if err != nil {
		return "", err
	}
	return HashString(string(data)), nil
}
