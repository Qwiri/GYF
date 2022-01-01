package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var StatsHandler = &Handler{
	AccessLevel: AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  model.StateInGame & ^model.StateCastVotes,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		stats := make(map[string]int)
		for _, c := range game.Clients {
			stats[c.Name] = game.GetStats(c.Name)
		}
		return model.NewResponse("STATS", stats).Respond(conn)
	}),
}