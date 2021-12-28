package main

import (
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"math/rand"
	"regexp"
	"strings"
)

///

const CharSet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789"

func GenerateRandomString(length int) string {
	var bob strings.Builder
	for i := 0; i < length; i++ {
		bob.WriteRune(rune(CharSet[rand.Intn(len(CharSet))]))
	}
	return bob.String()
}

///

var usernameExpr = regexp.MustCompile("^[A-Za-z0-9]{1,16}$")

func IsUserNameValid(username string) bool {
	return usernameExpr.MatchString(username)
}

///

func closeConnection(c *websocket.Conn) {
	if err := c.Close(); err != nil {
		log.WithError(err).Warn("[ws] cannot close connection to client. but yolo")
	}
}

func write(c *websocket.Conn, msg string) {
	if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.WithError(err).Warnf("[ws] cannot send %s to client", msg)
	}
}

func (g *Game) broadcast(msg string) {
	for _, client := range g.Clients {
		write(client.Connection, msg)
	}
}

func (g *Game) broadcastExcept(conn *websocket.Conn, msg string) {
	for _, client := range g.Clients {
		if client.Connection != conn {
			write(client.Connection, msg)
		}
	}
}

func (g *Game) GetClient(conn *websocket.Conn) *Client {
	for _, c := range g.Clients {
		if c.Connection == conn {
			return c
		}
	}
	return nil
}
