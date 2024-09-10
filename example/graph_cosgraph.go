//go:build !graph

package main

import (
	"math"
)

func NewCosGraph() *CosGraph{
	g := &CosGraph{}
	g._eegpngybr = 3
	return g
}

type CosGraph struct {
	Input float64
	InputGrad float64
	_eegpngybr float64
	_eegpngybrGrad float64
	_wfexjqgcd float64
	_wfexjqgcdGrad float64
	_jhoaiemtz float64
	_jhoaiemtzGrad float64
	Result float64
	ResultGrad float64
}

func (g *CosGraph) ClearGrads() {
	g.ResultGrad = 0
	g._jhoaiemtzGrad = 0
	g._wfexjqgcdGrad = 0
	g._eegpngybrGrad = 0
	g.InputGrad = 0
}

func (g *CosGraph) Forward() {
	g._wfexjqgcd = g.Input + g._eegpngybr
	g._jhoaiemtz = math.Cos(g._wfexjqgcd)
	g.Result = g._jhoaiemtz
}

func (g *CosGraph) Backward() {
	g._jhoaiemtzGrad += g.ResultGrad
	g._wfexjqgcdGrad += -math.Sin(g._wfexjqgcd) * g._jhoaiemtzGrad
	g.InputGrad += g._wfexjqgcdGrad
	g._eegpngybrGrad += g._wfexjqgcdGrad
}