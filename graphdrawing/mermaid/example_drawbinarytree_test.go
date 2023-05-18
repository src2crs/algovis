package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/binarytrees"
)

func ExampleDrawBinaryTree() {
	// Create a binary tree with the following structure:
	//   root
	//   ├── L
	//   │   ├── LL
	//   │   │   ├── LLL
	//   │   │   └── LLR
	//   │   └── LR
	//   | 	 	 ├── LRL
	//   | 	 	 ├── LRR
	//   └── R
	//       └── RL
	//       └── RR
	//
	root := binarytrees.NewRootNode()
	l := root.CreateChild("L")
	ll := l.CreateChild("L")
	ll.CreateChild("L")
	ll.CreateChild("R")
	lr := l.CreateChild("R")
	lr.CreateChild("L")
	lr.CreateChild("R")
	r := root.CreateChild("R")
	r.CreateChild("L")
	r.CreateChild("R")

	// Draw the tree.
	fmt.Println(DrawBinaryTree(root))

	// Output:
	// graph TD
	//   root
	//   root --- L
	//   root --- R
	//   L[L]
	//   L --- LL
	//   L --- LR
	//   LL[LL]
	//   LL ~~~ LLL
	//   LL ~~~ LLR
	//   LLL[LLL]:::hidden
	//   LLR[LLR]:::hidden
	//   LR[LR]
	//   LR ~~~ LRL
	//   LR ~~~ LRR
	//   LRL[LRL]:::hidden
	//   LRR[LRR]:::hidden
	//   R[R]
	//   R ~~~ RL
	//   R ~~~ RR
	//   RL[RL]:::hidden
	//   RR[RR]:::hidden
}
