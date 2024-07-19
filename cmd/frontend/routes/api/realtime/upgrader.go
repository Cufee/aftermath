package realtime

import (
	"net/http"
	"net/url"
	"time"

	"github.com/cufee/aftermath/internal/constants"
	"github.com/gorilla/websocket"
)

var frontendURL, _ = url.Parse(constants.FrontendURL)

var defaultUpgrader = &websocket.Upgrader{
	HandshakeTimeout: time.Millisecond * 500,
	WriteBufferSize:  1024,
	ReadBufferSize:   1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Host == frontendURL.Host || constants.DevMode
	},
}
