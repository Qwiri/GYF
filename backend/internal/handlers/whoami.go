package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

func handleWhoAmI(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	util.Respond(conn, model.NewResponse("WHOAMI", "YOU", client, "GAME", game, "MSG", message))
	return nil
}
