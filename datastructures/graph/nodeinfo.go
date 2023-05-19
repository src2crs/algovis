package graph

// NodeInfo is an interface for a node in a graph that is used for drawing the graph.
//
// A node has an id, a label and a visibility.
// The id is used to identify the node and possibly to describe the graph's structure.
// The label can be used to display a text on the node.
// Nodes can be invisible. In this case, they are not drawn, but they still occupy space.
// This can be used to influence the graph's layout.
//
// E.g. in a tree, the id could be the node's position in the tree
// and the label could be the node's value/data.
// For drawing a tree, invisible nodes can be used to fill up each node's children
// to ensure that all nodes are drawn at their correct position.
type NodeInfo interface {
	Id() string    // Id returns the node's id.
	Label() string // Label returns the node's label.
	Visible() bool // Visible returns whether the node is visible.
}
