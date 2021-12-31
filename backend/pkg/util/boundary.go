package util

type Boundaries []*BoundaryInt

type BoundaryInt struct {
	Min   *int
	Max   *int
	Exact *int
}

func (b *BoundaryInt) Applies(a int) bool {
	if b.Min != nil && a < *b.Min {
		return false
	}
	if b.Max != nil && a > *b.Max {
		return false
	}
	if b.Exact != nil && a != *b.Exact {
		return false
	}
	return true
}

func (by Boundaries) Applies(a int) bool {
	for _, b := range by {
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
