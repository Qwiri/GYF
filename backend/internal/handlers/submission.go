package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
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
		urlHash, err := util.URLHash(url)
		log.Infof("URL: %s, Hash: %s", url, urlHash)
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

		// check if GIF is allowed
		allowed, err := util.IsAllowed(url)
		if err != nil {
			return err
		}
		if !allowed {
			return gerrors.ErrGIFNotAllowed
		}

		// save submission
		topic.Submissions[client.Name] = &model.Submission{
			Creator: client,
			URL:     url,
		}

		// return a list with players we're waiting for
		waiting := append([]interface{}{client.Name}, topic.Waiting(game)...)
		if len(waiting) <= 1 {
			// TODO: if all voted, auto continue
			topic.CanSubmit = false
		}

		game.Broadcast(model.NewResponse("SUBMIT_GIF", waiting...))
		return nil
	}),
}
