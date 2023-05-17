package bintrees

// Node is a node in a binary tree.
// Each node has a path string, which is used to identify the node's position in the tree.
// The path string is a sequence of characters 'L' or 'R', where each character represents a direction in the tree.
// The path string is empty for the root node.
type Node struct {
	Path  string // Stores the position of the node in the tree.
	Left  *Node  // Left child node.
	Right *Node  // Right child node.
}

// NewNode creates a new tree node with a given path string.
func NewNode(path string) *Node {
	return &Node{Path: path}
}

// NewRootNode creates a new tree node with an empty path string.
func NewRootNode() *Node {
	return NewNode("")
}
