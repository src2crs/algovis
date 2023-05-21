package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/datastructures/tree"
)

// ExampleMermaid_binarytree_complete shows how to create a complete binary tree using
// the tree.Tree data type.
// It also shows how to draw it using the Mermaid function.
func ExampleMermaid_binarytree_complete() {
	// Create a complete binary tree of height 3
	t := tree.New()

	// Use the AddNodeAtPath function to add all the leaf nodes.
	// The inner nodes will be added automatically.
	t.AddNodeAtPath("0.0")
	t.AddNodeAtPath("0.1")
	t.AddNodeAtPath("1.0")
	t.AddNodeAtPath("1.1")

	// Draw the tree
	fmt.Println(Mermaid(t))

	// Output:
	// graph TD
	//   root[root]
	//   0[0]
	//   0.0[0.0]
	//   0.1[0.1]
	//   1[1]
	//   1.0[1.0]
	//   1.1[1.1]
	//   root --- 0
	//   0 --- 0.0
	//   0 --- 0.1
	//   root --- 1
	//   1 --- 1.0
	//   1 --- 1.1
}

// ExampleMermaid_ternarytree_complete shows how to create an incomplete binary tree using
// the tree.Tree data type.
// It also shows how to draw it using the Mermaid function.
func ExampleMermaid_ternarytree_complete() {
	// Create a ternary tree of height 2
	t := tree.New()

	// Use the AddNodeAtPath function to add all the leaf nodes.
	// The inner nodes will be added automatically.
	t.AddNodeAtPath("0.0")
	t.AddNodeAtPath("0.1")
	t.AddNodeAtPath("0.2")
	t.AddNodeAtPath("1.0")
	t.AddNodeAtPath("1.1")
	t.AddNodeAtPath("1.2")
	t.AddNodeAtPath("2.0")
	t.AddNodeAtPath("2.1")
	t.AddNodeAtPath("2.2")

	// Draw the tree
	fmt.Println(Mermaid(t))

	// Output:
	// graph TD
	//   root[root]
	//   0[0]
	//   0.0[0.0]
	//   0.1[0.1]
	//   0.2[0.2]
	//   1[1]
	//   1.0[1.0]
	//   1.1[1.1]
	//   1.2[1.2]
	//   2[2]
	//   2.0[2.0]
	//   2.1[2.1]
	//   2.2[2.2]
	//   root --- 0
	//   0 --- 0.0
	//   0 --- 0.1
	//   0 --- 0.2
	//   root --- 1
	//   1 --- 1.0
	//   1 --- 1.1
	//   1 --- 1.2
	//   root --- 2
	//   2 --- 2.0
	//   2 --- 2.1
	//   2 --- 2.2
}
