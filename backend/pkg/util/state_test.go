package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	assert.True(t, StateShowVotes.Contains(StateShowVotes))
	assert.True(t, StateShowVotes.In(StateShowVotes))

	assert.True(t, (StateShowVotes | StateSubmitGIF).Contains(StateShowVotes))
	assert.False(t, StateShowVotes.Contains(StateShowVotes|StateSubmitGIF))

	assert.True(t, StateShowVotes.In(StateShowVotes|StateSubmitGIF))
	assert.False(t, (StateShowVotes | StateSubmitGIF).In(StateShowVotes))

	assert.True(t, (StateShowVotes | StateSubmitGIF).In(StateShowVotes|StateSubmitGIF))

	assert.True(t, StateAny.Contains(StateShowVotes))
	assert.False(t, (StateAny & ^StateShowVotes).Contains(StateShowVotes))
	assert.False(t, StateShowVotes.In(StateAny & ^StateShowVotes))
}
