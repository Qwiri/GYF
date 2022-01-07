package util

import (
	"fmt"
	"strings"
)

type Bob struct {
	strings.Builder
}

//goland:noinspection ALL
func (e *Bob) NewLine(repeat ...int) {
	if len(repeat) == 1 {
		for i := 0; i < repeat[0]; i++ {
			e.NewLine()
		}
		return
	}
	e.WriteRune('\n')
}

//goland:noinspection ALL
func (e *Bob) Writef(msg string, args ...interface{}) {
	e.WriteString(fmt.Sprintf(msg, args...))
}

func (e *Bob) Bytes() []byte {
	return []byte(e.String())
}
