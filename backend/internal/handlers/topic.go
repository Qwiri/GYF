package handlers

import (
	"encoding/json"
	"strings"

	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var TopicListHandler = &handler.Handler{
	Description: "Returns a list with all topics from the current game",
	Syntax:      "",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateAny,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		return model.PTopicList(game).Respond(conn)
	}),
}

var TopicAddHandler = &handler.Handler{
	Description: "Adds a new topic to the game",
	Syntax:      "(topic!)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateLobby,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if len(game.Topics) >= game.Preferences.MaxTopics {
			return gerrors.ErrTooManyTopics
		}
		// send updated list to all leaders
		defer game.BroadcastTopicListToLeaders()

		topic := strings.TrimSpace(strings.Join(message, " "))
		if game.Topics.Exists(topic) {
			return gerrors.ErrTopicAlreadyExists
		}
		game.Topics.Add(topic)
		return model.PTopicAdd(topic).Respond(conn)
	}),
}

var TopicAddAllHandler = &handler.Handler{
	Description: "Adds all topics from a JSON array",
	Syntax:      "(...topics: Array<string>!)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateLobby,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) (err error) {
		data := strings.Join(message, " ")
		var topics []string
		if err = json.Unmarshal([]byte(data), &topics); err != nil {
			return
		}

		// send updated list to all leaders
		defer game.BroadcastTopicListToLeaders()

		// add all topics to list
		for _, t := range topics {
			t = strings.TrimSpace(t)
			if t == "" {
				continue
			}
			if game.Topics.Exists(t) {
				continue
			}
			// check if there are too many topics
			if len(game.Topics) >= game.Preferences.MaxTopics {
				return gerrors.ErrTooManyTopics
			}
			game.Topics.Add(t)
		}
		return nil
	}),
}

var TopicRemoveHandler = &handler.Handler{
	Description: "Removes a topic from the game",
	Syntax:      "(topic!)",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateLobby,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		topic := strings.TrimSpace(strings.Join(message, " "))
		if !game.Topics.Exists(topic) {
			return gerrors.ErrTopicNotFound
		}
		// send updated list to all leaders
		defer game.BroadcastTopicListToLeaders()

		game.Topics.Delete(topic)
		return model.PTopicRemove(topic).Respond(conn)
	}),
}

var TopicClearHandler = &handler.Handler{
	Description: "Removes all topics from the current game",
	Syntax:      "",
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundExact(0)),
	StateLevel:  util.StateLobby,
	Handler: handler.BasicHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client) error {
		// send updated list to all leaders
		defer game.BroadcastTopicListToLeaders()

		game.Topics.Clear()
		return nil
	}),
}
