package util

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
)

func CloseConnection(c *websocket.Conn) {
	if err := c.Close(); err != nil {
		log.WithError(err).Warn("[ws] cannot close connection to client. but yolo")
	}
}

func Respond(conn *websocket.Conn, response *model.Response) {
	if conn == nil {
		log.Warnf("tried to send '%s' to nil connection", response)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, response.Marshal()); err != nil {
		log.WithError(err).Warnf("[ws] cannot send %s to client", response.String())
	}
}
