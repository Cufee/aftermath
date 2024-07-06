package constants

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func mustGetEnv(key string, ignoreIf ...func() bool) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	for _, ignore := range ignoreIf {
		if ignore() {
			return "@ignored"
		}
	}
	panic(key + " environment variable must be set")
}
