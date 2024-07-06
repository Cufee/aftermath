package constants

import "os"

var (
	FrontendURL     = mustGetEnv("FRONTEND_URL")
	FrontendAppName = mustGetEnv("WEBAPP_NAME")
)

var (
	ServePrivateEndpointsEnabled = os.Getenv("PRIVATE_SERVER_ENABLED") == "true"
	ServePrivateEndpointsPort    = os.Getenv("PRIVATE_SERVER_PORT")
)
