package wargaming

import "unicode"

// returns true if the input is a valid wargaming nickname
// https://worldoftanks.com/en/content/guide/general/username_change/#:~:text=Full%20Guide,-Full%20Guide&text=Whatever%20the%20reason%20may%20be,Click%20CHANGE.
// Must be between 3 and 24 characters. Use Latin characters, numbers, and underscores only.
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
