package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"strings"
	"time"
)

type chatSpam struct {
	Previous string
	Time     time.Time
}

var chats = make(map[*model.Client]*chatSpam)

func JanitorChatChats() {
	for k, v := range chats {
		if time.Since(v.Time) > 5*time.Minute {
			delete(chats, k)
		}
	}
}

var ChatHandler = &handler.Handler{
	Description: "Sends a chat message",
	Syntax:      "(...message!)",
	AccessLevel: handler.AccessJoined,
	Bounds:      util.Bounds(util.BoundMin(1)),
	StateLevel:  util.StateAny,
	Handler: handler.MessagedHandler(func(conn *websocket.Conn, game *model.Game, client *model.Client, message []string) error {
		msg := strings.TrimSpace(strings.Join(message, " "))
		// min char length: 2
		if len(msg) < 2 {
			return gerrors.ErrChatMessageTooShort
		}
		// max char length: 100
		if len(msg) > 100 {
			return gerrors.ErrChatMessageTooLong
		}
		// time check
		if chat, ok := chats[client]; ok {
			if time.Since(chat.Time) < 1*time.Second {
				return gerrors.ErrChatMessageTooFast
			}
			if strings.EqualFold(chat.Previous, msg) {
				return gerrors.ErrChatMessageRepeat
			}
		}
		chats[client] = &chatSpam{msg, time.Now()}
		game.Broadcast(model.PChat(client, msg))
		return nil
	}),
}
