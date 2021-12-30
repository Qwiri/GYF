package model

import (
	"github.com/gofiber/websocket/v2"
	"time"
)

type Game struct {
	ID              string
	Clients         map[string]*Client
	Topics          map[string]*Topic
	Started         bool
	CurrentTopic    *Topic
	LastInteraction time.Time
}

func NewGame(id string) *Game {
	return &Game{
		ID:              id,
		Clients:         make(map[string]*Client),
		Topics:          make(map[string]*Topic),
		Started:         false,
		CurrentTopic:    nil,
		LastInteraction: time.Now(),
	}
}

func (g *Game) Broadcast(response *Response) {
	for _, client := range g.Clients {
		_ = response.Respond(client.Connection)
	}
}

func (g *Game) BroadcastExcept(conn *websocket.Conn, response *Response) {
	for _, client := range g.Clients {
		if client.Connection != conn {
			_ = response.Respond(client.Connection)
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

func (g *Game) SetLeader(client *Client) {
	// reset current leader
	for _, c := range g.Clients {
		if c.Leader {
			c.Leader = false
			g.Broadcast(NewResponse("CHANGE_ROLE", client.Name, "PLAYER"))
		}
	}
	client.Leader = true
	g.Broadcast(NewResponse("CHANGE_ROLE", client.Name, "LEADER"))
}

func (g *Game) LeaveClient(client *Client, reason string) {
	for k, v := range g.Clients {
		if v == client {
			delete(g.Clients, k)
			// announce client leave
			if client.Name != "" {
				g.Broadcast(NewResponse("PLAYER_LEAVE", client.Name, reason))
			}
		}
	}
	// check if the game is now empty
	if len(g.Clients) > 0 {
		// check if there is a leader left
		for _, v := range g.Clients {
			if v.Leader {
				return
			}
			if v != nil {
				g.SetLeader(v)
				break
			}
		}
	}
}
