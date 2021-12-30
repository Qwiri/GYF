package handlers

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"strings"
	"time"
)

var (
	ErrMessageTooShort = errors.New("message too short")
	ErrNotJoined       = errors.New("not joined")
	ErrUnknownCommand  = errors.New("unknown command")
)

type HandlerFunc func(conn *websocket.Conn, game *model.Game, client *model.Client, prefix string, message []string) error

func OnClientMessage(conn *websocket.Conn, game *model.Game, msg string) error {
	game.LastInteraction = time.Now()

	str := strings.Split(msg, " ")
	prefix := strings.ToUpper(str[0])

	// get client if the client already exists
	client := game.GetClient(conn)

	// Special Message: "WHOAMI"
	if prefix == "WHOAMI" {
		return handleWhoAmI(conn, game, client, prefix, str)
	}

	// commands below require extra args
	if len(str) <= 1 {
		log.WithField("message", str).Warn("message of client %+v too short")
		return ErrMessageTooShort
	}

	// Special Message: "JOIN"
	if prefix == "JOIN" {
		return handleJoin(conn, game, client, prefix, str[1:])
	}

	// make sure we know the client
	if client == nil {
		log.WithField("message", msg).Warn("got message from client which didn't join")
		return ErrNotJoined
	}
	client.LastInteraction = time.Now()

	switch prefix {
	case "LIST":
		return handleList(conn, game, client, prefix, str[1:])
	case "CHAT":
		return handleChat(conn, game, client, prefix, str[1:])
	case "TOPIC_LIST":
		return handleTopicList(conn, game, client, prefix, str[1:])
	case "TOPIC_ADD":
		return handleTopicAdd(conn, game, client, prefix, str[1:])
	case "TOPIC_REMOVE":
		return handleTopicRemove(conn, game, client, prefix, str[1:])
	}

	return ErrUnknownCommand
}
