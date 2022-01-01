package model

var DefaultPreferences = GamePreferences{
	AutoSkip:   true,
	MinPlayers: 3,
	MaxPlayers: 10,
	MinTopics:  1,
	MaxTopics:  30,
}

type GamePreferences struct {
	AutoSkip   bool
	MinPlayers int
	MaxPlayers int
	MinTopics  int
	MaxTopics  int
}
