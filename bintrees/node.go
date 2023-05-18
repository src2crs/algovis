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

// CreateChild creates a new child node of the current node.
// The new node is created on the left or right side of the current node, depending on the given direction.
// The path string of the new node is the path string of the current node, with the given direction appended.
// The new node is returned.
// If the given direction is neither "L" nor "R", a panic occurs.
// If a node already exists on the given side, it is overwritten.
func (n *Node) CreateChild(direction string) *Node {
	if direction == "L" {
		n.Left = NewNode(n.Path + direction)
		return n.Left
	} else if direction == "R" {
		n.Right = NewNode(n.Path + direction)
		return n.Right
	} else {
		panic("Invalid direction")
	}
}
