package gengraph

import "fmt"

type Buffer[T any] struct {
	Name string
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
