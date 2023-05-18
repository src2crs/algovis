package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/binarytrees"
)

// ExampleDrawNode_binarytree creates a binary tree with several nodes
// and draws their mermaid representations using DrawNode().
func ExampleDrawNode_binarytree() {
	// Create a binary tree with the following structure:
	//   root
	//   ├── L
	//   │   ├── LL
	//   │   │   ├── LLL (invisible)
	//   │   │   └── LLR (invisible)
	//   │   └── LR
	//   | 	 	 ├── LRL (invisible)
	//   | 	 	 ├── LRR (invisible)
	//   └── R
	//       └── RL (invisible)
	//       └── RR (invisible)
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
	for _, node := range []*binarytrees.Node{root, l, ll, lll, llr, lr, lrl, lrr, r, rl, rr} {
		fmt.Println(DrawNode(node.NodeInfo()))
	}

	// Output:
	// root
	// L[L]
	// LL[LL]
	// LLL[LLL]:::hidden
	// LLR[LLR]:::hidden
	// LR[LR]
	// LRL[LRL]:::hidden
	// LRR[LRR]:::hidden
	// R[R]
	// RL[RL]:::hidden
	// RR[RR]:::hidden
}
