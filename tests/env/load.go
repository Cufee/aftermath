package env

import (
	"path/filepath"
	"testing"

	"github.com/cufee/aftermath/tests/path"
	"github.com/joho/godotenv"
)

func LoadTestEnv(t *testing.T) {
	env, err := godotenv.Read(filepath.Join(path.Root(), ".env.test"))
	if err != nil {
		panic(err)
	}
	for k, v := range env {
		t.Setenv(k, v)
	}
}
