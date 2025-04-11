package wargaming

import "unicode"

// Returns true if a nickname is between 3 and 24 characters, uses Latin characters, numbers, and underscores only.
//
// https://worldoftanks.com/en/content/guide/general/username_change/
func ValidatePlayerNickname(input string) bool {
	if l := len(input); l < 3 || l > 24 {
		return false
	}
	for _, r := range input {
		if !(unicode.IsUpper(r) && r <= unicode.MaxASCII ||
			unicode.IsLower(r) && r <= unicode.MaxASCII ||
			unicode.IsDigit(r) ||
			r == '_') {
			return false
		}
	}
	return true
}
