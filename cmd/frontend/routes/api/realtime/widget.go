package realtime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cufee/aftermath/cmd/frontend/handler"
	"github.com/cufee/aftermath/internal/realtime"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var widgetWsMaxMessageSize int64 = 512
var widgetWsPongWait = time.Second * 60
var widgetWsPingPeriod = (widgetWsPongWait * 9) / 10

var WidgetSettings handler.WebSocket = func(ctx *handler.Context) (*websocket.Upgrader, func(conn *websocket.Conn) error, error) {
	widgetId := ctx.Path("widgetId")
	if widgetId == "" {
		ctx.SetStatus(http.StatusBadRequest)
		return nil, nil, nil
	}

	settings, err := ctx.Database().GetWidgetSettings(ctx.Context, widgetId)
	if err != nil {
		ctx.SetStatus(http.StatusNotFound)
		return nil, nil, nil
	}

	topicID := fmt.Sprintf("widget-settings-%s", settings.ID)
	err = ctx.PubSub().NewTopic(topicID)
	if err != nil && !errors.Is(err, realtime.ErrTopicRegistered) {
		ctx.SetStatus(http.StatusInternalServerError)
		return nil, nil, ctx.String("failed to register a pubsub topic")
	}

	listenCh, closeListener, err := ctx.PubSub().Listen(topicID)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return nil, nil, ctx.String("failed to register a pubsub listener")
	}

	return defaultUpgrader, func(conn *websocket.Conn) error {
		data := map[string]any{"widgetId": settings.ID}
		err = conn.WriteJSON(data)
		if err != nil {
			defer conn.Close()
			return conn.WriteMessage(websocket.TextMessage, []byte("failed to write message: "+err.Error()))
		}

		// listen to messages
		go func(conn *websocket.Conn) {
			defer conn.Close()
			defer closeListener()

			conn.SetReadLimit(widgetWsMaxMessageSize)
			conn.SetReadDeadline(time.Now().Add(widgetWsPongWait))
			conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(widgetWsPongWait)); return nil })

			for {
				var data map[string]any
				err := conn.ReadJSON(&data)
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						return
					}
					_ = conn.WriteMessage(websocket.TextMessage, []byte("invalid message: "+err.Error()))
					return
				}

				// we have nothing to do here rn
			}
		}(conn)

		// send messages
		go func(conn *websocket.Conn) {
			ticker := time.NewTicker(widgetWsPingPeriod)

			defer func() {
				ticker.Stop()
				closeListener()
				_ = conn.WriteMessage(websocket.TextMessage, []byte("topic listener closed"))
				_ = conn.Close()
			}()

			for {
				select {
				case <-ticker.C:
					conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
					if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}

				case data, ok := <-listenCh:
					if !ok {
						conn.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}

					command, ok := data.Data.(string)
					if !ok {
						continue
					}

					conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
					w, err := conn.NextWriter(websocket.TextMessage)
					if err != nil {
						return
					}

					payload, _ := json.Marshal(map[string]any{"command": command})
					w.Write(payload)
					if err := w.Close(); err != nil {
						return
					}
				}
			}
		}(conn)

		return nil
	}, nil
}
