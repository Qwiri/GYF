package handlers

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	websocket2 "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOnClientMessage(t *testing.T) {
	conn := &websocket.Conn{Conn: &websocket2.Conn{}}
	game := model.NewGame("")
	var err error

	// check payload too large
	err = OnClientMessage(conn, game, strings.Repeat("a", 4096), false)
	assert.NotEqual(t, err, gerrors.ErrPayloadTooLarge)
	err = OnClientMessage(conn, game, strings.Repeat("a", 4097), false)
	assert.Equal(t, err, gerrors.ErrPayloadTooLarge)

	// check message too short
	err = OnClientMessage(conn, game, "a", false)
	assert.NotEqual(t, err, gerrors.ErrMessageTooShort)
	err = OnClientMessage(conn, game, "", false)
	assert.Equal(t, err, gerrors.ErrMessageTooShort)

	// check unknown handler
	err = OnClientMessage(conn, game, "JOIN", false)
	assert.NotEqual(t, err, gerrors.ErrUnknownCommand)
	err = OnClientMessage(conn, game, "HAWDHAWDHAWDH", false)
	assert.Equal(t, err, gerrors.ErrUnknownCommand)

	// check devonly
	err = OnClientMessage(conn, game, "SKIP", true)
	assert.NotEqual(t, err, gerrors.ErrDevOnly)
	err = OnClientMessage(conn, game, "SKIP", false)
	assert.Equal(t, err, gerrors.ErrDevOnly)
}
