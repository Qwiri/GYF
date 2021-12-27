package main

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
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
