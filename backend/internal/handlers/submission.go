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
		urlHash, err := util.URLHash(url)
		if err != nil {
			return err
		}

		// check if GIF was already submitted
		for _, s := range topic.Submissions {
			subHash, err := util.URLHash(s.URL)
			if err != nil {
				return err
			}
			if subHash == urlHash {
				return gerrors.ErrGIFTaken
			}
		}

		// check if GIF provider is allowed
		allowed, err := util.IsAllowed(url)
		if err != nil {
			return err
		}
		if !allowed {
			return gerrors.ErrGIFNotAllowed
		}

		// save submission
		topic.Submissions[client.Name] = model.NewSubmission(client, url)

		// return a list with players we're waiting for
		waiting := append([]interface{}{client.Name}, util.WrapClientArray(game.WaitingForGIFSubmission(topic))...)
		if len(waiting) <= 1 {
			return game.ForceStartVote()
		}

		game.Broadcast(model.NewResponse("SUBMIT_GIF", waiting...))
		return nil
	}),
}
