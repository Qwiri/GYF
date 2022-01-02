package util

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

func (g GameState) Contains(state GameState) bool {
	return g&state == state
}

func (g GameState) In(other GameState) bool {
	return g&other == g
}
