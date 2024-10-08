# Example useage of `gengraph`
This example demonstrates how to use `gengraph` to generate graphs in a Go project. It also then benchmarks the graphs against both a native go implementation that I wrote manually, and my previous attempt at a graph computation library, `toygraph`. Spoiler altert: `gengraph` is `20x` times faster than `toygraph` in this test. The native implementaion is only `1.4x` faster than `gengraph`, but it requires you to manually write the forward and backward passes.

File overview:
- `main.go`: The main file that contains the main program logic, run when `$ go run .`. It is exluded from the build when building graphs.
- `main_gen.go`: Alternative main file, used only when generating graphs. It calls the `CreateGraphs` function, which lives in a different file, as syntax highlighting may not work in this file.
- `graphs.go`: Contains the `CreateGraphs` function, which generates the graphs.
- `graph_cosgraph.go` & `graph_divgraph.go`: Contains the `CosGraph` and `DivGraph` structs that were generated by `gengraph`.
- `Makefile`: Contains the build instructions for the project. It builds the project and generates the graphs.