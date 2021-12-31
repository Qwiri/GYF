package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var ListHandler = &Handler{
	AccessLevel: AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  model.StateAny,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
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
		return model.NewResponse("LIST", clientNames...).Respond(conn)
	}),
}
