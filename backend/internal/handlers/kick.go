package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var KickPlayerHandler = &handler.Handler{
	Description: "Kicks a player from the game",
	Syntax:      "(player!)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, msg []string) error {
		// find client by name
		kick, ok := game.Clients[msg[0]]
		if !ok {
			return gerrors.ErrClientNotFound
		}
		if kick == client {
			return gerrors.ErrSelf
		}
		// kick the fucker
		game.KickClient(kick)
		return nil
	}),
}
