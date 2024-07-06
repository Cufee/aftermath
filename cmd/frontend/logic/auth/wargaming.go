package auth

import "os"

var baseRedirectURL = os.Getenv("WG_AUTH_REDIRECT_URI")

func init() {
	if baseRedirectURL == "" {
		panic("WG_AUTH_REDIRECT_URI missing")
	}
}

func NewWargamingAuthRedirectURL(token string) string {
	return baseRedirectURL + "/" + token
}
