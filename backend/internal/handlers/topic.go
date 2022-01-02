package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var TopicListHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateAny,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.PTopicList(game).Respond(conn)
	}),
}

var TopicAddHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateLobby,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if len(game.Topics) >= game.Preferences.MaxTopics {
			return gerrors.ErrTooManyTopics
		}
		topic := strings.TrimSpace(strings.Join(message, " "))
		if game.Topics.Exists(topic) {
			return gerrors.ErrTopicAlreadyExists
		}
		game.Topics.Add(topic)
		return model.PTopicAdd(topic).Respond(conn)
	}),
}

var TopicRemoveHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateLobby,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		topic := strings.TrimSpace(strings.Join(message, " "))
		if !game.Topics.Exists(topic) {
			return gerrors.ErrTopicNotFound
		}
		game.Topics.Delete(topic)
		return model.PTopicRemove(topic).Respond(conn)
	}),
}
