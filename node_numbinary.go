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

func NumAdd[T Numerical](left BufferGetter[T], right BufferGetter[T]) *NumBinaryNode[T] {
	return numBinary(left, right, nbAdd)
}

func NumSub[T Numerical](left BufferGetter[T], right BufferGetter[T]) *NumBinaryNode[T] {
	return numBinary(left, right, nbSub)
}

func NumMul[T Numerical](left BufferGetter[T], right BufferGetter[T]) *NumBinaryNode[T] {
	return numBinary(left, right, nbMul)
}

func NumDiv[T Numerical](left BufferGetter[T], right BufferGetter[T]) *NumBinaryNode[T] {
	return numBinary(left, right, nbDiv)
}

func numBinary[T Numerical](left BufferGetter[T], right BufferGetter[T], op numBinaryOp) *NumBinaryNode[T] {
	left.Buf().AssertScalarShape()
	right.Buf().AssertScalarShape()
	g := left.Buf().OnGraph
	if g != right.Buf().OnGraph {
		panic("input buffers must be on the same graph")
	}
	n := &NumBinaryNode[T]{
		fromLeft:  left.Buf(),
		fromRight: right.Buf(),
		to:        &Buffer[T]{Name: randStringName(), OnGraph: g},
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
