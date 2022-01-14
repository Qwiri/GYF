package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strconv"
)

var ChangePermissionHandler = &handler.Handler{
	Description: "Changes the enhanced permission for the game",
	Syntax:      "(permission!)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) (err error) {
		var value int
		if value, err = strconv.Atoi(message[0]); err != nil {
			return
		}
		game.Preferences.Permissions = model.EnhancedPermission(uint8(value))
		game.Broadcast(model.PPreferences(game.Preferences))
		game.BroadcastTopicListToLeaders()
		return nil
	}),
}
