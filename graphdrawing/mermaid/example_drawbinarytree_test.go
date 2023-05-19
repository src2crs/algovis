package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/binarytrees"
	"github.com/src2crs/algovis/binarytrees/binarytreenode"
)

// ExampleDrawBinaryTree_nodes creates a binary tree with several nodes.
// It explicitly creates a root node and several child nodes.
// It then retrieves the mermaid representation of the tree using DrawBinaryTree().
func ExampleDrawBinaryTree_nodes() {
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
	root := binarytreenode.NewRootNode()
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
	fmt.Println(DrawBinaryTreeNodes(root))

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

// ExampleDrawBinaryTree_tree creates a binary tree with several nodes.
// It creates a binarytrees.BinaryTree and uses its CreateNode() method to create the nodes.
// It then retrieves the mermaid representation of the tree using DrawBinaryTree().
func ExampleDrawBinaryTree_tree() {
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
	tree := binarytrees.NewBinaryTree()
	tree.CreateNode("LLL")
	tree.CreateNode("LRL")
	tree.CreateNode("RL")

	// Draw the tree.
	fmt.Println(DrawBinaryTree(tree))

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
