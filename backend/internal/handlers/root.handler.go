package handlers

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"regexp"
	"strings"
	"time"
)

var (
	ErrMessageTooShort = errors.New("message too short")
	ErrUnknownCommand  = errors.New("unknown command")
	ErrNoAccess        = errors.New("no access to that command")
	ErrInvalidHandler  = errors.New("invalid handler func")
	ErrDevOnly         = errors.New("handler is dev only")
	ErrArgLength       = errors.New("unexpected arg length")
)

var SpacesRegEx = regexp.MustCompile(`\s+`)

type Handler struct {
	AccessLevel Access
	Handler     interface{}
	DevOnly     bool
	Bounds      util.Boundaries
}

var Handlers = map[string]*Handler{
	"WHOAMI":       WhoAmIHandler,
	"JOIN":         JoinHandler,
	"LIST":         ListHandler,
	"CHAT":         ChatHandler,
	"TOPIC_LIST":   TopicListHandler,
	"TOPIC_ADD":    TopicAddHandler,
	"TOPIC_REMOVE": TopicRemoveHandler,
}

type BasicHandler func(*websocket.Conn, *model.Game, *model.Client) error
type MessagedHandler func(*websocket.Conn, *model.Game, *model.Client, []string) error
type PrefixedHandler func(*websocket.Conn, *model.Game, *model.Client, string) error
type PrefixedMessagedHandler func(*websocket.Conn, *model.Game, *model.Client, string, []string) error

func OnClientMessage(conn *websocket.Conn, game *model.Game, msg string, devMode bool) error {
	game.LastInteraction = time.Now() // update game's last interaction for janitor grace

	msg = strings.TrimSpace(msg)
	if len(msg) == 0 {
		return ErrMessageTooShort
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
		return ErrUnknownCommand
	}

	// check if handler is dev-only
	if handler.DevOnly && !devMode {
		return ErrDevOnly
	}

	// check access for handler
	if !handler.AccessLevel.Allowed(client) {
		return ErrNoAccess
	}

	// check arg length
	if handler.Bounds != nil {
		if !handler.Bounds.Applies(len(str[1:])) {
			return ErrArgLength
		}
	}

	// execute handler
	switch hdl := handler.Handler.(type) {
	case BasicHandler:
		return hdl(conn, game, client)
	case MessagedHandler:
		return hdl(conn, game, client, str[1:])
	case PrefixedHandler:
		return hdl(conn, game, client, prefix)
	case PrefixedMessagedHandler:
		return hdl(conn, game, client, prefix, str[1:])
	default:
		log.Warnf("cannot find handler for %s", prefix)
		return ErrInvalidHandler
	}
}
