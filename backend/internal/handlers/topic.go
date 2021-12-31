package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
)

const MaxTopics = 30

var TopicListHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	Handler: BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		topics := make([]interface{}, len(game.Topics))
		for i, topic := range game.Topics {
			topics[i] = topic.Description
		}
		return model.NewResponse("TOPIC_LIST", topics...).Respond(conn)
	}),
}

var TopicAddHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		// check if game already started
		if game.Started {
			return gerrors.ErrGameStarted
		}
		if len(game.Topics) >= MaxTopics {
			return gerrors.ErrTooManyTopics
		}
		topic := strings.TrimSpace(strings.Join(message, " "))
		if game.Topics.Exists(topic) {
			return gerrors.ErrTopicAlreadyExists
		}
		game.Topics.Add(topic)
		return model.NewResponse("TOPIC_ADD").Respond(conn)
	}),
}

var TopicRemoveHandler = &Handler{
	AccessLevel: AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		// check if game already started
		if game.Started {
			return gerrors.ErrGameStarted
		}
		topic := strings.TrimSpace(strings.Join(message, " "))
		if !game.Topics.Exists(topic) {
			return gerrors.ErrTopicNotFound
		}
		game.Topics.Delete(topic)
		return model.NewResponse("TOPIC_REMOVE").Respond(conn)
	}),
}
