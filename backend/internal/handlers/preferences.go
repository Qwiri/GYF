package handlers

import (
	"encoding/json"
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
)

var ChangePreferenceHandler = &handler.Handler{
	AccessLevel: handler.AccessLeader,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, args []string) (err error) {
		data := []byte(strings.Join(args, " "))
		type changePreferencesPayload struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		}
		var payload changePreferencesPayload
		if err = json.Unmarshal(data, &payload); err != nil {
			return
		}
		switch payload.Key {
		case "AutoSkip":
			i, ok := payload.Value.(bool)
			if !ok {
				return gerrors.ErrTypeInvalid
			}
			game.Preferences.AutoSkip = i
		case "ShuffleTopics":
			i, ok := payload.Value.(bool)
			if !ok {
				return gerrors.ErrTypeInvalid
			}
			game.Preferences.ShuffleTopics = i
		default:
			return gerrors.ErrUnknownPreference
		}

		game.Broadcast(model.PPreferences(game.Preferences))
		return model.NewResponse("CHANGE_PREF").Respond(conn)
	}),
}
