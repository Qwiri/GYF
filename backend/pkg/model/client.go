package model

import (
	"github.com/gofiber/websocket/v2"
	"time"
)

type Client struct {
	Name            string
	Connection      *websocket.Conn
	Leader          bool
	LastInteraction time.Time
}

func NewClient(conn *websocket.Conn, username string) (c *Client) {
	return &Client{
		Name:            username,
		Connection:      conn,
		Leader:          false,
		LastInteraction: time.Now(),
	}
}
