package handler

import (
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
)

type Handler struct {
	Description  string
	Syntax       string
	AccessLevel  Access
	EnhancedPerm model.EnhancedPermission
	Bounds       util.Boundaries
	StateLevel   util.GameState
	Handler      interface{}
	DevOnly      bool
}
