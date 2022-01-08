package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var WhoAmIHandler = &handler.Handler{
	Description: "Returns information about the client and game",
	Syntax:      "",
	AccessLevel: handler.AccessAny,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateAny,
	DevOnly:     true,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.NewResponse("WHOAMI", "YOU", client, "GAME", game).Respond(conn)
	}),
}
