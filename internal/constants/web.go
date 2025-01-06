package constants

import (
	"net/url"
	"os"
)

var (
	FrontendURL     = MustGetEnv("FRONTEND_URL")
	FrontendHost    string
	FrontendAppName = MustGetEnv("WEBAPP_NAME")
)

var (
	ServePrivateEndpointsEnabled = os.Getenv("PRIVATE_SERVER_ENABLED") == "true"
	ServePrivateEndpointsPort    = os.Getenv("PRIVATE_SERVER_PORT")
)

func init() {
	var frontendURL, _ = url.Parse(FrontendURL)
	FrontendHost = frontendURL.Host
}
