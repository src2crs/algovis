package binarytrees

import "github.com/src2crs/algovis/binarytrees/binarytreenode"

// BinaryTree represents a binary tree.
// It has a single field, Root, which is a pointer to the root node of the tree.
type BinaryTree struct {
	root *binarytreenode.Node
}

// NewBinaryTree creates a new binary tree with an empty root node.
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{root: binarytreenode.NewRootNode()}
}

// IsEmpty returns true if the tree is empty.
// A tree is empty if its root node is empty.
func (t *BinaryTree) IsEmpty() bool {
	return t.root.IsEmpty()
}
