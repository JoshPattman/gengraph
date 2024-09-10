//go:build !graph

package main

import (

)

func NewDivGraph() *DivGraph{
	g := &DivGraph{}
	return g
}

type DivGraph struct {
	A float64
	AGrad float64
	B float64
	BGrad float64
	_iwndnzofj float64
	_iwndnzofjGrad float64
	Y float64
	YGrad float64
}

func (g *DivGraph) ClearGrads() {
	g.YGrad = 0
	g._iwndnzofjGrad = 0
	g.BGrad = 0
	g.AGrad = 0
}

func (g *DivGraph) Forward() {
	g._iwndnzofj = g.A / g.B
	g.Y = g._iwndnzofj
}

func (g *DivGraph) Backward() {
	g._iwndnzofjGrad += g.YGrad
	g.AGrad += g._iwndnzofjGrad * 1.0 / g.B
	g.BGrad += g._iwndnzofjGrad * -g.A / (g.B * g.B)
}