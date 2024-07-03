package auth

import (
	"net/http"
	"time"
)

func NewSessionCookie(value string, expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:   SessionCookieName,
		Value:  value,
		MaxAge: int(time.Until(expiresAt).Seconds()),

		Secure:   true,
		HttpOnly: true,
		Path:     DefaultCookiePath,
		Domain:   DefaultCookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
}

func NewNonceCookie(value string, expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:   AuthNonceCookieName,
		Value:  value,
		MaxAge: int(time.Until(expiresAt).Seconds()),

		Secure:   true,
		HttpOnly: true,
		Path:     DefaultCookiePath,
		Domain:   DefaultCookieDomain,
		SameSite: http.SameSiteLaxMode,
	}
}
