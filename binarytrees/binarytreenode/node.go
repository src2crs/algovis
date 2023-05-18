package binarytreenode

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
	if direction != "L" && direction != "R" {
		panic("Invalid direction")
	}

	result := NewNode(n.Path + direction)

	switch direction {
	case "L":
		n.Left = result
		if n.Right == nil {
			n.Right = NewNode(n.Path + "R")
		}
	case "R":
		n.Right = result
		if n.Left == nil {
			n.Left = NewNode(n.Path + "L")
		}
	}
	return result
}

// CreateNodeAtSubPos creates a new node below the current node.
// The new node is created at the given sub-position denoted by the given path string.
// The path string of the new node is the path string of the current node, with the given sub-position appended.
// The new node is returned.
// Any non existing nodes on the path to the sub-position are created with empty labels.
// If the given path string is empty, the current node is returned without changing it.
// If the node at the given sub-position already exists, it is overwritten.
func (n *Node) CreateNodeAtSubPos(path string) *Node {
	if path == "" {
		return n
	}
	nextdir, rest := path[:1], path[1:]
	switch nextdir {
	case "L":
		if n.Left == nil {
			n.Left = NewNode(n.Path + nextdir)
		}
		return n.Left.CreateNodeAtSubPos(rest)
	case "R":
		if n.Right == nil {
			n.Right = NewNode(n.Path + nextdir)
		}
		return n.Right.CreateNodeAtSubPos(rest)
	default:
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
// A node is empty if any of its children is nil.
func (n *Node) IsEmpty() bool {
	return n.Left == nil || n.Right == nil
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

// Visit traverses the tree in pre-order and calls the given function for each node.
// The given function is called with the current node as argument.
// If the given function returns false, the traversal is aborted.
// Returns true if the complete tree was traversed, false otherwise.
func (n *Node) Visit(f func(*Node) bool) bool {
	if !f(n) {
		return false
	}
	if n.Left != nil && !n.Left.Visit(f) {
		return false
	}
	if n.Right != nil && !n.Right.Visit(f) {
		return false
	}
	return true
}