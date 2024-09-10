package gengraph

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Graph struct {
	Nodes []Node
	Name  string
}

func NewGraph(name string) *Graph {
	return &Graph{Name: name}
}

func (g *Graph) Add(n Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) ToDefaultFile() error {
	return g.ToFile(fmt.Sprintf("graph_%s.go", strings.ToLower(g.Name)))
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
	return strings.Join([]string{
		"//go:build !graph",
		"package main",
		g.importsString(),
		g.consString(),
		g.structString(),
		g.clearGradsString(),
		g.forwardString(),
		g.backwardString(),
	}, "\n\n")
}

func (g *Graph) importsString() string {
	imports := map[string]bool{}
	for _, n := range g.Nodes {
		for _, imp := range n.Imports() {
			imports[imp] = true
		}
	}
	var res []string
	for imp := range imports {
		res = append(res, "\t\""+imp+"\"")
	}
	return fmt.Sprintf("import (\n%s\n)", strings.Join(res, "\n"))
}

func (g *Graph) consString() string {
	body := []string{
		fmt.Sprintf("g := &%s{}", g.Name),
	}
	for _, n := range g.Nodes {
		body = append(body, n.BufferInits()...)
	}
	body = append(body, "return g")
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func New%s() *%s{\n%s\n}", g.Name, g.Name, strings.Join(body, "\n"))
}

func (g *Graph) structString() string {
	body := []string{}
	for _, n := range g.Nodes {
		body = append(body, n.BufferDefs()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("type %s struct {\n%s\n}", g.Name, strings.Join(body, "\n"))
}

func (g *Graph) forwardString() string {
	body := []string{}
	for _, n := range g.Nodes {
		body = append(body, n.FwdLines()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) Forward() {\n%s\n}", g.Name, strings.Join(body, "\n"))
}

func (g *Graph) backwardString() string {
	body := []string{}
	for _, n := range slices.Backward(g.Nodes) {
		body = append(body, n.BackLines()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) Backward() {\n%s\n}", g.Name, strings.Join(body, "\n"))
}

func (g *Graph) clearGradsString() string {
	body := []string{}
	for _, n := range slices.Backward(g.Nodes) {
		body = append(body, n.GradBufferClears()...)
	}
	body = indentStrings(body, "\t")
	return fmt.Sprintf("func (g *%s) ClearGrads() {\n%s\n}", g.Name, strings.Join(body, "\n"))
}

func indentStrings(s []string, indent string) []string {
	var res []string
	for _, line := range s {
		res = append(res, indent+line)
	}
	return res
}
