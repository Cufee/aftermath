package constants

import (
	"os"
	"strconv"
	"time"
)

var (
	BlitzStarsApiURL = mustGetEnv("BLITZ_STARS_API_URL")
)

var (
	WargamingPrimaryAppID             = mustGetEnv("WG_PRIMARY_APP_ID")
	WargamingPrimaryAppRPS            int
	WargamingPrimaryAppRequestTimeout time.Duration
	WargamingPrimaryAppProxyHostList  = os.Getenv("WG_PROXY_HOST_LIST")
	WargamingCacheAppID               = mustGetEnv("WG_CACHE_APP_ID")
	WargamingCacheAppRPS              int
	WargamingCacheAppRequestTimeout   time.Duration
	WargamingCacheAppProxyHostList    = os.Getenv("WG_CACHE_PROXY_HOST_LIST")

	WargamingPublicAppID = mustGetEnv("WG_AUTH_APP_ID")
)

func init() {
	{
		WargamingPrimaryAppRPS, _ = strconv.Atoi(mustGetEnv("WG_PRIMARY_APP_RPS"))
		timeoutSec, _ := strconv.Atoi(mustGetEnv("WG_PRIMARY_REQUEST_TIMEOUT_SEC"))
		WargamingPrimaryAppRequestTimeout = time.Duration(timeoutSec) * time.Second
	}

	{

		WargamingCacheAppRPS, _ = strconv.Atoi(mustGetEnv("WG_CACHE_APP_RPS"))
		timeoutSec, _ := strconv.Atoi(mustGetEnv("WG_CACHE_REQUEST_TIMEOUT_SEC"))
		WargamingCacheAppRequestTimeout = time.Duration(timeoutSec) * time.Second
	}

}
