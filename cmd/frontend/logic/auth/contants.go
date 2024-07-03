package auth

import "os"

const SessionCookieName = "x-amth-session"
const AuthNonceCookieName = "x-amth-auth-nonce"

var (
	DefaultCookiePath   = os.Getenv("AUTH_COOKIE_PATH")
	DefaultCookieDomain = os.Getenv("AUTH_COOKIE_DOMAIN")
)

func init() {
	if DefaultCookiePath == "" {
		panic("AUTH_COOKIE_PATH cannot be left empty")
	}
	if DefaultCookieDomain == "" {
		panic("AUTH_COOKIE_DOMAIN cannot be left empty")
	}
}
