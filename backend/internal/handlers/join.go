package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"math/rand"
	"strings"
)

var JoinHandler = &handler.Handler{
	AccessLevel: handler.AccessGuest,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		username := strings.TrimSpace(message[0])
		// check if username is allowed
		if !util.IsNameValid(username) {
			return gerrors.ErrNameInvalid
		}
		// check if there is already a client with the same name?
		for _, c := range game.Clients {
			if strings.EqualFold(c.Name, username) {
				return gerrors.ErrNameExists
			}
		}
		// client already joined?
		if client != nil {
			if client.Name != "" {
				// client has already a name
				return gerrors.ErrAlreadyJoined
			}
			client.Name = username
		} else {
			client = model.NewClient(conn, username)
			// add client to game map
			game.Clients[client.Name] = client
		}

		log.Infof("[%s] %s joined", game.ID, client.Name)

		// is this the first player? (leader)
		if len(game.Clients) == 1 {
			game.SetLeader(client)
		}

		// broadcast player join and client list
		game.Broadcast(model.PJoin(client, game))
		game.Broadcast(model.PList(game.Clients))
		game.Broadcast(model.PState(game.State()))

		// send game-relevant things
		if util.StateInGame.Contains(game.State()) {
			game.Broadcast(model.PStats(game))
			game.BroadcastWaitingFor()

			// only send more info if we have a topic atm
			if topic := game.CurrentTopic; topic != nil {
				// for show votes, we don't need to update waiting-for
				if util.StateShowVotes.In(game.State()) {
					results := topic.Submissions.AsArray().Results()
					_ = model.PVoteResults(results...).Respond(conn)
				} else {
					// send topic to submit GIF
					if util.StateSubmitGIF.In(game.State()) {
						_ = model.PNextRound(topic.Description, game.Topics.PlayedTopicsCount(), len(game.Topics)).Respond(conn)
					}
					if util.StateCastVotes.In(game.State()) {
						submissions := topic.Submissions.AllExceptFrom(client)
						urls := submissions.URLs()
						// randomize URLs
						rand.Shuffle(len(urls), func(i, j int) {
							urls[i], urls[j] = urls[j], urls[i]
						})
						_ = model.PVoteStart(urls...).Respond(conn)
					}
				}
			}
		}

		// send preferences to player
		return model.PPreferences(game.Preferences).Respond(conn)
	}),
}
