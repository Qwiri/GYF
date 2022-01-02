package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var JoinHandler = &handler.Handler{
	AccessLevel: handler.AccessGuest,
	Bounds:      util.Bounds(util.BoundExact(1)),
	StateLevel:  util.StateLobby,
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
		// broadcast player join
		game.Broadcast(model.PJoin(client, game))

		// send preferences to player
		return model.PPreferences(game.Preferences).Respond(conn)
	}),
}
