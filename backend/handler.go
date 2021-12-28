package main

import (
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"strings"
	"time"
)

//
func onClientMessage(conn *websocket.Conn, game *Game, msg string) {
	str := strings.Split(msg, " ")
	if len(str) <= 1 {
		log.WithField("message", str).Warnf("message of client %+v too short")
		return
	}
	prefix := strings.ToUpper(str[0])

	// get client if the client already exists
	client := game.GetClient(conn)

	// Special Message: "JOIN"
	if prefix == "JOIN" {
		username := str[1]
		// check if username is allowed
		if !IsUserNameValid(username) {
			write(conn, "ERROR name invalid")
			return
		}
		// check if there is already a client with the same name?
		for _, c := range game.Clients {
			if strings.EqualFold(c.Name, username) {
				write(conn, "ERROR player already exists")
				return
			}
		}
		// client already joined?
		if client != nil {
			if client.Name != "" {
				// client has already a name
				write(conn, "ERROR already joined")
				return
			}
			client.Name = username
		} else {
			client = NewClient(conn, username)
			// add client to game map
			game.Clients[client.Name] = client
		}
		// broadcast player join
		game.broadcast("PLAYER_JOINED " + client.Name)
		log.Infof("Client %s joined game %s", client.Name, game.ID)
		return
	}

	// make sure we know the client
	if client == nil {
		log.WithField("message", msg).Warn("got message from client which didn't join")
		return
	}
	client.LastInteraction = time.Now()

	switch prefix {
	case "JOIN":
	}
}
