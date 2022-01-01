package gerrors

import "errors"

var (
	// Game Errors
	ErrGameNotFound    = errors.New("game not found")
	ErrGameStarted     = errors.New("game already started")
	ErrGameStateAccess = errors.New("no access in that state")

	// Handler Errors
	ErrMessageTooShort = errors.New("message too short")
	ErrUnknownCommand  = errors.New("unknown command")
	ErrNoAccess        = errors.New("no access to that command")
	ErrInvalidHandler  = errors.New("invalid handler func")
	ErrDevOnly         = errors.New("handler is dev only")
	ErrArgLength       = errors.New("unexpected arg length")

	// Handler Specific Errors
	// chat
	ErrMessageEmpty = errors.New("message empty")
	// join
	ErrNameInvalid   = errors.New("name invalid")
	ErrNameExists    = errors.New("name already exists")
	ErrAlreadyJoined = errors.New("already joined")
	// topics
	ErrTooManyTopics      = errors.New("too many topics")
	ErrTopicNotFound      = errors.New("topic not found")
	ErrTopicAlreadyExists = errors.New("topic already exists")
	ErrNoTopicsLeft       = errors.New("no topics left")

	// Requirements
	ErrTooFewPlayers = errors.New("too few players")
	ErrTooFewTopics  = errors.New("too few topics")

	// Submissions
	ErrGIFTaken         = errors.New("another player took that gif already")
	ErrGIFNotAllowed    = errors.New("url not allowed")
	ErrAlreadySubmitted = errors.New("already submitted that GIF")

	// Voting
	ErrSubmissionNotFound = errors.New("submission not found")
	ErrAlreadyVoted       = errors.New("already voted")
	ErrVoteSelf           = errors.New("cannot vote self")
)
