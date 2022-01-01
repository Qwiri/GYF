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

func PState(state GameState) *Response {
	return NewResponse("STATE", state)
}

func PVoteResults(results ...*SubmissionResult) *Response {
	return NewResponse("VOTE_RESULTS", util.WrapVoteResults(results...)...)
}
