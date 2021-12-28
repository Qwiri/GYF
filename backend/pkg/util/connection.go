package util

import (
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
)

func CloseConnection(c *websocket.Conn) {
	if err := c.Close(); err != nil {
		log.WithError(err).Warn("[ws] cannot close connection to client. but yolo")
	}
}

func Write(c *websocket.Conn, msg string) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.WithError(err).Warnf("[ws] cannot send %s to client", msg)
	}
}
