package handlers

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var (
	ErrNotLeader          = errors.New("not leader")
	ErrTooManyTopics      = errors.New("too many topics")
	ErrTopicNotFound      = errors.New("topic not found")
	ErrTopicAlreadyExists = errors.New("topic already exists")
)

const MaxTopics = 30

func handleTopicList(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, _ []string) error {
	if !client.Leader {
		return model.NewResponseWithError("TOPIC_LIST", ErrNotLeader).Respond(conn)
	}
	topics := make([]interface{}, len(game.Topics))
	i := 0
	for _, topic := range game.Topics {
		topics[i] = topic.Description
		i += 1
	}
	return model.NewResponse("TOPIC_LIST", topics...).Respond(conn)
}

func handleTopicAdd(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	if !client.Leader {
		return model.NewResponseWithError("TOPIC_ADD", ErrNotLeader).Respond(conn)
	}
	if len(game.Topics) >= MaxTopics {
		return model.NewResponseWithError("TOPIC_ADD", ErrTooManyTopics).Respond(conn)
	}
	topic := strings.Join(message, " ")
	if _, ok := game.Topics[topic]; ok {
		return model.NewResponseWithError("TOPIC_ADD", ErrTopicAlreadyExists).Respond(conn)
	}
	game.Topics[topic] = model.NewTopic(topic)
	return model.NewResponse("TOPIC_ADD").Respond(conn)
}

func handleTopicRemove(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	if !client.Leader {
		return model.NewResponseWithError("TOPIC_REMOVE", ErrNotLeader).Respond(conn)
	}
	topic := strings.Join(message, " ")
	if _, ok := game.Topics[topic]; !ok {
		return model.NewResponseWithError("TOPIC_REMOVE", ErrTopicNotFound).Respond(conn)
	}
	delete(game.Topics, topic)
	return model.NewResponse("TOPIC_REMOVE").Respond(conn)
}
