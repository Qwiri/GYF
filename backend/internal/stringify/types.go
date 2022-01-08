package stringify

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/util"
)

// Role

var Roles = map[string]handler.Access{
	"Guest":  handler.AccessGuest,
	"Joined": handler.AccessPlayer,
	"Leader": handler.AccessLeader,
}

var RoleOrder = []string{
	"Guest",
	"Joined",
	"Leader",
}

// States

var States = map[string]util.GameState{
	"Lobby":             util.StateLobby,
	"Submit GIF":        util.StateSubmitGIF,
	"Cast Votes":        util.StateCastVotes,
	"Show Vote Results": util.StateShowVotes,
}

var StateOrder = []string{
	"Lobby",
	"Submit GIF",
	"Cast Votes",
	"Show Vote Results",
}
