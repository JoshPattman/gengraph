package gengraph

import "fmt"

type Shape []int

func (s Shape) IsScalar() bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 && s[0] == 0 {
		return true
	}
	return false
}

func (a Shape) Equals(b Shape) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// As a special case, all scalars return 0, even if they have length 1.
func (s Shape) NumDims() int {
	if s.IsScalar() {
		return 0
	}
	return len(s)
}

func (s Shape) AssertScalar() {
	if !s.IsScalar() {
		panic(fmt.Sprintf("Expected scalar shape (either nil, {}, or {1}) but got %v", s))
	}
}

func (s Shape) AssertEquals(other Shape) {
	if !s.Equals(other) {
		panic(fmt.Sprintf("Expected equal shape %v but got %v", other, s))
	}
}
