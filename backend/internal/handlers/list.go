package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

func handleList(conn *websocket.Conn, game *model.Game, _ *model.Client, _ string, _ []string) error {
	// collect client names
	clientNames := make([]interface{}, len(game.Clients))
	var i = 0
	for _, c := range game.Clients {
		clientNames[i] = c.Name
		i += 1
	}
	util.Respond(conn, model.NewResponse("LIST", clientNames...))
	return nil
}
