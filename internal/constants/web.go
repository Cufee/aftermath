package constants

import (
	"net/url"
	"os"
)

var (
	FrontendURL     = mustGetEnv("FRONTEND_URL")
	FrontendHost    string
	FrontendAppName = mustGetEnv("WEBAPP_NAME")
)

var (
	ServePrivateEndpointsEnabled = os.Getenv("PRIVATE_SERVER_ENABLED") == "true"
	ServePrivateEndpointsPort    = os.Getenv("PRIVATE_SERVER_PORT")
)

func init() {
	var frontendURL, _ = url.Parse(FrontendURL)
	FrontendHost = frontendURL.Host
}
