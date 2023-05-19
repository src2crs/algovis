package mermaid

import (
	"strings"

	"github.com/src2crs/algovis/binarytrees"
	"github.com/src2crs/algovis/binarytrees/binarytreenode"
)

// DrawBinaryTreeNodes returns the mermaid representation of a binary tree.
func DrawBinaryTreeNodes(root *binarytreenode.Node) string {
	// Create a slice to store the node and edge strings
	resultstrings := []string{"graph TD"}

	// Create a traversal function that appends each node's NodeInfo and outgoing EdgeInfo strings to the slice.
	traversalFunc := func(node *binarytreenode.Node) bool {
		resultstrings = append(resultstrings, DrawNode(node.NodeInfo()))
		if !node.IsEmpty() {
			resultstrings = append(resultstrings, DrawEdge(node.EdgeInfo(node.Left)))
			resultstrings = append(resultstrings, DrawEdge(node.EdgeInfo(node.Right)))
		}
		return true
	}

	// Traverse the tree.
	root.Visit(traversalFunc)

	// return the mermaid string.
	return strings.Join(resultstrings, "\n  ")
}

// DrawBinaryTree returns the mermaid representation of a binary tree.
func DrawBinaryTree(tree *binarytrees.BinaryTree) string {
	return DrawBinaryTreeNodes(tree.Root())
}
