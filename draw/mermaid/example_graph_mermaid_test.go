package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/datastructures/graph/basicgraph"
)

// ExampleMermaid_graph_all_nodes_visible shows how to create a graph where all nodes are visible
// and how to draw it using the Mermaid function.
func ExampleMermaid_graph_all_nodes_visible() {
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

// ExampleMermaid_graph_one_node_invisible shows how to create a graph where one node is invisible
// and how to draw it using the Mermaid function.
func ExampleMermaid_graph_one_node_invisible() {
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

// ExampleMermaid_binarytree_graph_complete shows how to create a binary tree using a graph.
// It also shows how to draw it using the Mermaid function.

// The tree looks like this:
//
//	    root
//	    /  \
//	   L    R
//	 / \   / \
//	LL LR RL RR
func ExampleMermaid_binarytree_graph_complete() {
	g := basicgraph.New()

	g.AddEdgeBetweenIds("root", "L")
	g.AddEdgeBetweenIds("root", "R")
	g.AddEdgeBetweenIds("L", "LL")
	g.AddEdgeBetweenIds("L", "LR")
	g.AddEdgeBetweenIds("R", "RL")
	g.AddEdgeBetweenIds("R", "RR")

	fmt.Println(Mermaid(g))

	// Output:
	// graph TD
	//   root[root]
	//   L[L]
	//   R[R]
	//   LL[LL]
	//   LR[LR]
	//   RL[RL]
	//   RR[RR]
	//   root --- L
	//   root --- R
	//   L --- LL
	//   L --- LR
	//   R --- RL
	//   R --- RR
}

// ExampleMermaid_binarytree_graph_incomplete shows how to create an incomplete binary tree using a graph.
// It also shows how to draw it using the Mermaid function.
func ExampleMermaid_binarytree_graph_incomplete() {
	g := basicgraph.New()

	g.AddEdgeBetweenIds("root", "L")
	g.AddEdgeBetweenIds("root", "R")
	g.AddInvisibleNodeWithId("LL")
	g.AddEdgeBetweenIds("L", "LL")
	g.AddEdgeBetweenIds("L", "LR")
	g.AddInvisibleNodeWithId("RL")
	g.AddEdgeBetweenIds("R", "RL")
	g.AddEdgeBetweenIds("R", "RR")

	fmt.Println(Mermaid(g))

	// Output:
	// graph TD
	//   root[root]
	//   L[L]
	//   R[R]
	//   LL[LL]:::hidden
	//   LR[LR]
	//   RL[RL]:::hidden
	//   RR[RR]
	//   root --- L
	//   root --- R
	//   L ~~~ LL
	//   L --- LR
	//   R ~~~ RL
	//   R --- RR
}
