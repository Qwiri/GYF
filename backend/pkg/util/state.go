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

var States = map[string]GameState{
	"1. Lobby":             StateLobby,
	"2. Submit GIF":        StateSubmitGIF,
	"3. Cast Votes":        StateCastVotes,
	"4. Show Vote Results": StateShowVotes,
}

func (g GameState) Contains(state GameState) bool {
	return g&state == state
}

func (g GameState) In(other GameState) bool {
	return g&other == g
}
