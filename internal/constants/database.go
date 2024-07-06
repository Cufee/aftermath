package constants

var (
	DatabasePath = mustGetEnv("DATABASE_PATH")
	DatabaseName = mustGetEnv("DATABASE_NAME")
)
