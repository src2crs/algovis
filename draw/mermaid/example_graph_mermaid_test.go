package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/datastructures/graph/basicgraph"
)

// ExampleMermaid_all_nodes_visible shows how to create a graph where all nodes are visible
// and how to draw it using the Mermaid function.
func ExampleMermaid_all_nodes_visible() {
	g := basicgraph.New()

	g.AddEdgeBetweenIds("a", "b")
	g.AddEdgeBetweenIds("b", "c")
	g.AddEdgeBetweenIds("c", "a")

	fmt.Println(Mermaid(g))

	// Output:
	// graph TD
	//   a[a]
	//   b[b]
	//   c[c]
	//   a --- b
	//   b --- c
	//   c --- a
}

// ExampleMermaid_one_node_invisible shows how to create a graph where one node is invisible
// and how to draw it using the Mermaid function.
func ExampleMermaid_one_node_invisible() {
	g := basicgraph.New()

	g.AddEdgeBetweenIds("a", "b")
	g.AddInvisibleNodeWithId("c")
	g.AddEdgeBetweenIds("b", "c")
	g.AddEdgeBetweenIds("c", "a")

	fmt.Println(Mermaid(g))

	// Output:
	// graph TD
	//   a[a]
	//   b[b]
	//   c[c]:::hidden
	//   a --- b
	//   b ~~~ c
	//   c ~~~ a
}
