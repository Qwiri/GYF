package model

import (
	"fmt"
	"github.com/Qwiri/GYF/backend/pkg/config"
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"math/rand"
	"strings"
	"time"
)

type Game struct {
	ID              string
	Clients         ClientMap
	Topics          TopicArray
	CurrentTopic    *Topic
	state           util.GameState
	LastInteraction time.Time
	Preferences     *GamePreferences
}

/// Constructors

// NewGame returns a new game object with default values
func NewGame(id string) (game *Game) {
	pref := DefaultPreferences
	game = &Game{
		ID:              id,
		Clients:         make(ClientMap),
		CurrentTopic:    nil,
		state:           util.StateLobby,
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

/// Getters & Setters

// SetState sets the current GameState and broadcasts the change to all clients
func (g *Game) SetState(state util.GameState) {
	g.state = state
	g.Broadcast(PState(state))
}

// State returns the current GameState
func (g *Game) State() util.GameState {
	return g.state
}

// IsEmpty returns if there are clients in the current game
func (g *Game) IsEmpty() bool {
	return len(g.Clients) <= 0
}

/// Util

// Reset resets the Topic.Played values, Game.CurrentTopic, Game.State()
// and purges (if purgeClients was set to true) all clients
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
	g.SetState(util.StateLobby)
	g.LastInteraction = time.Now()

	g.BroadcastWaitingForNone()
}

// Broadcast sends a response to all clients
func (g *Game) Broadcast(response *Response) {
	for _, client := range g.Clients {
		_ = response.Respond(client.Connection)
	}
}

// BroadcastExcept sends a response to all clients except to a specific connection
func (g *Game) BroadcastExcept(conn *websocket.Conn, response *Response) {
	for _, client := range g.Clients {
		if client.Connection != conn {
			_ = response.Respond(client.Connection)
		}
	}
}

func (g *Game) BroadcastTopicListToLeaders() {
	p := PTopicList(g)
	for _, c := range g.Clients {
		if c.Leader || g.Preferences.Permissions&PermissionListTopics == PermissionListTopics {
			_ = p.RespondTo(c)
		}
	}
}

func (g *Game) BroadcastWaitingForAll() {
	g.Broadcast(PWaitingFor(g.Clients.Array()))
}
func (g *Game) BroadcastWaitingForNone() {
	g.Broadcast(PWaitingFor(ClientArray{}))
}
func (g *Game) BroadcastWaitingFor() {
	if g.CurrentTopic == nil {
		g.BroadcastWaitingForNone()
		return
	}
	var (
		topic   = g.CurrentTopic
		waiting ClientArray
	)
	if util.StateSubmitGIF.In(g.State()) {
		waiting = g.WaitingForGIFSubmission(topic)
	} else if util.StateCastVotes.In(g.State()) {
		waiting = g.WaitingForVote(topic)
	}
	g.Broadcast(PWaitingFor(waiting))
}

/// Play

// SetLeader sets the client as the sole leader and sends a
// CHANGE_ROLE message to all players with the old and the new leader.
func (g *Game) SetLeader(client *Client) {
	// remove old leader(s)
	for _, c := range g.Clients {
		if c != client && c.Leader {
			c.Leader = false
			g.Broadcast(PChangeRole(c, "PLAYER"))
			log.Infof("[%s] %s is no longer a leader", g.ID, client.Name)
		}
	}
	client.Leader = true
	g.Broadcast(PChangeRole(client, "LEADER"))

	// send topic list to leader
	g.BroadcastTopicListToLeaders()

	log.Infof("[%s] %s is now a leader", g.ID, client.Name)
}

// LeaveClient removes a client from the game and sends a PLAYER_LEAVE message to all other players
func (g *Game) LeaveClient(client *Client, reason string) {
	log.Infof("[%s] %s left (leader: %v)", g.ID, client.Name, client.Leader)

	// remove client from game
	g.Clients.Delete(client)

	// if game is now empty, reset game
	if g.IsEmpty() {
		g.Reset(true)
		return
	}

	// broadcast player list and leave
	g.Broadcast(PPlayerLeave(client, reason))
	g.Broadcast(PList(g.Clients))

	// check if client was a leader
	if client.Leader {
		log.Debugf("[%s] %s was a leader. Choosing a new one.", g.ID, client.Name)
		if leader := g.FindNonLeaderClient(); leader != nil {
			log.Debugf("[%s] New leader for game: %s", g.ID, leader.Name)
			g.SetLeader(leader)
		} else {
			log.Warnf("[%s] Cannot find new leader", g.ID)
		}
	}

	// if we aren't in-game, we don't need to check the game cycle
	if !g.State().In(util.StateInGame) {
		return
	}

	g.Broadcast(PStats(g))

	// check game cycle
	_ = g.CheckCycle(true, false)
}

// CheckCycle checks if we're waiting for clients and if not, continue the game
func (g *Game) CheckCycle(checkAutoSkip, force bool) (err error) {
	g.BroadcastWaitingFor()

	if checkAutoSkip && !g.Preferences.AutoSkip {
		return
	}

	var change bool

	switch g.State() {
	case util.StateLobby:
		// no need to do anything in the lobby state
	case util.StateShowVotes:
		// no need to do anything in the show vote state
		// since the leader should skip to the next round
		if force {
			err = g.ForceNextRound()
			change = true
		}

	case util.StateSubmitGIF:
		// check if we're currently waiting for another submission
		if g.CurrentTopic == nil {
			log.Warnf("[%s] Tried to check submission count, but current topic was nil", g.ID)
			return
		}
		if force || len(g.WaitingForGIFSubmission(g.CurrentTopic)) == 0 {
			err = g.ForceStartVote()
			change = true
		}

	case util.StateCastVotes:
		// check if we're currently waiting for another vote
		if g.CurrentTopic == nil {
			log.Warnf("[%s] Tried to check voters count, but current topic was nil", g.ID)
			return
		}
		if force || len(g.WaitingForVote(g.CurrentTopic)) == 0 {
			err = g.ForceShowVoteResults()
			change = true
		}
	}

	if change {
		// recheck cycle
		err = g.CheckCycle(checkAutoSkip, false)
	}

	return
}

// ForceEndGame ends the game
func (g *Game) ForceEndGame(reason string) (err error) {
	stats := make(map[string]int)
	for _, c := range g.Clients {
		stats[c.Name] = g.StatsForUser(c.Name)
	}
	config.Obj.Pushover.Send(
		fmt.Sprintf("STATS { %+v }", stats),
		fmt.Sprintf("Game (End) [%s]", g.ID))

	g.Broadcast(PStats(g))        // send stats for winning screen
	g.Broadcast(PGameEnd(reason)) // send game end
	g.Reset(false)
	return
}

// ForceNextRound starts the next round if a topic is available
func (g *Game) ForceNextRound() (err error) {

	// game freshly started?
	if g.Topics.PlayedTopicsCount() == 0 {
		config.Obj.Pushover.Send(
			fmt.Sprintf("CLIENTS { %+v }\nTOPICS [ %+v ]",
				strings.Join(g.Clients.Array().Names(), ", "),
				strings.Join(g.Topics.Strings(), ", ")),
			fmt.Sprintf("Game (Start) [%s]", g.ID))
	}

	// get next topic
	var topic *Topic
	if g.Preferences.ShuffleTopics {
		topic, err = g.Topics.RandomTopic()
	} else {
		topic, err = g.Topics.NextTopic()
	}
	if err != nil {
		if err == gerrors.ErrNoTopicsLeft {
			err = g.ForceEndGame("NO_TOPIC_LEFT")
		} else {
			log.WithError(err).Warnf("[%s] cannot get next round", g.ID)
		}
		return
	}

	// set topic.Played to true to prevent the topic from re-appearing
	topic.Played = true

	// update game state and meta
	g.SetState(util.StateSubmitGIF)
	g.CurrentTopic = topic

	g.Broadcast(PNextRound(topic.Description, g.Topics.PlayedTopicsCount(), len(g.Topics)))
	g.Broadcast(PStats(g))
	g.BroadcastWaitingForAll()
	return nil
}

// ForceStartVote starts the vote process of the previous submissions
func (g *Game) ForceStartVote() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}
	topic := g.CurrentTopic

	// do not allow more GIF submissions
	g.SetState(util.StateCastVotes)

	for _, client := range g.Clients {
		submissions := topic.Submissions.AllExceptFrom(client)
		urls := submissions.URLs()

		// randomize URLs
		rand.Shuffle(len(urls), func(i, j int) {
			urls[i], urls[j] = urls[j], urls[i]
		})

		if err = PVoteStart(urls...).RespondTo(client); err != nil {
			log.Warnf("cannot send vote to %s: %v", client.Name, err)
		}
	}

	g.BroadcastWaitingForAll()
	return nil
}

