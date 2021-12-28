package main

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strings"
)

func (gs *GYFServer) CreateRoutes(app *fiber.App) {
	app.Get("/game/list", gs.RouteListGames)
	app.Get("/game/create", gs.RouteCreateGame)

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
			if err := c.WriteMessage(websocket.TextMessage, []byte("ERROR game not found")); err != nil {
				log.WithError(err).Warn("[ws] cannot write error message")
			}
			closeConnection(c)
			return
		}

		// check if game is in progress
		if game.Started {
			log.Warnf("client tried to connect to game %s but the game was running", gameID)
			if err := c.WriteMessage(websocket.TextMessage, []byte("ERROR game already started")); err != nil {
				log.WithError(err).Warn("[ws] cannot write error message")
			}
			closeConnection(c)
			return
		}

		log.Infof("[ws] websocket connection with game %+v", gameID)

		for {
			if _, msg, err := c.ReadMessage(); err != nil {
				log.WithError(err).Warn("[ws] cannot read message from websocket")
				break
			} else {
				onClientMessage(c, game, strings.TrimSpace(string(msg)))
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
