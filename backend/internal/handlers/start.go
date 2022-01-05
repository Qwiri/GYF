package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var StartHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateLobby,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		// check if we have enough topics
		if len(game.Topics) < game.Preferences.MinTopics {
			return gerrors.ErrTooFewTopics
		}
		// check if we have too many topics
		if len(game.Topics) > game.Preferences.MaxTopics {
			return gerrors.ErrTooManyTopics
		}
		// check if we have enough players
		if len(game.Clients) < game.Preferences.MinPlayers {
			return gerrors.ErrTooFewPlayers
		}
		// check if we have too many players
		if len(game.Clients) > game.Preferences.MaxPlayers {
			return gerrors.ErrTooManyPlayers
		}
		// start next round
		game.Broadcast(model.PStart())
		return game.ForceNextRound()
	}),
}

var SkipHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateInGame,
	DevOnly:     true,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return game.CheckCycle(false, true)
	}),
}

var NextRoundHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateShowVotes,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return game.ForceNextRound()
	}),
}
