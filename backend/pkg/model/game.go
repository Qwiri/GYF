package model

import (
	"time"
)

type Game struct {
	ID              string
	Clients         ClientMap
	Topics          TopicArray
	CurrentTopic    *Topic
	state           GameState
	LastInteraction time.Time
	Preferences     *GamePreferences
}

func (g *Game) SetState(state GameState) {
	g.state = state
	g.Broadcast(PState(state))
}

func (g *Game) GetState() GameState {
	return g.state
}

func NewGame(id string) (game *Game) {
	pref := DefaultPreferences
	game = &Game{
		ID:              id,
		Clients:         make(ClientMap),
		CurrentTopic:    nil,
		state:           StateLobby,
		LastInteraction: time.Now(),
		Preferences:     &pref,
	}
	// TODO: remove dummy topics
	game.Topics = append(game.Topics,
		NewTopic("I'm Driving Home For Christmas"),
		NewTopic("This Christmas gift... Is not what I expected"),
		NewTopic("Excuse my look, I just fed the reindeer"),
		NewTopic("My reaction to a White Christmas"),
		NewTopic("Seeing family on Christmas Eve"),
		NewTopic("Hearing Last Christmas on the Radio"))
	return
}

func (g *Game) Reset(purgeClients bool) {
	if purgeClients {
		g.Clients = make(ClientMap)
	}

	// reset topics
	for _, t := range g.Topics {
		t.Played = false
		t.Submissions = make(SubmissionMap)
	}
	g.CurrentTopic = nil

	// reset meta
	g.SetState(StateLobby)
	g.LastInteraction = time.Now()
}

func (g *Game) IsEmpty() bool {
	return len(g.Clients) <= 0
}
