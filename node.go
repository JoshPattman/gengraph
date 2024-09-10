package gengraph

import (
	"math/rand"
)

type Numerical interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Node interface {
	FwdLines() []string
	BackLines() []string
	BufferDefs() []string
	BufferInits() []string
	GradBufferClears() []string
	Imports() []string
}

func randStringName() string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	s := make([]byte, 10)
	for i := range s {
		s[i] = alpha[rand.Intn(len(alpha))]
	}
	s[0] = '_'
	return string(s)
}
