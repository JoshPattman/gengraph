package gengraph

import "fmt"

type numBinaryOp int

const (
	nbAdd numBinaryOp = iota
	nbSub
	nbMul
	nbDiv
)

type NumBinaryNode[T Numerical] struct {
	fromLeft  *Buffer[T]
	fromRight *Buffer[T]
	to        *Buffer[T]
	op        numBinaryOp
}

func NumAdd[T Numerical](g *Graph, left *Buffer[T], right *Buffer[T]) *NumBinaryNode[T] {
	return numBinary(g, left, right, nbAdd)
}

func NumSub[T Numerical](g *Graph, left *Buffer[T], right *Buffer[T]) *NumBinaryNode[T] {
	return numBinary(g, left, right, nbSub)
}

func NumMul[T Numerical](g *Graph, left *Buffer[T], right *Buffer[T]) *NumBinaryNode[T] {
	return numBinary(g, left, right, nbMul)
}

func NumDiv[T Numerical](g *Graph, left *Buffer[T], right *Buffer[T]) *NumBinaryNode[T] {
	return numBinary(g, left, right, nbDiv)
}

func numBinary[T Numerical](g *Graph, left *Buffer[T], right *Buffer[T], op numBinaryOp) *NumBinaryNode[T] {
	n := &NumBinaryNode[T]{
		fromLeft:  left,
		fromRight: right,
		to:        &Buffer[T]{Name: randStringName()},
		op:        op,
	}
	g.Add(n)
	return n
}

func (n *NumBinaryNode[T]) FwdLines() []string {
	opStr := ""
	switch n.op {
	case nbAdd:
		opStr = "+"
	case nbSub:
		opStr = "-"
	case nbMul:
		opStr = "*"
	case nbDiv:
		opStr = "/"
	default:
		panic("unreachable")
	}
	return []string{fmt.Sprintf("%s = %s %s %s", n.to.UseString(), n.fromLeft.UseString(), opStr, n.fromRight.UseString())}
}

func (n *NumBinaryNode[T]) BufferDefs() []string {
	return []string{n.to.BufferDef(), n.to.GradBufferDef()}
}

func (n *NumBinaryNode[T]) BufferInits() []string { return nil }

func (n *NumBinaryNode[T]) Imports() []string { return nil }

func (n *NumBinaryNode[T]) Buf() *Buffer[T] { return n.to }

func (v *NumBinaryNode[T]) BackLines() []string {
	tGUS := v.to.GradUseString()
	lGUS := v.fromLeft.GradUseString()
	rGUS := v.fromRight.GradUseString()
	lUS := v.fromLeft.UseString()
	rUS := v.fromRight.UseString()
	switch v.op {
	case nbAdd:
		return []string{
			fmt.Sprintf("%s += %s", lGUS, tGUS),
			fmt.Sprintf("%s += %s", rGUS, tGUS),
		}
	case nbSub:
		return []string{
			fmt.Sprintf("%s += %s", lGUS, tGUS),
			fmt.Sprintf("%s -= %s", rGUS, tGUS),
		}
	case nbMul:
		return []string{
			fmt.Sprintf("%s += %s * %s", lGUS, tGUS, rUS),
			fmt.Sprintf("%s += %s * %s", rGUS, tGUS, lUS),
		}
	case nbDiv:
		// TODO check this is correct
		return []string{
			fmt.Sprintf("%s += %s * 1.0 / %s", lGUS, tGUS, rUS),
			fmt.Sprintf("%s += %s * -%s / (%s * %s)", rGUS, tGUS, lUS, rUS, rUS),
		}
	}
	panic("unreachable")
}

func (v *NumBinaryNode[T]) GradBufferClears() []string {
	return []string{
		fmt.Sprintf("%s = 0", v.to.GradUseString()),
	}
}
