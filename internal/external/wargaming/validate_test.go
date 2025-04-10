package wargaming

import (
	"testing"

	"github.com/matryer/is"
)

func TestValidatePlayerNickname(t *testing.T) {
	is := is.New(t)

	cases := map[string]bool{
		"Valid_Nickname": true,
		"valid_nickname": true,
		"123_nickname":   true,
		"_123_nickname_": true,

		"bang!":           false,
		"nickname server": false,
		"nickname-server": false,
		"nickname@server": false,
	}
	for test, expected := range cases {
		is.Equal(ValidatePlayerNickname(test), expected)
	}
}
