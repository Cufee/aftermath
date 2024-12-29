package constants

var (
	DatabasePath = MustGetEnv("DATABASE_PATH")
	DatabaseName = MustGetEnv("DATABASE_NAME")
)
