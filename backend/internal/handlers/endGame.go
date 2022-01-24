package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var EndGameHandler = &handler.Handler{
	Description: "Ends the Game (skips to the result screen)",
	Syntax:      "",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateInGame,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return game.ForceEndGame("Ended by Host")
	}),
}
