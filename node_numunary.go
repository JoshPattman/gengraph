package gengraph

import "fmt"

type numUnaryOp int

const (
	nuSin numUnaryOp = iota
	nuCos
	numRelu
)

type NumUnaryNode[T Numerical] struct {
	from *Buffer[T]
	to   *Buffer[T]
	op   numUnaryOp
}

func NumSin[T Numerical](g *Graph, in *Buffer[T]) *NumUnaryNode[T] {
	return numUnary(g, in, nuSin)
}

func NumCos[T Numerical](g *Graph, in *Buffer[T]) *NumUnaryNode[T] {
	return numUnary(g, in, nuCos)
}

func numUnary[T Numerical](g *Graph, in *Buffer[T], op numUnaryOp) *NumUnaryNode[T] {
	n := &NumUnaryNode[T]{
		from: in,
		to:   &Buffer[T]{Name: randStringName()},
		op:   op,
	}
	g.Add(n)
	return n
}

func (n *NumUnaryNode[T]) FwdLines() []string {
	if n.isMathOp() {
		opStr := ""
		switch n.op {
		case nuSin:
			opStr = "math.Sin"
		case nuCos:
			opStr = "math.Cos"
		default:
			panic("unreachable")
		}
		return []string{fmt.Sprintf("%s = %s(%s)", n.to.UseString(), opStr, n.from.UseString())}
	} else {
		switch n.op {
		case numRelu:
			return []string{
				fmt.Sprintf("if %s < 0 {", n.from.UseString()),
				fmt.Sprintf("  %s = 0", n.to.UseString()),
				"} else {",
				fmt.Sprintf("  %s = %s", n.to.UseString(), n.from.UseString()),
				"}",
			}
		default:
			panic("unreachable")
		}
	}
}

func (n *NumUnaryNode[T]) BufferDefs() []string {
	return []string{n.to.BufferDef(), n.to.GradBufferDef()}
}

func (n *NumUnaryNode[T]) BufferInits() []string { return nil }

func (n *NumUnaryNode[T]) Imports() []string {
	if n.isMathOp() {
		return []string{"math"}
	}
	return nil
}

func (n *NumUnaryNode[T]) Buf() *Buffer[T] { return n.to }

func (v *NumUnaryNode[T]) BackLines() []string {
	if v.isMathOp() {
		opStr := ""
		switch v.op {
		case nuSin:
			opStr = "math.Cos"
		case nuCos:
			opStr = "-math.Sin"
		default:
			panic("unreachable")
		}
		return []string{fmt.Sprintf("%s += %s(%s) * %s", v.from.GradUseString(), opStr, v.from.UseString(), v.to.GradUseString())}
	} else {
		switch v.op {
		case numRelu:
			return []string{
				fmt.Sprintf("if %s > 0 {", v.from.UseString()),
				fmt.Sprintf("  %s += %s", v.from.GradUseString(), v.to.GradUseString()),
				"}",
			}
		default:
			panic("unreachable")
		}
	}
}

func (v *NumUnaryNode[T]) GradBufferClears() []string {
	return []string{
		fmt.Sprintf("%s = 0", v.to.GradUseString()),
	}
}

func (n *NumUnaryNode[T]) isMathOp() bool {
	switch n.op {
	case nuSin, nuCos:
		return true
	}
	return false
}
