package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var WhoAmIHandler = &Handler{
	AccessLevel: AccessGuest | AccessJoined | AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  model.StateAny,
	DevOnly:     true,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.NewResponse("WHOAMI", "YOU", client, "GAME", game).Respond(conn)
	}),
}
