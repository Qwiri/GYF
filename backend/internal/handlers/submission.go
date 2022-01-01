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
	StateLevel:  model.StateSubmitGIF,
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if game.CurrentTopic == nil {
			return gerrors.ErrTopicNotFound
		}
		topic := game.CurrentTopic
		url := message[0]

		// check if GIF was already submitted
		if _, found, err := topic.Submissions.ByURLLoose(url); found || err != nil {
			if err != nil {
				return err
			}
			return gerrors.ErrGIFTaken
		}

		// check if GIF provider is allowed
		if allowed, err := util.IsAllowed(url); err != nil {
			return err
		} else if !allowed {
			return gerrors.ErrGIFNotAllowed
		}

		// save submission
		topic.Submissions[client.Name] = model.NewSubmission(client, url)

		// return a list with players we're waiting for
		waiting := game.WaitingForGIFSubmission(topic).Names()
		game.Broadcast(model.PSubmitGIF(client, waiting))

		// check game cycle
		return game.CheckCycle(true, false)
	}),
}
