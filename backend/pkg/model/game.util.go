package model

import "github.com/gofiber/websocket/v2"

///

type (
	ClientMap   map[string]*Client
	ClientArray []*Client
)

func (M ClientMap) Delete(client *Client) {
	for k, v := range M {
		if v == client {
			delete(M, k)
		}
	}
}

func (A ClientArray) Contains(client *Client) bool {
	for _, c := range A {
		if c == client {
			return true
		}
	}
	return false
}

func (A ClientArray) Names() (res []string) {
	res = make([]string, len(A))
	for i, v := range A {
		res[i] = v.Name
	}
	return
}

// WaitingForGIFSubmission returns all players we're currently waiting for
func (g *Game) WaitingForGIFSubmission(topic *Topic) (res ClientArray) {
	for _, c := range g.Clients {
		if _, ok := topic.Submissions[c.Name]; !ok {
			res = append(res, c)
		}
	}
	return
}

func (g *Game) WaitingForVote(topic *Topic) (res ClientArray) {
	var voters ClientArray
	for _, sub := range topic.Submissions {
		if len(sub.Voters) > 0 {
			voters = append(voters, sub.Voters...)
		}
	}
	for _, c := range g.Clients {
		// if the client hasn't submitted any GIF, we don't have to wait for that client
		if !topic.Submissions.HasSubmittedGIF(c) {
			continue
		}
		if !voters.Contains(c) {
			res = append(res, c)
		}
	}
	return
}

///

func (g *Game) ClientByConnection(conn *websocket.Conn) *Client {
	for _, c := range g.Clients {
		if c.Connection == conn {
			return c
		}
	}
	return nil
}

func (g *Game) StatsForUser(user string) (res int) {
	for _, topic := range g.Topics {
		for _, sub := range topic.Submissions {
			if sub.Creator.Name == user {
				res += len(sub.Voters)
			}
		}
	}
	return
}

func (g *Game) CreateLeader() *Client {
	for _, c := range g.Clients {
		if !c.Leader {
			c.Leader = true
		}
		return c
	}
	return nil
}
