//go:build !graph

package main

import (
	"fmt"
	"math"
	"time"

	"github.com/JoshPattman/toygraph"
)

func main() {
	fmt.Println("Differenting y=a/b with a=1, b=2")

	{
		g := NewDivGraph()
		g.A = 1.0
		g.B = 2.0
		g.Forward()
		fmt.Println("Y=", g.Y)
		g.YGrad = 1.0
		g.Backward()
		fmt.Println("AGrad=", g.AGrad, "BGrad=", g.BGrad)
	}

	fmt.Println("\nDifferenting y=cos(x+3) with timing")

	{
		tStart := time.Now()
		total := 0.0
		totalGrad := 0.0
		for i := 0.0; i < 1000000; i++ {
			total += baseline(i)
			totalGrad += baselineDiff(i, 1.0)
		}
		fmt.Println("BASELINE: Elapsed time=", time.Since(tStart), "total=", total, "totalGrad=", totalGrad)

		g := NewCosGraph()
		tStart = time.Now()
		total = 0.0
		totalGrad = 0.0
		for i := 0.0; i < 1000000; i++ {
			g.Input = i
			g.Forward()
			total += g.Result
			g.ClearGrads()
			g.ResultGrad = 1.0
			g.Backward()
			totalGrad += g.InputGrad
		}
		fmt.Println("GENGRAPH: Elapsed time=", time.Since(tStart), "total=", total, "totalGrad=", totalGrad)

		g2 := toygraph.NewGraph()
		a := toygraph.NumValue[float64](g2)
		b3 := toygraph.NumValue[float64](g2)
		toygraph.Set(b3.C, 3.0)
		added := toygraph.NumAdd(g2, a.C, b3.C)
		res := toygraph.NumCos(g2, added.C)
		tStart = time.Now()
		total = 0.0
		for i := 0.0; i < 1000000; i++ {
			toygraph.Set(a.C, i)
			g2.AllForward()
			total += toygraph.Get(res.C)
			g2.AllZeroGrads()
			toygraph.SetGrad(res.C, 1.0)
			g2.AllBackward()
			totalGrad += toygraph.GetGrad(a.C)
		}
		fmt.Println("GENGRAPH: Elapsed time=", time.Since(tStart), "total=", total, "totalGrad=", totalGrad)
	}
}

func baseline(x float64) float64 {
	return math.Cos(x + 3)
}

func baselineDiff(x float64, resGrad float64) float64 {
	return -math.Sin(x+3) * resGrad
}
