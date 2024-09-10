# gengraph: High-Performance Code Generation for Computational Graphs in Go

[gengraph](https://github.com/JoshPattman/gengraph) is a small project that allows you to build simple computational graphs in Go. Unlike other graph computation frameworks, `gengraph` leverages a code generation step to convert your graph into pure Go code. This minimizes runtime overhead for forward and gradient propagation, making it ideal for performance-critical applications such as robotics.

I built `gengraph` as a further learning excersise in computational graphs, following on from my other graph computation project, `github.com/JoshPattman/toygraph`. I wanted to see how I could improve the performance of the graph computation by using code generation, and I'm happy with the results so far.

**Note:** `gengraph` does not, and probably never will, support GPU processing. However, I plan to add support for matrices and vectors using `gonum` in the future.

## Usage

To generate your graphs using `gengraph`, you write them as Go files in the same project where you intend to use the graphs. Build tags are utilized to streamline the process.

### Steps:

1. **Set up build tags**:
   - Add `//go:build !graph` to the top of your main Go file (with `func main()`). This excludes it during graph generation.
   
     Example:
     ```go
     //go:build !graph

     package main

     func main() {
         // Your main program logic
     }
     ```

   - Functions in your main file cannot be used for graph generation. Move any such functions to another file without build constraints.

2. **Create a graph generation file**:
   - Create a file called `main_generate.go` with the following content:
     ```go
     //go:build graph

     package main

     func main() {
         CreateGraphs()
     }
     ```

   - The purpose of this file is to call the `CreateGraphs` function, which should live in a separate file that has no build constraints.

3. **Define `CreateGraphs` function**:
   - It's fine to put the `CreateGraphs()` function in `main_generate.go`. However, for better syntax checking in editors like VS Code, it's recommended to place `CreateGraphs()` in another file that isn't excluded by build constraints.

4. **Run the project**:
   - Use the following command to generate the graphs and run your project:
     ```bash
     $ go run -tags graph . && go run .
     ```

### Example

For a more detailed example, refer to the `example` subdirectory in this repository. However, here’s a quick overview of what might exist in the `CreateGraphs` function:

This code creates a new graph, named `CosGraph`, that performs the function `cos(x + 3)`. It then saves the generated code to a file in the local Go project called `graph_cosgraph.go`:

```go
// Create a new graph to build on, called CosGraph
g := G.NewGraph("CosGraph")

// Define an input variable, making sure to specify a capital letter for its first character so the field will be exported in the struct
a := G.Variable[float64](g, "Input")

// Define a constant variable, b, with a value of 3.0. The name of this variable (and any other unamed variables) will be generated.
b := G.Constant(g, 3.0)

// Add the input and the constant together
added := G.NumAdd(a, b)

// Calculate the cosine of the sum
res := G.NumCos(added)

// Alias the result to "Result" - This allows us to specify a name for the result so we can easily acsess it from the struct
G.Alias(res, "Result")

// Write the graph to the default file - graph_cosgraph.go
g.ToDefaultFile()
```

To use this graph in the `main` function, here is some example code:

```go
// Create a new instance of our cos graph
// (the NewCosGraph function is generated and stored in graph_cosgraph.go)
g := NewCosGraph()

// Set the input value to 5.0
// If your graph has multiple inputs, you can set them all by setting their respective struct variables
g.Input = 5.0

// Run the forward pass (calculate cos(5.0 + 3.0))
g.Forward()
fmt.Println("Result=", g.Result)

// Clear the gradients
g.ClearGrads()

// gengraph computes the gradients w.r.t to output gradient, not the partial gradients.
// So we need to set the output gradient to 1.0 to see how a change in each input creates a change of 1.0 in the output.
// Side note: If you want to calculate the partials, you can set all grads to 0 except for the output of the function you want to calculate (can use an alias to get named acsess), then run the backward pass.
g.ResultGrad = 1.0

// Run the backward pass, storing gradients at each step
g.Backward()

// Every variable in gengraph has a corresponding gradient variable, suffixed with "Grad"
fmt.Println("InputGrad=", g.InputGrad)
```

By following these steps, you can seamlessly integrate graph generation into your Go project with minimal runtime overhead. If you're working on performance-critical applications, such as robotics, gengraph may provide the efficiency boost you're looking for!