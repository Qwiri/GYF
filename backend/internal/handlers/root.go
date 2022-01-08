package handlers

import (
	"regexp"
	"strings"
	"time"

	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
)

const MaxPayloadLength = 4096

var SpacesRegEx = regexp.MustCompile(`\s+`)

var Handlers = map[string]*handler.Handler{
	"WHOAMI":        WhoAmIHandler,
	"JOIN":          JoinHandler,
	"LIST":          ListHandler,
	"CHAT":          ChatHandler,
	"TOPIC_LIST":    TopicListHandler,
	"TOPIC_ADD":     TopicAddHandler,
	"TOPIC_ADD_ALL": TopicAddAllHandler,
	"TOPIC_REMOVE":  TopicRemoveHandler,
	"TOPIC_CLEAR":   TopicClearHandler,
	"START":         StartHandler,
	"SKIP":          SkipHandler,
	"SUBMIT_GIF":    SubmitGIFHandler,
	"VOTE":          VoteCastHandler,
	"NEXT_ROUND":    NextRoundHandler,
	"STATS":         StatsHandler,
	"CHANGE_PREF":   ChangePreferenceHandler,
	"EXPLAIN":       ExplainHandler,
}

func OnClientMessage(conn *websocket.Conn, game *model.Game, msg string, devMode bool) error {
	game.LastInteraction = time.Now() // update game's last interaction for janitor grace

	if len(msg) > MaxPayloadLength {
		return gerrors.ErrPayloadTooLarge
	}

	msg = strings.TrimSpace(msg)
	if msg == "" {
		return gerrors.ErrMessageTooShort
	}

	str := SpacesRegEx.Split(msg, -1)
	prefix := strings.ToUpper(str[0])

	// get client if the client already exists
	// NOTICE: client CAN BE nil AT THIS MOMENT!
	client := game.ClientByConnection(conn)

	// find h
	h, ok := Handlers[prefix]
	if !ok {
		return gerrors.ErrUnknownCommand
	}

	// check if h is dev-only
	if h.DevOnly && !devMode {
		return gerrors.ErrDevOnly
	}

	// check access for h
	if !h.AccessLevel.Allowed(client) {
		return gerrors.ErrNoAccess
	}

	// check arg length
	if h.Bounds != nil {
		if !h.Bounds.Applies(len(str[1:])) {
			return gerrors.ErrArgLength
		}
	}

	// check game state
	if !game.State().In(h.StateLevel) {
		return gerrors.ErrGameStateAccess
	}

	var err error

	// execute h
	switch hdl := h.Handler.(type) {
	case handler.BasicHandler:
		err = hdl(conn, game, client)
	case handler.MessagedHandler:
		err = hdl(conn, game, client, str[1:])
	case handler.PrefixedHandler:
		err = hdl(conn, game, client, prefix)
	case handler.PrefixedMessagedHandler:
		err = hdl(conn, game, client, prefix, str[1:])
	case handler.HandlersHandler:
		err = hdl(conn, game, client, str[1:], Handlers)
	default:
		log.Warnf("cannot find h for %s", prefix)
		err = gerrors.ErrInvalidHandler
	}

	if err != nil {
		// send error
		err = model.NewResponseWithError(prefix, err).Respond(conn)
	}

	return err
}
