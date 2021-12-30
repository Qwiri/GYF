package handlers

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var ErrMessageEmpty = errors.New("message empty")

var ChatHandler = &Handler{
	AccessLevel: AccessJoined,
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		msg := strings.TrimSpace(strings.Join(message, " "))
		if msg == "" {
			return ErrMessageEmpty
		}
		game.Broadcast(model.NewResponse("CHAT", client.Name, msg))
		return nil
	}),
}
