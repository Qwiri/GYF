package util

import (
	"strconv"
	"strings"
)

type Boundaries []*BoundaryInt

type BoundaryInt struct {
	Min   *int
	Max   *int
	Exact *int
}

func (i *BoundaryInt) Applies(a int) bool {
	if i.Min != nil && a < *i.Min {
		return false
	}
	if i.Max != nil && a > *i.Max {
		return false
	}
	if i.Exact != nil && a != *i.Exact {
		return false
	}
	return true
}

func (B Boundaries) Applies(a int) bool {
	for _, b := range B {
		if !b.Applies(a) {
			return false
		}
	}
	return true
}

func Bounds(boundaries ...*BoundaryInt) Boundaries {
	var resp Boundaries
	for _, b := range boundaries {
		resp = append(resp, b)
	}
	return resp
}

func BoundMax(max int) *BoundaryInt {
	return &BoundaryInt{Max: Int(max)}
}
func BoundMin(min int) *BoundaryInt {
	return &BoundaryInt{Min: Int(min)}
}
func BoundExact(exact int) *BoundaryInt {
	return &BoundaryInt{Exact: Int(exact)}
}

func (i BoundaryInt) String() string {
	var bob strings.Builder
	bob.WriteRune('[')
	if i.Min != nil {
		bob.WriteString(strconv.Itoa(*i.Min))
		bob.WriteString(" <= ")
	}
	bob.WriteRune('x')
	if i.Max != nil {
		bob.WriteString(" <= ")
		bob.WriteString(strconv.Itoa(*i.Max))
	}
	if i.Exact != nil {
		bob.WriteString(" == ")
		bob.WriteString(strconv.Itoa(*i.Exact))
	}
	bob.WriteRune(']')
	return bob.String()
}

func (B Boundaries) String() string {
	str := make([]string, len(B))
	for i, b := range B {
		str[i] = b.String()
	}
	return strings.Join(str, ", ")
}
