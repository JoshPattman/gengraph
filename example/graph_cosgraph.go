package main

import (
	"math"
)

func NewCosGraph() *CosGraph{
	g := &CosGraph{}
	g._ubyhizzka = 3
	return g
}

type CosGraph struct {
	Input float64
	InputGrad float64
	_ubyhizzka float64
	_ubyhizzkaGrad float64
	_bleeansoc float64
	_bleeansocGrad float64
	_mignckyrw float64
	_mignckyrwGrad float64
	Result float64
	ResultGrad float64
}

func (g *CosGraph) ClearGrads() {
	g.ResultGrad = 0
	g._mignckyrwGrad = 0
	g._bleeansocGrad = 0
	g._ubyhizzkaGrad = 0
	g.InputGrad = 0
}

func (g *CosGraph) Forward() {
	g._bleeansoc = g.Input + g._ubyhizzka
	g._mignckyrw = math.Cos(g._bleeansoc)
	g.Result = g._mignckyrw
}

func (g *CosGraph) Backward() {
	g._mignckyrwGrad += g.ResultGrad
	g._bleeansocGrad += -math.Sin(g._bleeansoc) * g._mignckyrwGrad
	g.InputGrad += g._bleeansocGrad
	g._ubyhizzkaGrad += g._bleeansocGrad
}