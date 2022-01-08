package util

import (
	"fmt"
	"strconv"
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

func (e *Bob) Writef(msg string, args ...interface{}) {
	_, _ = e.WriteString(fmt.Sprintf(msg, args...))
}

func (e *Bob) Writefl(msg string, args ...interface{}) {
	e.Writef(msg, args...)
	e.NewLine()
}

func (e *Bob) Writel(msg string) {
	_, _ = e.WriteString(msg)
	e.NewLine()
}

func (e *Bob) Bytes() []byte {
	return []byte(e.String())
}

func (*Bob) If(cond bool, str string) string {
	if cond {
		return str
	}
	return ""
}

func (*Bob) IfElse(cond bool, str, els string) string {
	if cond {
		return str
	}
	return els
}

func (e *Bob) All(str ...interface{}) {
	for _, s := range str {
		var sv string

		switch t := s.(type) {
		case string:
			sv = t
		case int:
			sv = strconv.Itoa(t)
		case fmt.Stringer:
			sv = t.String()
		// special
		case rune:
			_, _ = e.WriteRune(t)
			continue
		}

		if sv != "" {
			_, _ = e.WriteString(sv)
		}
	}
}

func (e *Bob) Alll(str ...interface{}) {
	e.All(str...)
	e.NewLine()
}
