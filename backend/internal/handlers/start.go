package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

const (
	MinTopic   = 1
	MinPlayers = 3
)

var StartHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		if game.Started {
			return gerrors.ErrGameStarted
		}
		// check if we have enough topics
		if len(game.Topics) < MinTopic {
			return gerrors.ErrTooFewTopics
		}
		// check if we have enough players
		if len(game.Clients) < MinPlayers {
			return gerrors.ErrTooFewPlayers
		}
		game.Started = true
		// start next round
		return game.NextRound()
	}),
}

var SkipHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	DevOnly:     true,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		if !game.Started {
			return gerrors.ErrGameNotStarted
		}
		return game.NextRound()
	}),
}
