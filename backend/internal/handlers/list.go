package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var ListHandler = &handler.Handler{
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateAny,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.PList(game.Clients).Respond(conn)
	}),
}
