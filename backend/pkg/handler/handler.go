package handler

import (
	"github.com/Qwiri/GYF/backend/pkg/util"
)

type Handler struct {
	AccessLevel Access
	Bounds      util.Boundaries
	StateLevel  util.GameState
	Handler     interface{}
	DevOnly     bool
}
