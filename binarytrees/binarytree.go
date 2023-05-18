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

// CreateNode creates a new node in the tree at the given path.
// The path string is a sequence of characters 'L' or 'R', where each character represents a direction in the tree.
// The path string is empty for the root node.
func (t *BinaryTree) CreateNode(path string) {
	t.root.CreateNodeAtSubPos(path)
}

// CreateLabelledNode creates a new node in the tree at the given path with the given label.
// The path string is a sequence of characters 'L' or 'R', where each character represents a direction in the tree.
// The path string is empty for the root node.
func (t *BinaryTree) CreateLabelledNode(path, label string) {
	t.root.CreateNodeAtSubPos(path).SetLabel(label)
}
