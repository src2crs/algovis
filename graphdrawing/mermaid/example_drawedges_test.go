package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/binarytrees"
)

// ExampleDrawEdge_binarytree creates a binary tree with several nodes
// and draws the mermaid representationsof the edges using DrawEdge().
func ExampleDrawEdge_binarytree() {
	// Create a binary tree with the following structure:
	//   root
	//   ├── L
	//   │   ├── LL
	//   │   │   ├── LLL (invisible edge)
	//   │   │   └── LLR (invisible edge)
	//   │   └── LR
	//   | 	 	 ├── LRL (invisible edge)
	//   | 	 	 ├── LRR (invisible edge)
	//   └── R
	//       └── RL (invisible edge)
	//       └── RR (invisible edge)
	//
	root := binarytrees.NewRootNode()
	l := root.CreateChild("L")
	ll := l.CreateChild("L")
	lll := ll.CreateChild("L")
	llr := ll.CreateChild("R")
	lr := l.CreateChild("R")
	lrl := lr.CreateChild("L")
	lrr := lr.CreateChild("R")
	r := root.CreateChild("R")
	rl := r.CreateChild("L")
	rr := r.CreateChild("R")

	// Print the mermaid representation of each node.
	fmt.Println(DrawEdge(root.EdgeInfo(l)))
	fmt.Println(DrawEdge(root.EdgeInfo(r)))
	fmt.Println(DrawEdge(l.EdgeInfo(ll)))
	fmt.Println(DrawEdge(l.EdgeInfo(lr)))
	fmt.Println(DrawEdge(ll.EdgeInfo(lll)))
	fmt.Println(DrawEdge(ll.EdgeInfo(llr)))
	fmt.Println(DrawEdge(lr.EdgeInfo(lrl)))
	fmt.Println(DrawEdge(lr.EdgeInfo(lrr)))
	fmt.Println(DrawEdge(r.EdgeInfo(rl)))
	fmt.Println(DrawEdge(r.EdgeInfo(rr)))

	// Output:
	// root --- L
	// root --- R
	// L --- LL
	// L --- LR
	// LL ~~~ LLL
	// LL ~~~ LLR
	// LR ~~~ LRL
	// LR ~~~ LRR
	// R ~~~ RL
	// R ~~~ RR
}
