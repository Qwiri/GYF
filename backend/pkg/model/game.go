package model

import (
	"fmt"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
	"time"
)

type Game struct {
	ID              string
	Clients         map[string]*Client
	Topics          []*Topic
	Started         bool
	CurrentTopic    *Topic
	LastInteraction time.Time
}

func NewGame(id string) *Game {
	return &Game{
		ID:              id,
		Clients:         make(map[string]*Client),
		Topics:          nil,
		Started:         false,
		CurrentTopic:    nil,
		LastInteraction: time.Now(),
	}
}

func (g *Game) Broadcast(msg string) {
	for _, client := range g.Clients {
		util.Write(client.Connection, msg)
	}
}

func (g *Game) BroadcastExcept(conn *websocket.Conn, msg string) {
	for _, client := range g.Clients {
		if client.Connection != conn {
			util.Write(client.Connection, msg)
		}
	}
}

func (g *Game) GetClient(conn *websocket.Conn) *Client {
	for _, c := range g.Clients {
		if c.Connection == conn {
			return c
		}
	}
	return nil
}

func (g *Game) LeaveClient(client *Client, reason string) {
	for k, v := range g.Clients {
		if v == client {
			delete(g.Clients, k)
			// announce client leave
			if client.Name != "" {
				g.Broadcast(fmt.Sprintf("PLAYER_LEAVE %s %s", client.Name, reason))
			}
		}
	}
}
