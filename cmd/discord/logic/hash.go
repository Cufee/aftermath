package logic

import (
	"encoding/json"

	"github.com/cufee/aftermath/internal/logic"
)

/*
json encode the input and sha256 hashes the resulting string
*/
func hashAny(input any) (string, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return logic.HashString(string(data)), nil
}
