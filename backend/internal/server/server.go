package server

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"sync"
	"time"
)

type GYFServer struct {
	games   map[string]*model.Game
	gamesMu sync.RWMutex
	devMode bool
}

func NewServer(devMode bool) *GYFServer {
	return &GYFServer{
		games:   make(map[string]*model.Game),
		devMode: devMode,
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

func (gs *GYFServer) NewGame(idLen int) *model.Game {
	// find a free game ID
	var id string
	for id == "" || gs.GameExists(id) {
		id = util.GenerateRandomString(idLen)
	}
	return model.NewGame(id)
}

//////
