package handlers

import (
	"github.com/Qwiri/GYF/backend/internal/stringify"
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var ExplainHandler = &handler.Handler{
	Description: "Returns help for a handler",
	Syntax:      "(handler!)",
	AccessLevel: handler.AccessAny,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateAny,
	Handler: handler.HandlersHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, msg []string, handlers map[string]*handler.Handler) error {
		name := strings.ToUpper(msg[0])
		h, ok := handlers[name]
		if !ok {
			return gerrors.ErrInvalidHandler
		}
		var bob util.Bob
		stringify.HandlerToString(&bob, name, h)
		return model.NewResponse("EXPLAIN", util.WrapStringArray(strings.Split(bob.String(), "\n")...)...).Respond(conn)
	}),
	DevOnly: true,
}
