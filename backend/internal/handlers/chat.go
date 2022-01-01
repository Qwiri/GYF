package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var ChatHandler = &Handler{
	AccessLevel: AccessJoined,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  model.StateAny,
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		msg := strings.TrimSpace(strings.Join(message, " "))
		if msg == "" {
			return gerrors.ErrMessageEmpty
		}
		game.Broadcast(model.PChat(client, msg))
		return nil
	}),
}
