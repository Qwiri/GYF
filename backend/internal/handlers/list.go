package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
)

func handleList(conn *websocket.Conn, game *model.Game, _ *model.Client, _ string, _ []string) error {
	type listObj struct {
		Name   string `json:"name"`
		Leader bool   `json:"leader"`
	}
	// collect client names
	clientNames := make([]interface{}, len(game.Clients))
	var i = 0
	for _, c := range game.Clients {
		clientNames[i] = listObj{c.Name, c.Leader}
		i += 1
	}
	model.NewResponse("LIST", clientNames...).Respond(conn)
	return nil
}
