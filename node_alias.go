package gengraph

import "fmt"

type AliasNode[T any] struct {
	From *Buffer[T]
	To   *Buffer[T]
}

func Alias[T any](g *Graph, from *Buffer[T], name string) *AliasNode[T] {
	n := &AliasNode[T]{From: from, To: &Buffer[T]{Name: name}}
	g.Add(n)
	return n
}

func (n *AliasNode[T]) FwdLines() []string {
	return []string{
		fmt.Sprintf("%s = %s", n.To.UseString(), n.From.UseString()),
	}
}

func (n *AliasNode[T]) BufferDefs() []string {
	return []string{n.To.BufferDef(), n.To.GradBufferDef()}
}

func (n *AliasNode[T]) BufferInits() []string {
	return []string{}
}

func (n *AliasNode[T]) Imports() []string {
	return []string{}
}

func (v *AliasNode[T]) BackLines() []string {
	return []string{
		fmt.Sprintf("%s += %s", v.From.GradUseString(), v.To.GradUseString()),
	}
}

func (v *AliasNode[T]) GradBufferClears() []string {
	return []string{
		fmt.Sprintf("%s = 0", v.To.GradUseString()),
	}
}
