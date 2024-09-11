package gengraph

import "fmt"

type BufferGetter[T any] interface {
	Buf() *Buffer[T]
}

type Buffer[T any] struct {
	Name    string
	OnGraph *Graph
	Shape   []int
}

func (v *Buffer[T]) BufferDef() string {
	return fmt.Sprintf("%s %T", v.Name, *new(T))
}

func (v *Buffer[T]) GradBufferDef() string {
	return fmt.Sprintf("%sGrad %T", v.Name, *new(T))
}

func (v *Buffer[T]) UseString() string {
	return "g." + v.Name
}

func (v *Buffer[T]) GradUseString() string {
	return "g." + v.Name + "Grad"
}

func (v *Buffer[T]) Buf() *Buffer[T] {
	return v
}

func (v *Buffer[T]) NumDims() int {
	return len(v.Shape)
}

func (v *Buffer[T]) AssertScalarShape() {
	if v.NumDims() != 0 {
		panic(fmt.Sprintf("Expected scalar shape (either nil or {}) but got %v", v.Shape))
	}
}
