package model

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/gofiber/websocket/v2"
	"time"
)

type Game struct {
	ID              string
	Clients         map[string]*Client
	Topics          Topics
	Started         bool
	CurrentTopic    *Topic
	LastInteraction time.Time
}

func NewGame(id string) (game *Game) {
	game = &Game{
		ID:              id,
		Clients:         make(map[string]*Client),
		Started:         false,
		CurrentTopic:    nil,
		LastInteraction: time.Now(),
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

func (g *Game) EndGame() {
	// reset topic plays
	for _, t := range g.Topics {
		t.Played = false
		t.Submissions = make(map[string]*Submission)
	}
	g.Started = false
	g.Broadcast(NewResponse("GAME_END", "TOPIC_END"))
}

func (g *Game) NextRound() (err error) {
	// get next topic
	var topic *Topic
	if topic, err = g.Topics.Next(); err != nil {
		if err == gerrors.ErrNoTopicsLeft {
			// TODO: End Game
		} else {
		}
		return
	}
	topic.Played = true
	topic.CanSubmit = true
	g.CurrentTopic = topic
	g.Broadcast(NewResponse("NEXT_ROUND", topic.Description, g.Topics.PlayedCount(), len(g.Topics)))
	return nil
}
