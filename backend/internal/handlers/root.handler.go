package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"regexp"
	"strings"
	"time"
)

var SpacesRegEx = regexp.MustCompile(`\s+`)

type Handler struct {
	AccessLevel Access
	Bounds      util.Boundaries
	StateLevel  model.GameState
	Handler     interface{}
	DevOnly     bool
}

var Handlers = map[string]*Handler{
	"WHOAMI":       WhoAmIHandler,
	"JOIN":         JoinHandler,
	"LIST":         ListHandler,
	"CHAT":         ChatHandler,
	"TOPIC_LIST":   TopicListHandler,
	"TOPIC_ADD":    TopicAddHandler,
	"TOPIC_REMOVE": TopicRemoveHandler,
	"START":        StartHandler,
	"SKIP":         SkipHandler,
	"SUBMIT_GIF":   SubmitGIFHandler,
	"VOTE":         VoteCastHandler,
	"NEXT_ROUND":   NextRoundHandler,
	"STATS":        StatsHandler,
}

type BasicHandler func(*websocket.Conn, *model.Game, *model.Client) error
type MessagedHandler func(*websocket.Conn, *model.Game, *model.Client, []string) error
type PrefixedHandler func(*websocket.Conn, *model.Game, *model.Client, string) error
type PrefixedMessagedHandler func(*websocket.Conn, *model.Game, *model.Client, string, []string) error

func OnClientMessage(conn *websocket.Conn, game *model.Game, msg string, devMode bool) error {
	game.LastInteraction = time.Now() // update game's last interaction for janitor grace

	msg = strings.TrimSpace(msg)
	if len(msg) == 0 {
		return gerrors.ErrMessageTooShort
	}

	str := SpacesRegEx.Split(msg, -1)
	prefix := strings.ToUpper(str[0])

	// get client if the client already exists
	// NOTICE: client CAN BE nil AT THIS MOMENT!
	client := game.GetClient(conn)

	// update client's last interaction for janitor grace
	if client != nil {
		client.LastInteraction = time.Now()
	}

	// find handler
	handler, ok := Handlers[prefix]
	if !ok {
		return gerrors.ErrUnknownCommand
	}

	// check if handler is dev-only
	if handler.DevOnly && !devMode {
		return gerrors.ErrDevOnly
	}

	// check access for handler
	if !handler.AccessLevel.Allowed(client) {
		return gerrors.ErrNoAccess
	}

	// check arg length
	if handler.Bounds != nil {
		if !handler.Bounds.Applies(len(str[1:])) {
			return gerrors.ErrArgLength
		}
	}

	// check game start
	if !handler.StateLevel.Allowed(game) {
		return gerrors.ErrGameStateAccess
	}

	var err error

	// execute handler
	switch hdl := handler.Handler.(type) {
	case BasicHandler:
		err = hdl(conn, game, client)
	case MessagedHandler:
		err = hdl(conn, game, client, str[1:])
	case PrefixedHandler:
		err = hdl(conn, game, client, prefix)
	case PrefixedMessagedHandler:
		err = hdl(conn, game, client, prefix, str[1:])
	default:
		log.Warnf("cannot find handler for %s", prefix)
		err = gerrors.ErrInvalidHandler
	}

	if err != nil {
		// send error
		err = model.NewResponseWithError(prefix, err).Respond(conn)
	}

	return err
}
