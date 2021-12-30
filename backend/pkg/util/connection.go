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
