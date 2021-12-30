package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
	"strings"
)

func handleChat(_ *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	game.Broadcast(model.NewResponse("CHAT", client.Name, strings.Join(message, " ")))
	return nil
}
