package stringify

import (
	"github.com/Qwiri/GYF/backend/pkg/handler"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"strings"
)

func HandlerToString(bob *util.Bob, name string, handler *handler.Handler) {
	// write header
	bob.Writef("### %s", name)
	if handler.DevOnly {
		_, _ = bob.WriteString(" 🔰[^1]")
	}
	bob.NewLine(2)

	// write description
	if handler.Description != "" {
		bob.Writefl("> %s\n", strings.TrimSpace(handler.Description))
	}

	// write syntax
	bob.Alll('`', name,
		bob.If(handler.Syntax != "", " "),
		handler.Syntax,
		"`  ")
	bob.Alll("👉 args: ", handler.Bounds)

	// access
	bob.Writel(`!!! danger "Access"`)
	for _, role := range RoleOrder {
		a := Roles[role]
		bob.Alll('\t', "- [",
			bob.IfElse(handler.AccessLevel&a == a, "x", " "),
			"] ", role)
	}
	bob.NewLine()

	// states
	bob.Writel(`!!! hint "States"`)
	for _, stateName := range StateOrder {
		stateVal := States[stateName]
		bob.Alll('\t', "- [",
			bob.IfElse(stateVal.In(handler.StateLevel), "x", " "),
			"] ", stateName)
	}
}
