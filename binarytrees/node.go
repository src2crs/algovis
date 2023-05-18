package binarytrees

import "github.com/src2crs/algovis/graphdrawing"

// Node is a node in a binary tree.
// Each node has a path string, which is used to identify the node's position in the tree.
// The path string is a sequence of characters 'L' or 'R', where each character represents a direction in the tree.
// The path string is empty for the root node.
type Node struct {
	Path  string // Stores the position of the node in the tree.
	label string // Label to be used when drawing the node. If empty, the path string is used instead.
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

// GetChild returns the child node of the current node on the given side.
// If no node exists on the given side, nil is returned.
// If the given direction is neither "L" nor "R", a panic occurs.
func (n *Node) GetChild(direction string) *Node {
	switch direction {
	case "L":
		return n.Left
	case "R":
		return n.Right
	default:
		panic("Invalid direction")
	}
}

// IsEmpty returns true if the current node is empty.
// A node is empty if it has no children.
func (n *Node) IsEmpty() bool {
	return n.Left == nil && n.Right == nil
}

// Label returns the label of the current node for drawing.
// If the label is empty, the path string is returned instead.
func (n *Node) Label() string {
	if n.label != "" {
		return n.label
	}
	return n.Path
}

// SetLabel sets the label of the current node for drawing.
func (n *Node) SetLabel(label string) {
	n.label = label
}

// Name returns the name of the current node for drawing.
// The name is the path string of the current node
// or "root" if the current node is the root node.
func (n *Node) Name() string {
	if n.Path == "" {
		return "root"
	}
	return n.Path
}

// NodeInfo returns information for drawing the current node.
func (n *Node) NodeInfo() graphdrawing.NodeInfo {
	return graphdrawing.NodeInfo{
		Name:    n.Name(),
		Label:   n.Label(),
		Visible: !n.IsEmpty(),
	}
}

// EdgeInfo returns information for drawing an edge from the current node to the given target node.
// The target node must not be nil.
// If the current node or the target node is empty, the edge is not visible.
func (n *Node) EdgeInfo(target *Node) graphdrawing.EdgeInfo {
	if target == nil {
		panic("Target node must not be nil")
	}
	return graphdrawing.EdgeInfo{
		Source:  n.Name(),
		Target:  target.Name(),
		Visible: !n.IsEmpty() && !target.IsEmpty(),
	}
}
