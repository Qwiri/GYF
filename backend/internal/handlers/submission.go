package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var SubmitGIFHandler = &handler.Handler{
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateSubmitGIF,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if game.CurrentTopic == nil {
			return gerrors.ErrTopicNotFound
		}
		topic := game.CurrentTopic
		url := message[0]

		// check if GIF was already submitted
		if sub, found, err := topic.Submissions.ByURLLoose(url); found || err != nil {
			if err != nil {
				return err
			}
			if sub.Creator == client {
				return gerrors.ErrAlreadySubmitted
			}
			return gerrors.ErrGIFTaken
		}

		// check if GIF provider is allowed
		if allowed, err := util.IsURLAllowed(url); err != nil {
			return err
		} else if !allowed {
			return gerrors.ErrGIFNotAllowed
		}

		// save submission
		sub := model.NewSubmission(client, url)
		topic.Submissions[client.Name] = sub

		// broadcast player submit
		game.Broadcast(model.PSubmitGIF(client))

		return game.CheckCycle(true, false)
	}),
}
