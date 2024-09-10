package gengraph

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Graph struct {
	nodes     []Node
	name      string
	jacobians []*jacobian
}

func NewGraph(name string) *Graph {
	return &Graph{name: name}
}

func (g *Graph) Add(n Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) ToDefaultFile() error {
	return g.ToFile(fmt.Sprintf("graph_%s.go", strings.ToLower(g.name)))
}

func (g *Graph) ToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(g.String())
	return err
}

func (g *Graph) String() string {
	parts := []string{
		//"//go:build !graph",
		"package main",
		g.importsString(),
		g.consString(),
		g.structString(),
		g.clearGradsString(),
		g.forwardString(),
		g.backwardString(),
	}
	for _, jac := range g.jacobians {
		parts = append(parts, jac.String())
	}
	return strings.Join(parts, "\n\n")
}

func (g *Graph) importsString() string {
	imports := map[string]bool{}
	for _, n := range g.nodes {
		for _, imp := range n.Imports() {
			imports[imp] = true
		}
	}
	if len(g.jacobians) > 0 {
		imports["gonum.org/v1/gonum/mat"] = true
	}
	var res []string
	for imp := range imports {
		res = append(res, "\t\""+imp+"\"")
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(res, "\n"))
}

func (g *Graph) consString() string {
	body := []string{
		fmt.Sprintf("g := &%s{}", g.name),
	}
	for _, n := range g.nodes {
		body = append(body, n.BufferInits()...)
	}
	body = append(body, "return g")
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func New%s() *%s{\n%s\n}", g.name, g.name, strings.Join(body, "\n"))
}

func (g *Graph) structString() string {
	body := []string{}
	for _, n := range g.nodes {
		body = append(body, n.BufferDefs()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("type %s struct {\n%s\n}", g.name, strings.Join(body, "\n"))
}

func (g *Graph) forwardString() string {
	body := []string{}
	for _, n := range g.nodes {
		body = append(body, n.FwdLines()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) Forward() {\n%s\n}", g.name, strings.Join(body, "\n"))
}

func (g *Graph) backwardString() string {
	body := []string{}
	for _, n := range slices.Backward(g.nodes) {
		body = append(body, n.BackLines()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) Backward() {\n%s\n}", g.name, strings.Join(body, "\n"))
}

func (g *Graph) clearGradsString() string {
	body := []string{}
	for _, n := range slices.Backward(g.nodes) {
		body = append(body, n.GradBufferClears()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) ClearGrads() {\n%s\n}", g.name, strings.Join(body, "\n"))
}

func (g *Graph) IncludeJacobian(name string, spaceDims, jointDims []BufferGetter[float64]) {
	g.jacobians = append(g.jacobians, &jacobian{name, g, spaceDims, jointDims})
}

func indentStrings(s []string, indent string) []string {
	var res []string
	for _, line := range s {
		res = append(res, indent+line)
	}
	return res
}

type jacobian struct {
	name      string
	graph     *Graph
	spaceDims []BufferGetter[float64]
	jointDims []BufferGetter[float64]
}

func (j *jacobian) String() string {
	body := []string{
		fmt.Sprintf("jac.Zero()"),
	}
	for spaceRow := range j.spaceDims {
		// Clear gradients
		body = append(body, "", "g.ClearGrads()")
		// Set all space grads to 0 except the current row
		for i, spaceVar := range j.spaceDims {
			val := 0
			if i == spaceRow {
				val = 1
			}
			body = append(body, fmt.Sprintf("%s = %d", spaceVar.Buf().GradUseString(), val))
		}
		// Backward pass
		body = append(body, "g.Backward()")
		// Set the jacobian row
		for jointCol, jointVar := range j.jointDims {
			body = append(body, fmt.Sprintf("jac.Set(%d, %d, %s)", spaceRow, jointCol, jointVar.Buf().GradUseString()))
		}
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s)%sJacobian(jac *mat.Dense) {\n%s\n}", j.graph.name, j.name, strings.Join(body, "\n"))
}
