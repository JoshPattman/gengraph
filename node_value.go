package gengraph

import (
	"fmt"
)

type Value[T any] struct {
	Var      *Buffer[T]
	isConst  bool
	constVal T
}

func Variable[T any](g *Graph, name string) *Value[T] {
	v := &Value[T]{Var: &Buffer[T]{Name: name, OnGraph: g}}
	g.Add(v)
	return v
}

func Constant[T any](g *Graph, val T) *Value[T] {
	v := &Value[T]{Var: &Buffer[T]{Name: randStringName(), OnGraph: g}, isConst: true, constVal: val}
	g.Add(v)
	return v
}

func (v *Value[T]) Buf() *Buffer[T] {
	return v.Var
}

func (v *Value[T]) FwdLines() []string {
	return []string{}
}

func (v *Value[T]) BufferDefs() []string {
	return []string{v.Var.BufferDef(), v.Var.GradBufferDef()}
}

func (v *Value[T]) BufferInits() []string {
	if v.isConst {
		return []string{
			fmt.Sprintf("%s = %v", v.Var.UseString(), v.constVal),
		}
	}
	return []string{}
}

func (v *Value[T]) Imports() []string {
	return []string{}
}

func (v *Value[T]) BackLines() []string {
	return []string{}
}

func (v *Value[T]) GradBufferClears() []string {
	return []string{
		fmt.Sprintf("%s = 0", v.Var.GradUseString()),
	}
}
