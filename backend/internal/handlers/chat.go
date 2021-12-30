package handlers

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var ErrMessageEmpty = errors.New("message empty")

func handleChat(_ *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	msg := strings.TrimSpace(strings.Join(message, " "))
	if msg == "" {
		return ErrMessageEmpty
	}
	game.Broadcast(model.NewResponse("CHAT", client.Name, msg))
	return nil
}
