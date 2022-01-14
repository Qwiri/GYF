package model

var DefaultPreferences = GamePreferences{
	AutoSkip:      true,
	MinPlayers:    3,
	MaxPlayers:    100,
	MinTopics:     1,
	MaxTopics:     30,
	ShuffleTopics: true,
	Permissions:   0,
}

type EnhancedPermission uint8

const (
	PermissionListTopics = 1 << iota
	PermissionCreateTopics
	PermissionDeleteTopics
)

type GamePreferences struct {
	AutoSkip      bool
	MinPlayers    int
	MaxPlayers    int
	MinTopics     int
	MaxTopics     int
	ShuffleTopics bool
	Permissions   EnhancedPermission
}
