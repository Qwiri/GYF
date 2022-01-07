package main

import (
	"github.com/Qwiri/GYF/backend/internal/handlers"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

const (
	commentStart = "[//]: # (handlers_start)"
	commentEnd   = "[//]: # (handlers_end)"
)

func main() {
	var (
		data []byte
		err  error
		stat os.FileInfo
	)
	join := path.Join("..", "docs", "socket.md")

	// stat
	if stat, err = os.Stat(join); err != nil {
		panic(err)
		return
	}

	// read docs file
	if data, err = os.ReadFile(join); err != nil {
		panic(err)
		return
	}

	var bob util.Bob
	capture := true
	lines := strings.Split(string(data), "\n")
	i := 0
	for _, line := range lines {
		i += 1
		if line == commentStart {
			capture = false

			_, _ = bob.WriteString(line)
			bob.NewLine()

			gen := generate()
			_, _ = bob.WriteString(gen.String())
			bob.NewLine()
			continue
		}

		if line == commentEnd && !capture {
			capture = true
		}

		if !capture {
			continue
		}

		_, _ = bob.WriteString(line)

		if i != len(lines) {
			bob.NewLine()
		}
	}

	// write Bob's output to docs file
	if err = os.WriteFile(join, bob.Bytes(), stat.Mode()); err != nil {
		panic(err)
	}
}

func generate() (bob util.Bob) {
	bob.Writef("*Generated on %s*", time.Now().Format("02.01.2006 15:04:05"))
	bob.NewLine(2)

	// collect names
	names := make([]string, len(handlers.Handlers))
	i := 0
	for name := range handlers.Handlers {
		names[i] = name
		i += 1
	}
	// order ascending
	sort.Strings(names)

	stateNames := make([]string, len(util.States))
	i = 0
	for name := range util.States {
		stateNames[i] = name
		i += 1
	}
	// order ascending
	sort.Strings(stateNames)

	for j, name := range names {
		if j != 0 {
			bob.NewLine()
			_, _ = bob.WriteString("---")
			bob.NewLine(2)
		}

		handler := handlers.Handlers[name]

		bob.Writef("### %s", name)
		if handler.DevOnly {
			_, _ = bob.WriteString(" 🔰[^1]")
		}
		bob.NewLine(2)

		// Add description
		if handler.Description != "" {
			bob.Writef("> %s", handler.Description)
			bob.NewLine(2)
		}

		// Add syntax
		bob.All("`", name, bob.If(handler.Syntax != "", " "), handler.Syntax, "`")
		bob.NewLine(2)

		// Access
		for _, stateName := range stateNames {
			stateVal := util.States[stateName]
			bob.Alll("- [", bob.IfElse(handler.AcceptsState(stateVal), "x", " "), "] ", stateName)
		}
	}

	return
}
