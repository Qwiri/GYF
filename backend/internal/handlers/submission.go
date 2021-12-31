package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var SubmitGIFHandler = &Handler{
	AccessLevel: AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(1)),
	GameStarted: util.Bool(true),
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		// TODO: set CurrentTopic.CanSubmit to false on vote
		if game.CurrentTopic == nil {
			return gerrors.ErrTopicNotFound
		}
		topic := game.CurrentTopic

		// check if we can submit new gifs
		if !topic.CanSubmit {
			return gerrors.ErrCannotSubmit
		}

		url := message[0]

		// TODO: validate URL
		topic.Submissions[client.Name] = &model.Submission{
			Creator: client,
			URL:     url,
		}

		// return a list with players we're waiting for
		var waiting = make([]interface{}, 1)
		waiting[0] = client.Name // first player is player that voted
		for _, c := range game.Clients {
			if _, ok := topic.Submissions[c.Name]; !ok {
				waiting = append(waiting, c.Name)
			}
		}

		if len(waiting) <= 1 {
			// TODO: if all voted, auto continue
			topic.CanSubmit = false
		}

		game.Broadcast(model.NewResponse("SUBMIT_GIF", waiting...))
		return nil
	}),
}
