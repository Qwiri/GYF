package handler

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/gofiber/websocket/v2"
)

type BasicHandler func(*websocket.Conn, *model.Game, *model.Client) error
type MessagedHandler func(*websocket.Conn, *model.Game, *model.Client, []string) error
type PrefixedHandler func(*websocket.Conn, *model.Game, *model.Client, string) error
type PrefixedMessagedHandler func(*websocket.Conn, *model.Game, *model.Client, string, []string) error
type HandlersHandler func(*websocket.Conn, *model.Game, *model.Client, []string, map[string]*Handler) error
