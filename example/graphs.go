package main

import G "github.com/JoshPattman/gengraph"

func CreateGraphs() {
	{
		// Create a new graph to build on, called CosGraph
		g := G.NewGraph("CosGraph")
		// Define an input variable, making sure to specify a capital letter for its first character so the field will be exported in the struct
		a := G.Variable[float64](g, "Input")
		// Define a constant variable, b, with a value of 3.0. The name of this variable (and any other unamed variables) will be generated.
		b := G.Constant(g, 3.0)
		// Add the input and the constant together
		added := G.NumAdd(g, a.Buf(), b.Buf())
		// Calculate the cosine of the sum
		res := G.NumCos(g, added.Buf())
		// Alias the result to "Result" - This allows us to specify a name for the result so we can easily acsess it from the struct
		G.Alias(g, res.Buf(), "Result")

		// Write the graph to the default file - graph_cosgraph.go
		g.ToDefaultFile()
	}
	{
		g := G.NewGraph("DivGraph")
		a := G.Variable[float64](g, "A")
		b := G.Variable[float64](g, "B")
		res := G.NumDiv(g, a.Buf(), b.Buf())
		G.Alias(g, res.Buf(), "Y")

		g.ToDefaultFile()
	}
}
