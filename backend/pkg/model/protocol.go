package model

import "github.com/Qwiri/GYF/backend/pkg/util"

func PGameEnd(reason string) *Response {
	return NewResponse("GAME_END", reason)
}

func PNextRound(description string, currentTopic, totalTopics int) *Response {
	return NewResponse("NEXT_ROUND", description, currentTopic, totalTopics)
}

func PVoteStart(urls ...string) *Response {
	return NewResponse("VOTE_START", util.WrapStringArray(urls...)...)
}

func PState(state util.GameState) *Response {
	return NewResponse("STATE", state)
}

func PVoteResults(results ...*SubmissionResult) *Response {
	return NewResponse("VOTE_RESULTS", WrapVoteResults(results...)...)
}

func PJoin(client *Client, game *Game) *Response {
	return NewResponse("JOIN",
		client.Name,
		len(game.Clients),
		game.Preferences.MinPlayers,
		game.Preferences.MaxPlayers,
	)
}

func PStats(game *Game) *Response {
	stats := make(map[string]int)
	for _, c := range game.Clients {
		stats[c.Name] = game.StatsForUser(c.Name)
	}
	return NewResponse("STATS", stats)
}

func PChat(client *Client, message string) *Response {
	return NewResponse("CHAT", client.Name, message)
}

func PList(clients ClientMap) *Response {
	type listObj struct {
		Name   string `json:"name"`
		Leader bool   `json:"leader"`
	}
	// collect client names
	clientArray := make([]interface{}, len(clients))
	var i = 0
	for _, c := range clients {
		clientArray[i] = listObj{c.Name, c.Leader}
		i += 1
	}
	return NewResponse("LIST", clientArray...)
}

func PSubmitGIF(client *Client, waiting []string) *Response {
	args := append([]string{client.Name}, waiting...)
	return NewResponse("SUBMIT_GIF", util.WrapStringArray(args...)...)
}

func PTopicList(game *Game) *Response {
	topics := make([]interface{}, len(game.Topics))
	for i, topic := range game.Topics {
		topics[i] = topic.Description
	}
	return NewResponse("TOPIC_LIST", topics...)
}

func PTopicAdd(topic string) *Response {
	return NewResponse("TOPIC_ADD", topic)
}

func PTopicRemove(topic string) *Response {
	return NewResponse("TOPIC_REMOVE", topic)
}

func PVote(client *Client, waiting []string) *Response {
	var data []interface{}
	// add voter's name
	data = append(data, client.Name)
	// add waiting
	data = append(data, util.WrapStringArray(waiting...)...)
	return NewResponse("VOTE", data...)
}

func PPreferences(pref *GamePreferences) *Response {
	return NewResponse("PREFERENCES", pref)
}

func PChangeRole(client *Client, role string) *Response {
	return NewResponse("CHANGE_ROLE", client.Name, role)
}

func PPlayerLeave(client *Client, reason string) *Response {
	return NewResponse("PLAYER_LEAVE", client.Name, reason)
}
