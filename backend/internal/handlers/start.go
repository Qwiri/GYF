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
	StateLevel:  model.StateLobby,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		// check if we have enough topics
		if len(game.Topics) < MinTopic {
			return gerrors.ErrTooFewTopics
		}
		// check if we have enough players
		if len(game.Clients) < MinPlayers {
			return gerrors.ErrTooFewPlayers
		}
		// start next round
		return game.ForceNextRound()
	}),
}

var SkipHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  model.StateInGame,
	DevOnly:     true,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		switch game.State {
		case model.StateSubmitGIF:
			return game.ForceStartVote()
		case model.StateCastVotes:
			return game.ForceShowVoteResults()
		case model.StateShowVotes:
			return game.ForceNextRound()
		}
		return gerrors.ErrUnknownState
	}),
}

var NextRoundHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  model.StateShowVotes,
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return game.ForceNextRound()
	}),
}
