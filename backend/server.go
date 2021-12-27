package main

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strings"
	"sync"
	"time"
)

type GYFServer struct {
	games   map[string]*Game
	gamesMu sync.RWMutex
}

func NewServer() *GYFServer {
	return &GYFServer{
		games: make(map[string]*Game),
	}
}

// JanitorCheck checks for games which had no interaction for the last
// {timeout} duration
func (gs *GYFServer) JanitorCheck(timeout time.Duration) {
	gs.gamesMu.Lock()
	defer gs.gamesMu.Unlock()

	n := time.Now()
	for id, game := range gs.games {
		if n.Sub(game.LastInteraction) > timeout {
			log.WithFields(log.Fields{
				"game": game,
				"sub":  n.Sub(game.LastInteraction),
			}).Info("   [Janitor] cleaned game")
			delete(gs.games, id)
		}
	}
}

//////

func (gs *GYFServer) GameExists(gameID string) (o bool) {
	gs.gamesMu.RLock()
	defer gs.gamesMu.RUnlock()

	_, o = gs.games[gameID]
	return
}

func (gs *GYFServer) NewGame(idLen int) *Game {
	// find a free game ID
	var id string
	for id == "" || gs.GameExists(id) {
		id = GenerateRandomString(idLen)
	}
	return &Game{
		ID:              id,
		LastInteraction: time.Now(),
	}
}

//////

func (gs *GYFServer) CreateRoutes(app *fiber.App) {
	app.Get("/game", gs.RouteListGames)
	app.Get("/game/create", gs.RouteCreateGame)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		gameID := c.Params("id")
		log.Infof("[ws] got connection to id %s", gameID)

		game, ok := gs.games[gameID]
		if !ok || c.Locals("allowed") != true {
			if err := c.Close(); err != nil {
				log.WithError(err).Warn("[ws] cannot close websocket connection")
			}
			return
		}
		log.Infof("[ws] websocket connection with game %+v", gameID)

		// update last interaction for game (janitor)
		game.UpdateInteraction()

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			messageType int
			msg         []byte
			err         error
		)
		for {
			if messageType, msg, err = c.ReadMessage(); err != nil {
				log.WithError(err).Warn("[ws] cannot read message from websocket")
				break
			}
			str := strings.TrimSpace(string(msg))
			log.Infof("[ws] received (%d) %s", messageType, str)

			var bob strings.Builder
			bob.WriteString("Pong! (")
			bob.WriteString(str)
			bob.WriteRune(')')

			if err = c.WriteMessage(messageType, []byte(bob.String())); err != nil {
				log.WithError(err).Warn("[ws] cannot write message to webhook")
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
