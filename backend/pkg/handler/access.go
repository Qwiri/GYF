package handler

import "github.com/Qwiri/GYF/backend/pkg/model"

type Access uint8

const (
	AccessGuest Access = 1 << iota
	AccessPlayer
	AccessLeader
)

const (
	AccessJoined = AccessPlayer | AccessLeader
	AccessAny    = AccessGuest | AccessJoined
)

func (a Access) Allowed(client *model.Client) bool {
	// client is a guest
	if client == nil {
		return a&AccessGuest == AccessGuest
	}
	if a&AccessPlayer == AccessPlayer {
		// access requires joined access (if client has a name, the client has access)
		return client.Name != ""
	}
	// access requires leader access
	if a&AccessLeader == AccessLeader {
		return client.Leader
	}
	return false
}