// ForceShowVoteResults sends the voting results to all clients
func (g *Game) ForceShowVoteResults() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}
	topic := g.CurrentTopic

	var b util.Bob
	b.Writefl("'%s' [%d/%d]", topic.Description, g.Topics.PlayedTopicsCount(), len(g.Topics))
	for _, sub := range topic.Submissions {
		b.Alll(sub.Creator.Name, ": ", sub.URL, " (", len(sub.Voters), " votes)")
	}
	config.Obj.Pushover.Send(
		b.String(),
		fmt.Sprintf("Game (VRes) [%s]", g.ID))

	// do not allow more votes
	g.SetState(util.StateShowVotes)

	// broadcast results
	results := topic.Submissions.AsArray().Results()
	g.Broadcast(PVoteResults(results...))

	// broadcast stats
	g.Broadcast(PStats(g))

	// clear waiting for
	g.BroadcastWaitingForNone()
	return
}

/// Advanced Getters

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
		if len(topic.Submissions.AllExceptFrom(c)) == 0 {
			continue
		}
		if !voters.Contains(c) {
			res = append(res, c)
		}
	}
	return
}

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

///

func (g *Game) FindNonLeaderClient() *Client {
	for _, c := range g.Clients {
		if !c.Leader {
			return c
		}
	}
	return nil
}
