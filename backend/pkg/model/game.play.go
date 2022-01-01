package model

import (
	"github.com/Qwiri/GYF/backend/pkg/gerrors"
	"github.com/apex/log"
)

// SetLeader sets the client as the sole leader and sends a
// CHANGE_ROLE message to all players with the old and the new leader.
func (g *Game) SetLeader(client *Client) {
	// remove old leader(s)
	for _, c := range g.Clients {
		if c.Leader {
			c.Leader = false
			g.Broadcast(PChangeRole(c, "PLAYER"))
			log.Infof("[%s] %s is no longer a leader", g.ID, client.Name)
		}
	}
	client.Leader = true
	g.Broadcast(PChangeRole(client, "LEADER"))
	log.Infof("[%s] %s is now a leader", g.ID, client.Name)
}

// LeaveClient removes a client from the game and sends a PLAYER_LEAVE message to all other players
func (g *Game) LeaveClient(client *Client, reason string) {
	// remove client from game
	g.Clients.Delete(client)
	g.Broadcast(PPlayerLeave(client, reason))

	// if game is now empty, reset game
	if g.IsEmpty() {
		g.Reset(true)
		return
	}

	// check if client was a leader
	if client.Leader {
		if leader := g.CreateLeader(); leader != nil {
			log.Debugf("[%s] New leader for game: %s", g.ID, leader.Name)
		} else {
			log.Warnf("[%s] Cannot find new leader", g.ID)
		}
	}

	// if we aren't ingame, we don't need to check the game cycle
	if !StateInGame.Allowed(g) {
		return
	}

	// check game cycle
	_ = g.CheckCycle(true, false)
}

// CheckCycle checks if we're waiting for clients and if not, continue the game
func (g *Game) CheckCycle(checkAutoSkip, force bool) (err error) {
	if checkAutoSkip && !g.Preferences.AutoSkip {
		return
	}

	switch g.GetState() {
	case StateLobby:
		// no need to do anything in the lobby state
	case StateShowVotes:
		// no need to do anything in the show vote state
		// since the leader should skip to the next round
		if force {
			err = g.ForceNextRound()
		}

	case StateSubmitGIF:
		// check if we're currently waiting for another submission
		if g.CurrentTopic == nil {
			log.Warnf("[%s] Tried to check submission count, but current topic was nil", g.ID)
			return
		}
		if force || len(g.WaitingForGIFSubmission(g.CurrentTopic)) == 0 {
			err = g.ForceStartVote()
		}

	case StateCastVotes:
		// check if we're currently waiting for another vote
		if g.CurrentTopic == nil {
			log.Warnf("[%s] Tried to check voters count, but current topic was nil", g.ID)
			return
		}
		if force || len(g.WaitingForVote(g.CurrentTopic)) == 0 {
			err = g.ForceShowVoteResults()
		}

	}

	return
}

func (g *Game) ForceEndGame(reason string) (err error) {
	g.Broadcast(PStats(g))        // send stats for winning screen
	g.Broadcast(PGameEnd(reason)) // send game end
	g.Reset(false)
	return
}

func (g *Game) ForceNextRound() (err error) {
	// get next topic
	var topic *Topic
	if topic, err = g.Topics.NextTopic(); err != nil {
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
	g.SetState(StateSubmitGIF)
	g.CurrentTopic = topic

	g.Broadcast(PNextRound(topic.Description, g.Topics.PlayedTopicsCount(), len(g.Topics)))
	return nil
}

func (g *Game) ForceStartVote() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}
	topic := g.CurrentTopic

	// do not allow more GIF submissions
	g.SetState(StateCastVotes)

	for _, client := range g.Clients {
		submissions := topic.Submissions.AllExceptFrom(client)
		urls := submissions.URLs()
		if err = PVoteStart(urls...).RespondTo(client); err != nil {
			log.Warnf("cannot send vote to %s: %v", client.Name, err)
		}
	}

	return nil
}

func (g *Game) ForceShowVoteResults() (err error) {
	if g.CurrentTopic == nil {
		return gerrors.ErrTopicNotFound
	}
	topic := g.CurrentTopic

	// do not allow more votes
	g.SetState(StateShowVotes)

	// broadcast results
	results := topic.Submissions.AsArray().Results()
	g.Broadcast(PVoteResults(results...))

	return
}
