package gerrors

import "errors"

var (
	// ErrGameNotFound is returned when a game id was specified but the according game was not found
	ErrGameNotFound = errors.New("game not found")
	// ErrGameStateAccess is returned when a client tries to execute a command that is disabled in the current state
	// e.g. vote is disabled in the lobby
	ErrGameStateAccess = errors.New("no access in that state")
	// ErrPayloadTooLarge is returned when a client sent a payload larger than 4 KB
	ErrPayloadTooLarge = errors.New("payload too large")

	// ErrMessageTooShort is returned when a client sends an empty websocket message
	ErrMessageTooShort = errors.New("message too short")
	// ErrUnknownCommand is returned when a client requests an unknown command
	ErrUnknownCommand = errors.New("unknown command")
	// ErrNoAccess is returned when a client tries to execute a command which he has no access to
	// e.g. if a player requests a leader-only command
	ErrNoAccess = errors.New("no access to that command")
	// ErrInvalidHandler is returned when a handler has an invalid handler-function
	// valid: handler.BasicHandler, handler.MessagedHandler, handler.PrefixedHandler, handler.PrefixedMessagedHandler
	ErrInvalidHandler = errors.New("invalid handler func")
	// ErrDevOnly is returned if a handler was marked dev-only and the server runs in a production environment
	ErrDevOnly = errors.New("handler is dev only")
	// ErrArgLength is returned when a handler's bounds requirements were not met
	ErrArgLength = errors.New("unexpected arg length")

	ErrTypeInvalid       = errors.New("type invalid")
	ErrUnknownPreference = errors.New("unknown preference")

	// ErrNameInvalid is used by the join handler and is returned when the requested username doesn't meet the requirements
	ErrNameInvalid = errors.New("name invalid")
	// ErrNameExists is used by the join handler and is returned when the requested username was already occupied by another player
	ErrNameExists = errors.New("name already exists")
	// ErrAlreadyJoined is used by the join handler and is returned when a client tries to join a game which the client already joined
	ErrAlreadyJoined = errors.New("already joined")

	// ErrTopicNotFound is returned when a client tries to remove a topic that doesn't exist
	ErrTopicNotFound = errors.New("topic not found")
	// ErrTopicAlreadyExists is returned when a client tries to add a topic that already exists
	ErrTopicAlreadyExists = errors.New("topic already exists")
	// ErrNoTopicsLeft is returned when a next round should be started but no more topics are left and thus the game should be stopped
	ErrNoTopicsLeft = errors.New("no topics left")

	// ErrTooFewPlayers is returned when a client tries to start the game and too few players are in the lobby
	ErrTooFewPlayers = errors.New("too few players")
	// ErrTooManyPlayers is returned when a client tries to start the game and too many players are in the lobby
	ErrTooManyPlayers = errors.New("too many players")
	// ErrTooFewTopics is returned when a client tries to start the game and too few topics were added
	ErrTooFewTopics = errors.New("too few topics")
	// ErrTooManyTopics is returned when a client tries to start the game and too many topics were added
	ErrTooManyTopics = errors.New("too many topics")

	// ErrGIFTaken is returned when a client tries to submit a GIF, but the GIF was already submitted by another player
	ErrGIFTaken = errors.New("another player took that gif already")
	// ErrGIFNotAllowed is returned when a client tries to submit a GIF from an unauthorized URL
	ErrGIFNotAllowed = errors.New("url not allowed")
	// ErrAlreadySubmitted is returned when a client tries to submit a GIF which the client already submitted
	ErrAlreadySubmitted = errors.New("already submitted that GIF")

	// ErrSubmissionNotFound is used by the vote handler and is returned when clients try to vote on an unknown GIF
	ErrSubmissionNotFound = errors.New("submission not found")
	// ErrAlreadyVoted is used by the vote handler and is returned when clients try to vote again
	ErrAlreadyVoted = errors.New("already voted")
	// ErrVoteSelf is used by the vote handler and is returned when clients try to vote on their own GIFs
	ErrVoteSelf = errors.New("cannot vote self")

	ErrChatMessageTooShort = errors.New("chat message too short")
	ErrChatMessageTooLong  = errors.New("chat message too long")
	ErrChatMessageTooFast  = errors.New("typing too fast")
	ErrChatMessageRepeat   = errors.New("chat message repeated")

	ErrInvalidPassword = errors.New("password mismatch")
	ErrClientNotFound  = errors.New("client not found")
	ErrSelf            = errors.New("self")
)
