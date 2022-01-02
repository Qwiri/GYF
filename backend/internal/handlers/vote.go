package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var VoteCastHandler = &handler.Handler{
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateCastVotes,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if game.CurrentTopic == nil {
			return gerrors.ErrTopicNotFound
		}
		topic := game.CurrentTopic
		u := message[0]

		// check if already voted
		if topic.Submissions.HasVoted(client) {
			return gerrors.ErrAlreadyVoted
		}

		// get submission for URL
		var (
			submission *model.Submission
			found      bool
		)
		if submission, found = topic.Submissions.ByURL(u); !found {
			return gerrors.ErrSubmissionNotFound
		}

		// check if client is submitter
		if submission.Creator == client {
			return gerrors.ErrVoteSelf
		}

		// add client as voter
		submission.Voters = append(submission.Voters, client)

		// send response with waiting-for players
		return game.CheckCycle(true, false)
	}),
}
