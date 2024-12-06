package env

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

func LoadTestEnv(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	env, err := godotenv.Read(filepath.Join(basepath, "../../.env.test"))
	if err != nil {
		panic(err)
	}
	for k, v := range env {
		t.Setenv(k, v)
	}
}
