package gengraph

import "math/rand"

var randNameGen = rand.New(rand.NewSource(0))

func randStringName() string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	s := make([]byte, 10)
	for i := range s {
		s[i] = alpha[randNameGen.Intn(len(alpha))]
	}
	s[0] = '_'
	return string(s)
}
