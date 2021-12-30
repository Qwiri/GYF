package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
)

func handleWhoAmI(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	model.NewResponse("WHOAMI", "YOU", client, "GAME", game, "MSG", message).Respond(conn)
	return nil
}
