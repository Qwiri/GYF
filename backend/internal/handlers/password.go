package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
)

var ChangePasswordHandler = &handler.Handler{
	Description: "Changes the lobby password",
	Syntax:      "(password?)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMax(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, args []string) error {
		var newPass string
		if len(args) > 0 {
			newPass = args[0]
		}
		game.Password = newPass
		log.Debugf("changed password for game '%s' to '%s'", game.ID, newPass)
		return model.NewResponse("CHANGE_PASS").Respond(conn)
	}),
}
