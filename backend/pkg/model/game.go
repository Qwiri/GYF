package model

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"time"
)

type GameState int

const (
	StateLobby GameState = 1 << iota
	StateSubmitGIF
	StateCastVotes
	StateShowVotes
)

const (
	StateAny    = StateLobby | StateSubmitGIF | StateCastVotes | StateShowVotes
	StateInGame = StateSubmitGIF | StateCastVotes | StateShowVotes
)

func (g GameState) Allowed(game *Game) bool {
	return game.State&g == game.State
}

type Game struct {
	ID              string
	Clients         map[string]*Client
	Topics          Topics
	CurrentTopic    *Topic
	State           GameState
	LastInteraction time.Time
}

func NewGame(id string) (game *Game) {
	game = &Game{
		ID:              id,
		Clients:         make(map[string]*Client),
		CurrentTopic:    nil,
		State:           StateLobby,
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

func (g *Game) ForceEndGame(reason string) (err error) {
	// reset topic plays
	for _, t := range g.Topics {
		t.Played = false
		t.Submissions = make(map[string]*Submission)
	}
	g.State = StateLobby
	g.Broadcast(NewResponse("GAME_END", reason))
	return
}

func (g *Game) ForceNextRound() (err error) {
	// get next topic
	var topic *Topic
	if topic, err = g.Topics.Next(); err != nil {
		if err == gerrors.ErrNoTopicsLeft {
			err = g.ForceEndGame("NO_TOPIC_LEFT")
		} else {
			log.WithError(err).Warn("cannot get next round")
		}
		return
	}
	topic.Played = true

	g.State = StateSubmitGIF
	g.CurrentTopic = topic
	g.Broadcast(NewResponse("NEXT_ROUND", topic.Description, g.Topics.PlayedCount(), len(g.Topics)))
	return nil
}

func (g *Game) ForceStartVote() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}

	// do not allow more votes
	g.State = StateCastVotes

	for _, client := range g.Clients {
		// build list of URLs for client
		var urls []interface{}
		for _, sub := range g.CurrentTopic.Submissions {
			// skip created submission
			if sub.Creator == client {
				continue
			}
			urls = append(urls, sub.URL)
		}
		if err = NewResponse("VOTE_START", urls...).Respond(client.Connection); err != nil {
			log.Warnf("cannot send vote to %s: %v", client.Name, err)
		}
	}

	return nil
}

func (g *Game) ForceShowVoteResults() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}

	g.State = StateShowVotes

	type voteResult struct {
		URL     string   `json:"url"`
		Creator string   `json:"creator"`
		Voters  []string `json:"voters"`
	}
	var results []interface{}
	for _, sub := range g.CurrentTopic.Submissions {
		var voters = make([]string, len(sub.Voters))
		for i, voter := range sub.Voters {
			voters[i] = voter.Name
		}
		results = append(results, voteResult{
			URL:     sub.URL,
			Creator: sub.Creator.Name,
			Voters:  voters,
		})
	}
	g.Broadcast(NewResponse("VOTE_RESULTS", results...))
	return
}

func (g *Game) GetStats(user string) (res int) {
	for _, topic := range g.Topics {
		for _, sub := range topic.Submissions {
			if sub.Creator.Name == user {
				res += len(sub.Voters)
			}
		}
	}
	return
}
