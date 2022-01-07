package handler

import (
	"github.com/Qwiri/GYF/backend/pkg/util"
)

type Handler struct {
	Description string
	Syntax      string
	AccessLevel Access
	Bounds      util.Boundaries
	StateLevel  util.GameState
	Handler     interface{}
	DevOnly     bool
}

func (h *Handler) AcceptsState(state util.GameState) bool {
	return h.StateLevel.Contains(state)
}
