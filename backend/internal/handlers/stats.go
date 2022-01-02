package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var StatsHandler = &handler.Handler{
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateInGame & ^util.StateCastVotes,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.PStats(game).Respond(conn)
	}),
}
