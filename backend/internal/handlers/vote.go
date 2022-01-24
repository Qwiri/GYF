package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var VoteCastHandler = &handler.Handler{
	Description: "Votes for a submission",
	Syntax:      "(url!)",
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateCastVotes,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		if game.CurrentTopic == nil {
			return gerrors.ErrTopicNotFound
		}
		topic := game.CurrentTopic
		u := message[0]

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

		// check if already voted and remove if found
		for _, sub := range topic.Submissions {
			sub.Voters.Delete(client)
		}

		// add client as voter
		submission.Voters = append(submission.Voters, client)

		// send response with waiting-for players
		return game.CheckCycle(true, false)
	}),
}
