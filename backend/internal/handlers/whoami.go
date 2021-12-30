package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
)

var WhoAmIHandler = &Handler{
	AccessLevel: AccessGuest | AccessJoined | AccessLeader,
	DevOnly:     true,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.NewResponse("WHOAMI", "YOU", client, "GAME", game).Respond(conn)
	}),
}
