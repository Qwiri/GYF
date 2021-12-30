package server

import (
	"errors"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"strings"

	"github.com/Qwiri/GYF/backend/internal/handlers"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	ErrGameNotFound = errors.New("game not found")
	ErrGameStarted  = errors.New("game already started")
)

func (gs *GYFServer) CreateRoutes(app *fiber.App) {
	if gs.devMode {
		app.Get("/game/list", gs.RouteListGames)
	}

	app.Get("/game/create", gs.RouteCreateGame)

	// SOCKET
	app.Use("/game/socket", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/game/socket/:id", websocket.New(func(c *websocket.Conn) {
		gameID := c.Params("id")
		log.Infof("[ws] got connection to id %s", gameID)

		// make sure the game exists
		game, ok := gs.games[gameID]
		if !ok || c.Locals("allowed") != true {
			if err := model.NewResponseWithError("JOIN", ErrGameNotFound).Respond(c); err != nil {
				log.WithError(err).Warn("[ws] cannot write error message")
			}
			util.CloseConnection(c)
			return
		}

		// check if game is in progress
		if game.Started {
			log.Warnf("client tried to connect to game %s but the game was running", gameID)
			if err := model.NewResponseWithError("JOIN", ErrGameStarted).Respond(c); err != nil {
				log.WithError(err).Warn("[ws] cannot write error message")
			}
			util.CloseConnection(c)
			return
		}

		log.Infof("[ws] websocket connection with game %+v", gameID)

		for {
			if _, msg, err := c.ReadMessage(); err != nil {
				log.WithError(err).Warn("[ws] cannot read message from websocket")
				break
			} else if err = handlers.OnClientMessage(c, game, strings.TrimSpace(string(msg)), gs.devMode); err != nil {
				// send error to client
				_ = model.NewResponseWithError("ERROR", err).Respond(c)
				log.WithError(err).Warn("handling client message resulted in an error")
			}
		}

		// invalidate connection (remove clients)
		for _, client := range game.Clients {
			if client.Connection == c {
				game.LeaveClient(client, "disconnected")
				log.Warnf("leaving client %s@%s (disconnected)", client.Name, game.ID)
			}
		}
	}))
}

func (gs *GYFServer) RouteCreateGame(ctx *fiber.Ctx) error {
	game := gs.NewGame(8)

	// add game to GYFServer
	gs.gamesMu.Lock()
	gs.games[game.ID] = game
	gs.gamesMu.Unlock()

	// return game object to user
	return ctx.JSON(game)
}

func (gs *GYFServer) RouteListGames(ctx *fiber.Ctx) error {
	return ctx.JSON(gs.games)
}
