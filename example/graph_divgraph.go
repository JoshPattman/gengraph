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
	_evwshoipg float64
	_evwshoipgGrad float64
	Y float64
	YGrad float64
}

func (g *DivGraph) ClearGrads() {
	g.YGrad = 0
	g._evwshoipgGrad = 0
	g.BGrad = 0
	g.AGrad = 0
}

func (g *DivGraph) Forward() {
	g._evwshoipg = g.A / g.B
	g.Y = g._evwshoipg
}

func (g *DivGraph) Backward() {
	g._evwshoipgGrad += g.YGrad
	g.AGrad += g._evwshoipgGrad * 1.0 / g.B
	g.BGrad += g._evwshoipgGrad * -g.A / (g.B * g.B)
}