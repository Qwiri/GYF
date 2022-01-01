package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

var VoteCastHandler = &Handler{
	AccessLevel: AccessJoined,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  model.StateCastVotes,
	Handler: MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
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
		waiting := game.WaitingForVote(topic)
		game.Broadcast(model.PVote(client, waiting.Names()))

		return game.CheckCycle(true, false)
	}),
}
