package constants

import "fmt"

var (
	DatabaseConnString = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", MustGetEnv("DATABASE_USER"), MustGetEnv("DATABASE_PASSWORD"), MustGetEnv("DATABASE_HOST"), MustGetEnv("DATABASE_NAME"))
)
